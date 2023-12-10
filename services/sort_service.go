// services/sort_service.go
package services

import (
	"sort"
	"sync"
	"time"
	"fmt"
	"backend/models"
)

func SortSubArraySequential(subArray []int) {
	sort.Ints(subArray)
}

func SortSubArrayConcurrent(subArray []int, wg *sync.WaitGroup, ch chan bool) {
	defer func() {
		wg.Done()
		ch <- true
	}()

	sort.Ints(subArray)
}

func ProcessSortSequential(request models.SortRequest) models.SortResponse {
	startTime := time.Now().UnixNano()

	for _, subArray := range request.ToSort {
		SortSubArraySequential(subArray)
	}

	timeTaken := time.Now().UnixNano()

	response := models.SortResponse{
		SortedArrays:    request.ToSort,
		TimeNanoseconds: int64(timeTaken-startTime),
	}
	fmt.Println(startTime)
	fmt.Println(timeTaken)
	fmt.Println(startTime)
	
	return response
}

func ProcessSortConcurrent(request models.SortRequest) models.SortResponse {
	var wg sync.WaitGroup
	ch := make(chan bool, len(request.ToSort))

	startTime := time.Now()

	for _, subArray := range request.ToSort {
		wg.Add(1)
		go SortSubArrayConcurrent(subArray, &wg, ch)
	}

	wg.Wait()
	close(ch)

	for range ch {
		// Drain the channel to wait for all goroutines to finish
	}

	timeTaken := time.Since(startTime)

	response := models.SortResponse{
		SortedArrays:    request.ToSort,
		TimeNanoseconds: int64(timeTaken),
	}

	return response
}