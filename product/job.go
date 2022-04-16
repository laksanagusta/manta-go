package product

import (
	"sync"
)

type Job interface {
	DispatchWorkers(jobs <-chan []interface{}, wg *sync.WaitGroup)
}

type job struct {
	service Service
}

func NewJob(service Service) *job {
	return &job{service}
}

const totalWorker = 20

func (j *job) DispatchWorkers(jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				j.service.CreateProductBulk(workerIndex, counter, job)
				wg.Done()
				counter++
			}
		}(workerIndex, jobs, wg)
	}
}
