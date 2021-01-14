package main

import (
	"log"
	"strings"

	pkg "github.com/celikelozdinc/logger/pkg/api"
)

func main() {
	EscapeAnalysis()
	Goroutines()
}

// EscapeAnalysis is prepared for demo purpose
func EscapeAnalysis() {
	dummy := new(int)
	log.Printf("Value of dummy -> %d\n", dummy)
	var common string = "SIEMENS"
	log.Printf("Format at 1st version -> %s", pkg.FmtWithv1(common))
	log.Printf("Format at 2nd version -> %s", pkg.FmtWithv2(common))
}

// Goroutines is prepared for demo purpose
func Goroutines() {
	log.Println(strings.Repeat("#", 20))
	done := make(chan bool)
	shared := make(map[int]string)
	go func() {
		log.Println("...Entering goroutine...")
		shared[0] = "dardanelles"
		log.Println("...Exiting goroutine...")
		done <- true
	}()
	shared[1] = "hellespont"
	_ = <-done
	for i, v := range shared {
		log.Printf("at [%d] -> %q", i, v)
	}
	log.Println(strings.Repeat("#", 20))
}
