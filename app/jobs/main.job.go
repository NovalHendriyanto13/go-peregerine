package main

import (
	"log"
	"github.com/hibiken/asynq"
	DI "peregerine/DI"
	"peregerine/app/jobs/resume"
)

func main() {
	di, err := DI.Build()
	if err != nil {
		panic(err)
	}
	server := di.QueueServer

	mux := asynq.NewServeMux()

	// List of init worker for the jobs
	ResumeHandler := resume.NewResumeHandler()

	// List of Handle Function of worker jobs
	mux.HandleFunc(resume.TaskName, ResumeHandler.Process)

	log.Println("🚀 Worker started...")
	if err := server.Qs.Server.Run(mux); err != nil {
		di.Logger.Log.Info(err.Error())
		log.Println(err)
	}
}