package main

import (
	"log"
	"github.com/hibiken/asynq"
	DI "peregerine/DI"
	"peregerine/app/jobs/resume"
)

func main() {
	di, err := DI.BuildWorker()
	if err != nil {
		panic(err)
	}
	defer di.Close()

	server := di.QueueServer

	mux := asynq.NewServeMux()

	// List of init worker for the jobs
	ResumeHandler := resume.NewResumeHandler(di.Redis.RedisRepo, di.Logger)

	// List of Handle Function of worker jobs
	mux.HandleFunc(resume.TaskName, ResumeHandler.Process)

	log.Println("🚀 Worker started...")
	// Run blocks until SIGINT/SIGTERM, then waits for in-flight tasks to finish
	if err := server.Qs.Server.Run(mux); err != nil {
		di.Logger.Log.Error(err.Error())
		log.Println(err)
	}
}
