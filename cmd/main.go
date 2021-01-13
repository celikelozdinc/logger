package main

import (
	"log"

	pkg "github.com/celikelozdinc/logger/pkg/api"
)

func main() {
	var common string = "SIEMENS"
	log.Printf("Format at 1st version -> %s", pkg.FmtWithv1(common))
	log.Printf("Format at 2nd version -> %s", pkg.FmtWithv2(common))
}
