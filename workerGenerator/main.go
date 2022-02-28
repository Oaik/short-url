package workerGenerator

import (
	"fmt"
	"short-url/insertInstance"
	"short-url/randomUrlGenerator"
	"short-url/sqlInstance"
)

type WorkerUrlGenerator struct {
	workerChan   chan string
	numOfWorkers int
	SqlInstance  *sqlInstance.SqlInstance
	randomUrlGen *randomUrlGenerator.RandomUrlGenerator
}

func New(sqlInstance *sqlInstance.SqlInstance, numOfWorkers, numOfRandomLetters, numOfRandomDigits int) *WorkerUrlGenerator {
	e := WorkerUrlGenerator{}
	e.numOfWorkers = numOfWorkers
	e.workerChan = make(chan string, e.numOfWorkers)
	e.SqlInstance = sqlInstance
	e.randomUrlGen = randomUrlGenerator.New(numOfRandomLetters, numOfRandomDigits)
	for i := 0; i < numOfWorkers; i++ {
		go e.Worker()
	}
	return &e
}

func (e *WorkerUrlGenerator) Worker() {
	for websiteUrl := range e.workerChan {
		sqlInst := insertInstance.New()
		sqlInst.AddPair("website", websiteUrl)
		for {
			newShortUrl := e.randomUrlGen.GenerateUrl()
			fmt.Println("Trying to linking: ", websiteUrl, " to this shorturl: ", newShortUrl)
			if e.SqlInstance.CheckIfShortUrlNotExists(newShortUrl) {
				sqlInst.AddPair("shorturl", newShortUrl)
				sqlInst.Execute(e.SqlInstance.PostgresDB, "url")
				// e.SqlInstance.InsertShortUrlAndWebsite(newShortUrl, websiteUrl)
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
