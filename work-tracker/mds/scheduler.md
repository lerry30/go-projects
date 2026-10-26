Yes, cron is the right tool, and since you already have the data layer in `internal`, an in-process scheduler with `robfig/cron` fits well. First, one thing to settle about your times.

## Which Monday do the early-morning times belong to?

If this is one overnight window that starts Monday at 9pm, then 1:00am, 2:00am, 4:00am, 4:15am and 6:00am fall on **Tuesday**. If you mean all eight times on Monday itself, they all use weekday `1`. I'll assume the overnight window. Changing it is a one-field edit.

## Data-driven schedule

Storing the times as data is easier to read and change than eight cron strings:

```go
// internal/jobs/schedule.go
package jobs

import (
	"fmt"
	"time"
)

type Slot struct {
	Day    time.Weekday
	Hour   int
	Minute int
}

var WeeklySlots = []Slot{
	{time.Monday, 21, 0},
	{time.Monday, 23, 0},
	{time.Monday, 23, 15},
	{time.Tuesday, 1, 0},
	{time.Tuesday, 2, 0},
	{time.Tuesday, 4, 0},
	{time.Tuesday, 4, 15},
	{time.Tuesday, 6, 0},
}

// Spec returns the cron expression: "minute hour * * weekday"
func (s Slot) Spec() string {
	return fmt.Sprintf("%d %d * * %d", s.Minute, s.Hour, int(s.Day))
}
```

## Scheduler in `internal`

It reuses the same repositories or services your handlers use:

```go
// internal/jobs/scheduler.go
package jobs

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	c      *cron.Cron
	ctx    context.Context
	cancel context.CancelFunc
}

func NewScheduler(loc *time.Location, svc *SomeService) (*Scheduler, error) {
	ctx, cancel := context.WithCancel(context.Background())
	s := &Scheduler{
		c: cron.New(
			cron.WithLocation(loc),
			cron.WithChain(
				cron.Recover(cron.DefaultLogger),            // a panic won't kill the process
				cron.SkipIfStillRunning(cron.DefaultLogger), // no overlapping runs
			),
		),
		ctx:    ctx,
		cancel: cancel,
	}

	for _, slot := range WeeklySlots {
		if _, err := s.c.AddFunc(slot.Spec(), func() {
			if err := svc.Run(s.ctx); err != nil {
				log.Println("job failed:", err)
			}
		}); err != nil {
			cancel()
			return nil, err
		}
	}
	return s, nil
}

func (s *Scheduler) Start() { s.c.Start() }

func (s *Scheduler) Stop(ctx context.Context) error {
	cronCtx := s.c.Stop()
	defer s.cancel()
	select {
	case <-cronCtx.Done():
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
```

## Wiring in `cmd/api/main.go`

```go
loc, _ := time.LoadLocation("Asia/Manila")

sched, err := jobs.NewScheduler(loc, svc)
if err != nil {
	log.Fatal(err)
}
sched.Start()

srv := &http.Server{Addr: ":8080", Handler: router}
go func() {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}()

ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()
<-ctx.Done()

shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

_ = srv.Shutdown(shutdownCtx)
_ = sched.Stop(shutdownCtx)
```

## When in-process isn't the best choice

In-process is fine if you run **one instance** and can tolerate a missed run when the server is down. Otherwise consider these:

- **Multiple API instances:** every instance fires every job, so you'd get 2x or 3x runs. Either run the scheduler in only one instance or add a lock, such as a Postgres advisory lock (`pg_try_advisory_lock`) or a Redis `SET NX` with an expiry.
- **Separate worker binary:** add `cmd/worker/main.go` that only runs the scheduler. It uses the same `internal` packages, but deploys and restarts independently of the API, so an API deploy at 11:10pm doesn't interrupt a job.
- **OS-level trigger:** add `cmd/job/main.go` that runs once and exits, and trigger it from system cron or a systemd timer. With only 8 runs a week, this is simple and robust, and systemd's `Persistent=true` can catch up on missed runs.

For a single-server setup, I'd start with the in-process version and move to a separate worker binary if you run into any of the problems above. Also make each job idempotent (safe to run twice), because restarts and retries will eventually cause a double run. If you tell me what the jobs do, I can suggest how to guard against that.