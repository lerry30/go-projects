package in_process_scheduling

import (
	"fmt"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robfig/cron/v3"
)

/*
func RunSchedule(db *pgxpool.Pool) {
	schedulesService := NewSchedulesService(db)

	usersSched := Sched(schedulesService)

	//fmt.Println(usersSched[0].FirstName)

	for _, t := range usersSched[0].SchedTimeLogs["Monday"] {
		fmt.Println(t.Format())
	}
}
	*/

type Scheduler struct {
	backgroundService *cron.Cron
	ctx context.Context
	cancel context.CancelFunc
}

func NewScheduler(loc *time.Location, db *pgxpool.Pool) (*Scheduler, error) {
	ctx, cancel := context.WithCancel(context.Background())
	s := &Scheduler{
		backgroundService: cron.New(
			cron.WithLocation(loc),
			cron.WithChain(
				cron.Recover(cron.DefaultLogger), // A panic won't kill the process
				cron.SkipIfStillRunning(cron.DefaultLogger), // No overlapping runs
			),
		),
		ctx: ctx,
		cancel: cancel,
	}

	schedulesService := NewSchedulesService(db)
	usersSched := Sched(schedulesService)

	weekdaysSched := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	for _, user := range usersSched {
		for _, wds := range weekdaysSched {
			//fmt.Println(user.SchedTimeLogs[wds][0].Format("2006-01-02 15:04:05"))
			for i, _ := range user.SchedTimeLogs[wds] {
				//fmt.Println(user.Spec(wds, i))
				if _, err := s.backgroundService.AddFunc(user.Spec(wds, i), func() {
					fmt.Println("test")
				}); err != nil {
					cancel()
					return nil, err
				}
			}
		}
	}

	return s, nil
}

func (s *Scheduler) Start() { s.backgroundService.Start() }

func (s *Scheduler) Stop(ctx context.Context) error {
	cronCtx := s.backgroundService.Stop()
	defer s.cancel()

	select {
	case <-cronCtx.Done():
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}