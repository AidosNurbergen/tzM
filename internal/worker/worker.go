package worker

import (
	ps "TZ/internal/utils"
	"sync"
)

type Task struct {
	Data []ps.Data
}

type Result struct {
	Sum int64
}

func Worker(id int, tasks <-chan Task, results chan<- Result, done chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		var sum int64
		for _, item := range task.Data {
			sum += item.A + item.B
		}
		results <- Result{Sum: sum}
	}
	done <- id //dlya 2 proverki workera o zaverweni raboty
}

func StartWorkers(numWorkers int, tasks chan Task, results chan Result, done chan int) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go Worker(i, tasks, results, done, &wg)
	}

	wg.Wait()
	close(results)
	close(done)
}

func DispatchTasks(data []ps.Data, chunkSize int, tasks chan Task) {
	//delenie i itpravka v kanal
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		tasks <- Task{Data: data[i:end]}
	}
	close(tasks)
}
