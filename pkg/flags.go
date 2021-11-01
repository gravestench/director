package pkg

import "flag"

const (
	FlagNameDebug = "debug"
	FlagDescDebug = "print debug info"
)

var _ = flag.Bool(FlagNameDebug, false, FlagDescDebug)

const (
	FlagNameProfileCPU = "cpuprofile"
	FlagDescProfileCPU = "write cpu profile to file"
)

var _ = flag.String(FlagNameProfileCPU, "", FlagDescProfileCPU)

const (
	FlagNameTrace = "trace"
	FlagDescTrace = "write trace to file"
)

var _ = flag.String(FlagDescTrace, "", FlagDescTrace)
