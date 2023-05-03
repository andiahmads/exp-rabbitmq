package cronjob

// import (
// 	"design-pattern/corn_design/service"
// 	"log"
// 	"strconv"
// )

// type ServiceHitungLuas struct {
// 	HitungFactory service.Hitung
// }

// func newServiceHitung(hitungFactory service.Hitung) *ServiceHitungLuas {
// 	return &ServiceHitungLuas{hitungFactory}
// }

// func (s *ServiceHitungLuas) Start() {
// 	cronJob, found := CronJobs[HITUNG_LUAS]
// 	if !found {
// 		log.Printf("job `%s` not registered", HITUNG_LUAS)
// 	}

// 	if cronJob.IsStart {
// 		log.Printf("Job scheduler %s already running.", cronJob.Name)
// 	}

// 	entryID, err := cronJob.JobCron.AddFunc("@every "+strconv.Itoa(cronJob.Interval)+"s", s.Execute)
// 	if err != nil {
// 		log.Printf("error creating cron job %s", cronJob.Name)
// 	}

// 	cronJob.JobEntryID = entryID
// 	cronJob.JobCron.Start()
// 	cronJob.IsStart = true

// }

// func (s *ServiceHitungLuas) Execute() {
// 	scheduleCron := CronJobs[HITUNG_LUAS]
// 	log.Printf("job scheduler %s execute", scheduleCron.Name)

// 	_, err := s.HitungFactory.Luas()
// 	if err != nil {
// 		log.Printf("error creating cron job %s", err.Error())

// 	}

// }
