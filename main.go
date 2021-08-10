package main

import "github.com/mrcoggsworth/go-progbar/sysutils"

func main() {
	sliceExample := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x",
		"y", "z",
	}

	pbi := sysutils.ProgressBarInfo{
		SliceItems:       sliceExample,
		CompleteSymbol:   "#",
		UncompleteSymbol: ".",
		CountdownLength:  100,
	}

	pbi.StartProgressBar()
}
