package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	dataCh := make(chan int)
	wg.Add(1)
	go Source(wg, "data1.dat", dataCh)

	wg.Add(1)
	go Source(wg, "data2.dat", dataCh)

	evenCh, oddCh := Splitter(dataCh)
	evenSumCh := Sum(evenCh)
	oddSumCh := Sum(oddCh)

	wg.Wait()
	close(dataCh)

	done := Merger(evenSumCh, oddSumCh)
	<-done
}

func Source(wg *sync.WaitGroup, fileName string, dataCh chan<- int) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if val, err := strconv.Atoi(line); err == nil {
			dataCh <- val
		}
	}

}

func Splitter(dataCh <-chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	go func() {
		for no := range dataCh {
			if no%2 == 0 {
				evenCh <- no
			} else {
				oddCh <- no
			}
		}
		close(evenCh)
		close(oddCh)
	}()
	return evenCh, oddCh
}

func Sum(valCh <-chan int) <-chan int {
	sumCh := make(chan int)
	go func() {
		var sum int
		for val := range valCh {
			sum += val
		}
		sumCh <- sum
		// close(sumCh)
	}()
	return sumCh
}

func Merger(evenSumCh <-chan int, oddSumCh <-chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		resultFile, err := os.Create("total.txt")
		if err != nil {
			log.Fatalln(err)
		}
		defer resultFile.Close()
		for range 2 {
			select {
			case evenSum := <-evenSumCh:
				fmt.Fprintf(resultFile, "Even Total : %d\n", evenSum)
			case oddSum := <-oddSumCh:
				fmt.Fprintf(resultFile, "Odd Total : %d\n", oddSum)
			}
		}
		close(doneCh)
	}()
	return doneCh
}
