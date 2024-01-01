package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Kbgjtn/__eet.git/cases"
    "github.com/Kbgjtn/__eet.git/man"
)

func lab() {
    fmt.Printf("Res: %+v\n", 0)
}

func main() {
	startTime := time.Now()

    lab()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Runtime: %s\n", elapsedTime)
	printMemoryUsage()
}

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Usage: %v KB\n", m.Alloc/1024)
}
