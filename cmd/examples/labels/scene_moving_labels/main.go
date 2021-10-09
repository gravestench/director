package main

import (
	"flag"
	director "github.com/gravestench/director/pkg"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var traceit = flag.String("trace", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *traceit != "" {
		f, err := os.Create("trace.out")
		if err != nil {
			log.Fatalf("failed to create trace output file: %v", err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatalf("failed to close trace file: %v", err)
			}
		}()

		if err := trace.Start(f); err != nil {
			log.Fatalf("failed to start trace: %v", err)
		}
		defer trace.Stop()
	}
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	d := director.New()

	d.AddScene(&MovingLabelsScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
