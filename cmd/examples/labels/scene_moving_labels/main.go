package main

import (
	"flag"
	director "github.com/gravestench/director/pkg"
	"log"
	"os"
	"runtime/pprof"
)
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
func main() {
	flag.Parse()
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
