package director

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"

	"github.com/gravestench/director/pkg"
)

// Director is the scene manager
type Director = pkg.Director

var ( // these are for command line flags
	debug = flag.Bool("debug", false, "print debug info")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	traceit = flag.String("trace", "", "write cpu profile to file")
)

// New returns a new director instance
func New() *Director {
	instance := pkg.New()

	flag.Parse()

	handleProfiling()
	handleDebug(instance)

	return instance
}

func handleDebug(d *Director) {
	if *debug {
		d.AddScene(Scene.Primitives()[ScenePrimitiveTickGraph])
	}
}

func handleProfiling() {
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

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}

		defer pprof.StopCPUProfile()
	}
}
