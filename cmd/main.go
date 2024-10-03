package main

import (
	"TZ/internal/handler"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {

	//Разблокируйте этот кусок кода чтобы запустить вручную а остальные в комент)
	//	if len(os.Args) < 3 {
	//		fmt.Println("Команда для запуска: go run cmd/main.go <файл.json> <num-workers>")
	//		return
	//	}
	//	fileName := os.Args[1]
	//	numWorkers, err := strconv.Atoi(os.Args[2])
	//	if err != nil {
	//		fmt.Printf("Invalid number of workers: %s\n", os.Args[2])
	//		return
	//	}
	//	err = handler.ProcessFile(fileName, numWorkers)
	//	if err != nil {
	//		fmt.Printf("Error processing file: %v\n", err)
	//	}

	rand.Seed(time.Now().UnixNano())
	fileName := "./data.json"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", fileName)
		return
	}

	done := make(chan bool) //kanal dlya sinkhrona do sled zapuska

	for {
		numWorkers := rand.Intn(46) + 5

		fmt.Printf("Processing file with %d workers...\n", numWorkers)

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := handler.ProcessFile(fileName, numWorkers)
			if err != nil {
				fmt.Printf("Error processing file: %v\n", err)
			}
		}()
		go func() {
			wg.Wait()
			done <- true
		}()

		<-done
		fmt.Println("Процесс завершен следуший запуск после 7 секунд...")

		time.Sleep(7 * time.Second)
	}
}
