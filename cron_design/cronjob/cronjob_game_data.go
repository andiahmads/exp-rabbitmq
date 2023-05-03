package cronjob

import (
	gameService "design-pattern/cron_design/service"
	"fmt"
	"log"
	"strconv"
)

type serviceGameCron struct {
	GameService gameService.Service
}

func newServiceGame(serviceGameData gameService.Service) *serviceGameCron {
	return &serviceGameCron{
		serviceGameData,
	}
}

func (s *serviceGameCron) Start() {
	cronJob, found := CronJobs[GAME_DATA]

	fmt.Println(found)
	if !found {
		log.Printf("job `%s` not registered", GAME_DATA)
		return
	}

	if cronJob.IsStart {
		log.Printf("Job scheduler %s already running.", GAME_DATA)
		return
	}

	entryID, err := cronJob.JobCron.AddFunc("@every "+strconv.Itoa(cronJob.Interval)+"s", s.Execute)
	if err != nil {
		log.Printf("error creating cron job %s", cronJob.Name)
		return
	}

	cronJob.JobEntryID = entryID
	cronJob.JobCron.Start()
	cronJob.IsStart = true

}

func (s *serviceGameCron) Execute() {
	fmt.Println("running")
	scheduleCron := CronJobs[GAME_DATA]
	log.Printf("job scheduler %s execute", scheduleCron.Name)

	err := s.GameService.InserData()
	if err != nil {
		log.Printf("error creating cron job %s", err.Error())

	}

}
