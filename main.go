package main

import (
	"short-url/httpServer"
	"short-url/postgres"
	"short-url/workerGenerator"

	_ "github.com/lib/pq"
)

var workerGen *workerGenerator.WorkerUrlGenerator

func main() {
	databaseConnection := postgres.New()

	workerGen = workerGenerator.New()
	workerGen.Init(databaseConnection)

	for i := 0; i < workerGenerator.NumOfWorkers; i++ {
		go workerGen.Worker()
	}

	server := httpServer.New()
	server.Init(workerGen)
	server.StartServer()

	databaseConnection.Db.Close()
	workerGen.Close()
}
