package workerGenerator

import (
	"fmt"
	"short-url/postgres"
	"short-url/randomUrlGenerator"
)

const (
	NumOfWorkers int = 5
)

type WorkerUrlGenerator struct {
	workerChan         chan string
	DatabaseConnection *postgres.PostgresDB
	randomUrlGen       *randomUrlGenerator.RandomUrlGenerator
}

func New() *WorkerUrlGenerator {
	w := WorkerUrlGenerator{}
	return &w
}

func (e *WorkerUrlGenerator) Init(databaseConnection *postgres.PostgresDB) {
	e.workerChan = make(chan string, NumOfWorkers)
	e.DatabaseConnection = databaseConnection
	e.randomUrlGen = randomUrlGenerator.New()
}

func (e *WorkerUrlGenerator) Worker() {
	for websiteUrl := range e.workerChan {
		for {
			newShortUrl := e.randomUrlGen.GenerateUrl()
			fmt.Println("Trying to linking: ", websiteUrl, " to this shorturl: ", newShortUrl)
			if e.DatabaseConnection.CheckIfShortUrlNotExists(newShortUrl) {
				e.DatabaseConnection.InsertShortUrlAndWebsite(newShortUrl, websiteUrl)
				break
			}
		}
	}
}

func (e *WorkerUrlGenerator) AddWebsite(website string) {
	e.workerChan <- website
}

func (e *WorkerUrlGenerator) Close() {
	close(e.workerChan)
}
