package main

import (
	cronjobService "design-pattern/cron_design/cronjob"
	"design-pattern/cron_design/repository"
	"design-pattern/cron_design/service"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "user:password@tcp(127.0.0.1:3306)/cakra_internal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connection to database good...")

	gameRepository := repository.NewRepository(db)
	gameService := service.NewService(gameRepository)

	cronManager := cronjobService.NewManager(gameService)

	cronManager.Start()
	cronManager.JobInfo()

	fmt.Println("cron stating....")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// When you push CTRL+C close worker gracefully
	<-sig

}
