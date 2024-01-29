package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	router := gin.Default()

	// Create a new cron scheduler
	c := cron.New()

	// Define your cron jobs
	job1 := cronJob{ID: "job1", Schedule: "*/5 * * * *", Action: func() { fmt.Println("Executing Job 1 , every 5 minute") }}
	job2 := cronJob{ID: "job2", Schedule: "*/10 * * * *", Action: func() { fmt.Println("Executing Job 2") }}
	job3 := cronJob{ID: "job3", Schedule: "0 0 * * *", Action: func() { fmt.Println("Executing Job 3") }}

	// Add the cron jobs to the scheduler
	addCronJob(c, &job1)
	addCronJob(c, &job2)
	addCronJob(c, &job3)

	// Start the cron scheduler
	c.Start()

	// Define your API routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Cron Job Server"})
	})

	// Run the Gin server on port 8080
	err := router.Run(":5000")
	if err != nil {
		log.Fatal(err)
	}
}

type cronJob struct {
	ID       string
	Schedule string
	Action   func()
}

func addCronJob(c *cron.Cron, job *cronJob) {
	_, err := c.AddFunc(job.Schedule, func() {
		fmt.Printf("[%s] Running cron job...\n", job.ID)
		job.Action()
	})
	if err != nil {
		log.Fatalf("Failed to add cron job %s: %v", job.ID, err)
	}
}
