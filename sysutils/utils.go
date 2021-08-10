package sysutils

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Utilities interface {
	BackGroundColor() string
	ForeGroundColor() string
	TextStyle() string
}

var ClearCode = "\033[0m"

type AnsiColor struct {
	BgColor, FgColor, TextStyle string
}

type ProgressBarInfo struct {
	SliceItems                       []string
	CompleteSymbol, UncompleteSymbol string
	CountdownLength                  int
}

func (p *ProgressBarInfo) StartProgressBar() {
	percentComplete := 0.0

	red := AnsiColor{
		BgColor:   "0",
		FgColor:   "\033[38;5;197m",
		TextStyle: "0",
	}

	yellow := AnsiColor{
		BgColor:   "0",
		FgColor:   "\033[38;5;220m",
		TextStyle: "0",
	}

	green := AnsiColor{
		BgColor:   "0",
		FgColor:   "\033[38;5;48m",
		TextStyle: "0",
	}

	for {

		for j := 0; j < len(p.SliceItems) + 1; j++ {

			time.Sleep(200 * time.Millisecond)

			percentConversion := 100.0 / float64(len(p.SliceItems))

			progressDisplay := fmt.Sprintf("%.2f - [%s%s]", percentComplete, strings.Repeat(p.CompleteSymbol, int(math.Round(percentComplete))), strings.Repeat(p.UncompleteSymbol, int((float64(p.CountdownLength) - math.Round(percentComplete)))))

			if percentComplete <= 32 {
				fmt.Printf("%s%s%s\r", red.FgColor, progressDisplay, ClearCode)
			} else if percentComplete >= 33 && percentComplete <= 64 {
				fmt.Printf("%s%s%s\r", yellow.FgColor, progressDisplay, ClearCode)
			} else {
				fmt.Printf("%s%s%s\r", green.FgColor, progressDisplay, ClearCode)
			}

			percentComplete += percentConversion
			// p.CountdownLength -= int(percentComplete)

			if percentComplete >= 100{
				break
			}
			
		}
		break
	}
	fmt.Println("\nDone...")

}
