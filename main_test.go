package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type outputSpeedTest struct {
	input            []string
	speed            int
	expectedDuration time.Duration
}

var outputSpeedTests = []outputSpeedTest{
	{
		input:            []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		speed:            1,
		expectedDuration: time.Second * 10,
	},
	{
		input:            []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		speed:            2,
		expectedDuration: time.Second * 5,
	},
	{
		input:            []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		speed:            4,
		expectedDuration: time.Millisecond * 2500,
	},
}

func TestOutputSpeed(t *testing.T) {
	for _, test := range outputSpeedTests {
		var wg sync.WaitGroup
		wg.Add(1)
		inChan := make(chan string)
		start := time.Now()
		go output(&wg, inChan, test.speed)
		for _, in := range test.input {
			inChan <- in
		}
		close(inChan)
		wg.Wait()
		finished := time.Now()
		expected := start.Add(test.expectedDuration)

		assert.WithinDuration(t, expected, finished, time.Millisecond*10)
	}
}
