package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"

	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/component"
	"dolsh/ui/folder"
	"dolsh/ui/login"
	"dolsh/ui/player"
	"dolsh/ui/setup"
)

var profile = flag.Bool("profile", false, "write cpu.prof and trace.out")

func main() {
	flag.Parse()

	if *profile {
		prof, err := os.Create("cpu.prof")
		if err != nil {
			panic(err)
		}
		defer prof.Close()

		pprof.StartCPUProfile(prof)
		defer pprof.StopCPUProfile()

		tf, err := os.Create("trace.out")
		if err != nil {
			panic(err)
		}
		defer tf.Close()

		trace.Start(tf)
		defer trace.Stop()
	}

	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	a, err := app.New()
	if err != nil {
		fmt.Println("config error:", err)
		os.Exit(1)
	}

	var start tea.Model
	switch a.RestoreSession() {
	case app.SessionReady:
		start = player.New(a)
	case app.SessionNeedsLibrary:
		start = library.New(a, component.Size{})
	case app.SessionNeedsLogin:
		start = login.New(a, component.Size{})
	default:
		start = setup.New(a)
	}

	p := tea.NewProgram(start)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
