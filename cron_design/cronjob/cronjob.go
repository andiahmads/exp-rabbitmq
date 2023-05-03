package cronjob

import (
	"fmt"

	gameService "design-pattern/cron_design/service"

	cron "gopkg.in/robfig/cron.v2"
)

const (
	GAME_DATA     = "game_data" //
	SCHEDULEQUEUE = "schedule_queue"
)

// Manager cron manager
type Manager struct {
	scheduleApplyGame *serviceGameCron
}

type JobInfo struct {
	Module     interface{}
	Name       string
	Interval   int
	Limit      int
	Function   interface{}
	JobCron    *cron.Cron
	IsIdle     bool
	IsStart    bool
	JobEntryID cron.EntryID
}

var CronJobs map[string]*JobInfo = make(map[string]*JobInfo)

// NewManager new cron manager
func NewManager(gameDataService gameService.Service) *Manager {

	scheduleGameData := newServiceGame(gameDataService)
	cronManager := Manager{
		scheduleGameData,
	}

	return &cronManager
}

// Start start all job
func (manager *Manager) Start() {
	manager.scheduleApplyGame.Start()
}

// JobInfo register job
func (manager *Manager) JobInfo() map[string]*JobInfo {
	fmt.Println(CronJobs)

	CronJobs[GAME_DATA] = &JobInfo{
		Name:     GAME_DATA,
		Interval: 10, // 10 seconds
		Limit:    1,
		JobCron:  cron.New(),
	}

	return CronJobs
}
