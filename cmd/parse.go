package cmd

import (
	"fmt"

	"git.parallelcoin.io/dev/9/pkg/util/cl"
)

func Parse(args []string) int {
	// parse commandline
	err := parseCLI(args)
	if err != nil {
		log <- cl.Error{err}
	}
	// read configuration

	return 0
}

func parseCLI(args []string) error {
	log <- cl.Info{args}
	// collect set of items in commandline
	if len(args) < 2 {
		log <- cl.Info{"no args given, starting GUI"}
		return nil
	}
	commandsFound := make(map[string]int)
	for _, x := range args[1:] {
		for _, y := range commandsList {
			if commands[y].RE.Match([]byte(x)) {
				if _, ok := commandsFound[y]; ok {
					log <- cl.Debug{"found", y, x}
					commandsFound[y]++
					break
				} else {
					log <- cl.Debug{"found", y, x}
					commandsFound[y] = 1
					break
				}
			}
		}
	}

	var withHandlersNames []string
	withHandlers := make(Commands)
	for i := range commandsFound {
		if commands[i].Handler != nil {
			log <- cl.Debug{"found", i}
			withHandlers[i] = commands[i]
			withHandlersNames = append(withHandlersNames, i)
		}
	}
	// search the precedents of each in the case of multiple
	// with handlers and delete the one that has another in the
	// list of matching handlers. If one is left we can run it,
	// otherwise return an error.
	var resolved []string
	if len(withHandlersNames) > 1 {

		var common [][]string
		for _, x := range withHandlersNames {
			i := intersection(withHandlersNames, withHandlers[x].Precedent)
			log <- cl.Debug{"intersection", withHandlersNames, ".", withHandlers[x].Precedent, "==", i}
			common = append(common, i)
		}
		for _, x := range common {
			for _, y := range x {
				if y != "" {
					resolved = append(resolved, y)
				}
			}
		}

		for _, i := range resolved {
			log <- cl.Debug{"--> resolved", i}
		}
	} else if len(withHandlersNames) == 1 {
		resolved = []string{withHandlersNames[0]}
	}
	if len(resolved) < 1 {
		err := fmt.Errorf("unable to resolve which command to run, found multiple: %v", withHandlersNames)
		return err
	} else {
		log <- cl.Debug{"running", resolved}
	}

	return nil
}

func intersection(a, b []string) (out []string) {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				out = append(out, x)
			}
		}
	}
	return
}

func resolveCommand(withHandlers Commands) Commands {
	return withHandlers
}
