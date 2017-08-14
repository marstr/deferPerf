package deferPerf

import (
	"time"
)

func executeStandard() {
	subject := make(chan struct{})
	close(subject)
}

func executeDeferred() {
	subject := make(chan struct{})
	defer close(subject)
}

func executeLongStandard() {
	subject := make(chan struct{})
	time.Sleep(time.Second)
	close(subject)
}

func executeLongDeferred() {
	subject := make(chan struct{})
	defer close(subject)
	time.Sleep(time.Second)
}
