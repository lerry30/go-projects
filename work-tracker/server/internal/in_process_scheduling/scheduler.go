package in_process_scheduling

import (
	"fmt"
	"context"
	"time"
	"sync"
	//"log"

	//"tracker/internal/webhook"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	mu sync.Mutex
	backgroundService *cron.Cron
	ctx context.Context
	cancel context.CancelFunc
	user *User
	running bool
}

var weekdaysSched = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
var serviceIdentities = []string{"login", "1st-break-out", "1st-break-in", "lunch-out", "lunch-in", "2nd-break-out", "2nd-break-in", "logout"}

func NewScheduler(loc *time.Location, user *User) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &Scheduler{
		backgroundService: cron.New(
			cron.WithLocation(loc),
			cron.WithChain(
				cron.Recover(cron.DefaultLogger), // A panic won't kill the process
				cron.SkipIfStillRunning(cron.DefaultLogger), // No overlapping runs
			),
		),
		ctx: ctx,
		cancel: cancel,
		user: user,
		running: false,
	}
}

func (s *Scheduler) Create() error {
	//service := webhook.NewWebhookAppscript(user.FirstName, user.LastName, user.Endpoint, user.Status, user.SpreadSheetID)
	for _, wds := range weekdaysSched {
		//fmt.Println(user.SchedTimeLogs[wds][0].Format("2006-01-02 15:04:05"))
		for i, _ := range s.user.SchedTimeLogs[wds] {
			//fmt.Println(user.Spec(wds, i))
			/*var srvIt string
			if i >= len(serviceIdentities) {
				srvIt = "log"
			} else {
				srvIt = serviceIdentities[i]
			}*/

			if _, err := s.backgroundService.AddFunc(s.user.Spec(wds, i), func() {
				fmt.Println("Test schedule service")
				/*
				if err := service.NotifyAppscript(s.ctx, srvIt); err != nil {
					log.Println("job failed: ", err)
				}
					*/
			}); err != nil {
				s.cancel()
				return err
			}
		}
	}
	return nil
}

func (s *Scheduler) Start() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.running {
		return
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.backgroundService.Start()
	s.running = true
}

func (s *Scheduler) Stop(ctx context.Context) error {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return nil
	}
	s.running = false
	cronCtx := s.backgroundService.Stop()
	cancel := s.cancel
	s.mu.Unlock() // don't hold the lock while waiting

	defer cancel()

	select {
	case <-cronCtx.Done():
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Scheduler) Restart(ctx context.Context) error {
	if err != s.Stop(ctx); err != nil {
		return err
	}
	s.Start()
	return nil
}