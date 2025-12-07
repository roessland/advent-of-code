package aocutil

import (
	"log"
	"os"
	"runtime/pprof"
)

func Profile() func() {
	log.Println("Profiling")
	f, err := os.Create("pprof.txt")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
		log.Println("Done profiling")
		log.Println("Run go tool pprof -http localhost:8080 pprof.txt")
	}
}
