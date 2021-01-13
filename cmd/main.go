package main

import (
	"log"

	pkg "github.com/celikelozdinc/logger/pkg/api"
)

func main() {
	var common string = "SIEMENS"
	log.Printf("Format at 1st step -> %s", pkg.Fmt(common))

}
