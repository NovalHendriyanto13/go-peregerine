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

	ResumeHandler := resume.NewResumeHandler()
	mux.HandleFunc(resume.TaskName, ResumeHandler.Process) // ini di ambil di folder jobs method nya

	log.Println("🚀 Worker started...")
	if err := server.Qs.Server.Run(mux); err != nil {
		// di.Logger.Logger.Log.info(err)
		log.Println(err)
	}
}