package config

import (
	"encoding/json"
	"fmt"
	"git.parallelcoin.io/dev/9/pkg/util"
	"io/ioutil"
	"os"
	"path/filepath"

	"git.parallelcoin.io/dev/9/pkg/util/cl"
)

var datadir *string = new(string)

func (app *App) Parse(args []string) int {
	app.Config = MakeConfig(app)

	// parse commandline
	cmd, tokens := app.ParseCLI(args)
	if cmd == nil {
		cmd = app.Commands["help"]
	}
	// get datadir from cli args if given
	if dd, ok := tokens["datadir"]; ok {
		datadir = &dd.Value
		pwd, _ := os.Getwd()
		*datadir = filepath.Join(pwd, *datadir)
		dd.Value = *datadir
		// fmt.Println("datadir set", dd.Value, datadir)
		app.Cats["app"]["datadir"].Value.Put(*datadir)
		DataDir = *datadir
	} else {
		// fmt.Println("datadir default")
		ddd := util.AppDataDir("9", false)
		app.Cats["app"]["datadir"].Put(ddd)
		datadir = &ddd
		DataDir = *datadir
	}
	// now we can initialise the App
	for i, x := range app.Cats {
		for j := range x {
			temp := app.Cats[i][j]
			temp.App = app
			app.Cats[i][j] = temp
		}
	}
	for i, x := range app.Cats {
		for j := range x {
			app.Cats[i][j].Init(app.Cats[i][j])
		}
	}
	// set AppDataDir for running as node
	// fmt.Println("cmd.Name", cmd.Name)
	aa := CleanAndExpandPath(filepath.Join(
		*datadir,
		cmd.Name),
		*datadir)
	app.Config.AppDataDir, app.Config.LogDir = &aa, &aa

	configFile := CleanAndExpandPath(filepath.Join(
		*datadir, "config"), *datadir)
	*app.Config.ConfigFile = configFile
	if !FileExists(configFile) {
		// fmt.Println("config file not found: creating new one at ", configFile)
		if EnsureDir(configFile) {
			// fmt.Println("created new directory to store data", datadir)
		}
		fh, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		j, e := json.MarshalIndent(app, "", "\t")
		if e != nil {
			panic(e)
		}
		_, err = fmt.Fprint(fh, string(j))
		if err != nil {
			panic(err)
		}
	}
	// fmt.Println("reading config", configFile)
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	// fmt.Println("unmarshalling config")
	e := json.Unmarshal(conf, app)
	if e != nil {
		panic(e)
	}
	// app.Config = MakeConfig(app)
	// spew.Dump(app)

	if app.Config.LogLevel != nil {
		// fmt.Println("setting debug level to", *app.Config.LogLevel)
		cl.Register.SetAllLevels(*app.Config.LogLevel)
	}
	// run as configured
	r := cmd.Handler(
		args,
		tokens,
		app)
	// fmt.Println("finished parse", cmd.Name, cmd.Handler, r)
	return r
}

func (app *App) ParseCLI(args []string) (cmd *Command, tokens Tokens) {
	// fmt.Println("args", args)
	// cmds = make(Commands)
	cmd = new(Command)
	// collect set of items in commandline
	if len(args) < 2 {
		fmt.Print("No args given, printing help:\n\n")
		args = append(args, "h")
	}
	commandsFound := make(map[string]int)
	tokens = make(Tokens)
	for _, x := range args[1:] {
		for i, y := range app.Commands {
			if y.RE.MatchString(x) {
				if _, ok := commandsFound[i]; ok {
					// TODO change token to struct{val,command}
					tokens[i] = Token{x, *y}
					commandsFound[i]++
					break
				} else {
					tokens[i] = Token{x, *y}
					commandsFound[i] = 1
					break
				}
			}
		}
	}
	// fmt.Println("tokens", tokens)
	// fmt.Println("commandsFound", commandsFound)
	var withHandlersNames []string
	withHandlers := make(Commands)
	for i := range commandsFound {
		if app.Commands[i].Handler != nil {
			withHandlers[i] = app.Commands[i]
			withHandlersNames = append(withHandlersNames, i)
		}
	}
	invoked := make(Commands)
	for i, x := range withHandlers {
		invoked[i] = x
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
			common = append(common, i)
		}
		for _, x := range common {
			for _, y := range x {
				if y != "" {
					resolved = append(resolved, y)
				}
			}
		}
		if len(resolved) > 1 {
			resolved = uniq(resolved)
			withHandlers = make(Commands)
			common = [][]string{}
			withHandlersNames = resolved
			resolved = []string{}
			for _, x := range withHandlersNames {
				withHandlers[x] = app.Commands[x]
			}
			for _, x := range withHandlersNames {
				i := intersection(withHandlersNames, withHandlers[x].Precedent)
				common = append(common, i)
			}
			for _, x := range common {
				for _, y := range x {
					if y != "" {
						resolved = append(resolved, y)
					}
				}
			}
			resolved = uniq(resolved)
		}
	} else if len(withHandlersNames) == 1 {
		resolved = []string{withHandlersNames[0]}
	}
	if len(resolved) < 1 {
		err := fmt.Errorf(
			"\nunable to resolve which command to run:\n\tinput: '%s'",
			withHandlersNames)
		fmt.Println(err)
		return nil, tokens
	}
	*cmd = *app.Commands[resolved[0]]
	return cmd, tokens
}
