package main

import (
	"fmt"
	"math/rand"
	"os"
	. "strings"
	"time"

	. "github.com/tokenshift/dice"
)

type option int
const (
	OPT_AVERAGE option = iota
	OPT_MAXIMUM
	OPT_MINIMUM
	OPT_RANDOM
	OPT_VERBOSE
)

func main() {
	rand.Seed( time.Now().UTC().UnixNano())

	options := []option{}
	specs := []Dice{}

	for _, arg := range os.Args[1:] {
		if HasPrefix(arg, "-") {
			switch arg {
			case "-a", "--avg", "--average", "--mean":
				options = append(options, OPT_AVERAGE)
			case "-M", "--max", "--maximum":
				options = append(options, OPT_MAXIMUM)
			case "-m", "--min", "--minimum":
				options = append(options, OPT_MINIMUM)
			case "-r", "--result", "--rand", "--random", "--roll":
				options = append(options, OPT_RANDOM)
			case "-v", "--verbose":
				options = append(options, OPT_VERBOSE)
			default:
				fmt.Fprintln(os.Stderr, "Invalid option:", arg)
				os.Exit(1)
			}
		} else {
			dice, err := Parse(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			specs = append(specs, dice)
		}
	}

	if len(options) == 0 {
		options = append(options, OPT_RANDOM)
	}

	if len(specs) == 0 {
		specs = append(specs, MustParse("1d6"))
	}

	for _, spec := range specs {
		for _, option := range options {
			switch option {
			case OPT_AVERAGE:
				fmt.Print(spec.Mean())
			case OPT_MAXIMUM:
				fmt.Print(spec.Max())
			case OPT_MINIMUM:
				fmt.Print(spec.Min())
			case OPT_RANDOM:
				fmt.Print(spec.RollAll())
			case OPT_VERBOSE:
				rolls, mod := spec.RollEach()
				total := 0

				for i, r := range rolls {
					if i > 0 {
						fmt.Print("+")
					}

					total += r.Result
					fmt.Print(r.Result)
				}

				if mod != 0 {
					total += mod
					fmt.Printf("+%d", mod)
				}

				fmt.Printf(" = %d", total)
			}

			fmt.Print("\t")
		}
		fmt.Println("")
	}
}