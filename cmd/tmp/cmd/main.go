package main

import (
	"fmt"

	. "git.parallelcoin.io/dev/9/cmd/tmp"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	a := App("9",
		Version("v1.9.9"),
		Group("app",
			File("cpuprofile",
				Usage("write cpu profile to this file, empty disables cpu profiling"),
			),
		),
	)
	spew.Dump(a)
	fmt.Println(a.Version())
}