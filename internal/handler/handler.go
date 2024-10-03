package handler

import (
	ps "TZ/internal/utils"
	"TZ/internal/worker"
	"fmt"
)

func ProcessFile(fileName string, numWorkers int) error {

	data, err := ps.ParseJSON(fileName)
	if err != nil {
		return fmt.Errorf("failed to parse JSON file: %w", err)
	}

	msSize := (len(data) + numWorkers - 1) / numWorkers

	tasks := make(chan worker.Task, numWorkers)
	results := make(chan worker.Result, numWorkers)
	done := make(chan int, numWorkers) // Канал для завершённых рктин

	go worker.StartWorkers(numWorkers, tasks, results, done)

	worker.DispatchTasks(data, msSize, tasks)

	var totalSum int64
	/*{
		var completedWorkers int

		for {
			select {
			case result, ok := <-results:
				if ok {
					totalSum += result.Sum
				}
			case _, ok := <-done:
				if ok {
					completedWorkers++
					//fmt.Printf("Worker %d finished\n", i)
				} else {
					//fmt.Printf("All workers finished. Total workers: %d\n", completedWorkers)
					fmt.Printf("Total summ: %d\n", totalSum)
					return nil
				}
			}
		}
	}*/
	{
		for result := range results {
			totalSum += result.Sum
		}

		fmt.Printf("Total summm: %d\n", totalSum)
		return nil
	}
}
