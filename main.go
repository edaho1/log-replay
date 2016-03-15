package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	inputFile string
	speed     int64
)

func init() {
	flag.StringVar(&inputFile, "i", "", "Input file (if not using stdin)")
	flag.Int64Var(&speed, "s", 100, "Lines per second")
}

func main() {
	flag.Parse()

	var scanner *bufio.Scanner
	if inputFile != "" {
		fd, err := os.Open(inputFile)
		if err != nil {
			log.Fatalf("error opening input file: %s", err)
		}
		scanner = bufio.NewScanner(fd)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	readChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go output(&wg, readChan, speed)

	for scanner.Scan() {
		line := scanner.Text()
		readChan <- line
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %s", err)
	}
	close(readChan)
	wg.Wait()
}

func output(wg *sync.WaitGroup, input <-chan string, speed int64) {
	tickerInterval := time.Nanosecond * time.Duration(1e9/speed)
	ticker := time.NewTicker(tickerInterval)
	for line := range input {
		select {
		case <-ticker.C:
			fmt.Fprintf(os.Stdout, "%s\n", line)
		}
	}
	wg.Done()
}
