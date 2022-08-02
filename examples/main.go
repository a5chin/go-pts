package main

import (
	"log"
	"pts"
)

func main() {
	path := "assets/sample.pts"
	points := pts.ReadPTS(path)

	log.Print(points.Describe())
}
