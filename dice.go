package dice

import (
	. "bytes"
	"fmt"
	"math/rand"
	"sort"
	. "strconv"

	. "github.com/prataprc/goparsec"
)

// A single rolled die.
type Die struct {
	Sides, Result int
}

// A set of dice, plus a single constant/modifier that will be added to any
// result.
type Dice struct {
	ds map[int]int
	mod int
}

// Parse a dice spec, panicking if it is not parseable.
func MustParse(spec string) Dice {
	dice, err := Parse(spec)

	if err != nil {
		panic(err)
	}
	
	return dice
}

// Parse a dice spec of the form `d10+2d6+3` and return a rollable Dice object.
func Parse(spec string) (Dice, error) {
	s := NewScanner([]byte(spec))
	result, s := parseDiceSpec(s)

	if s.Endof() {
		return result.(Dice), nil
	} else {
		return Dice{}, fmt.Errorf("Could not parse \"%s\"", spec)
	}
}

// Roll each of the dice and return all of the individual results.
func (d Dice) RollEach() ([]Die, int) {
	results := []Die{}

	for _, sides := range d.diceDescending() {
		results = append(results, Die { sides, rand.Intn(sides)+1 })
	}

	return results, d.mod
}

// Roll all of the dice and return the total.
func (d Dice) RollAll() int {
	total := 0

	for sides, count := range d.ds {
		for i := 0; i < count; i+=1 {
			total += rand.Intn(sides)+1
		}
	}

	return total + d.mod
}

// Return the mean/average roll for a set of dice.
func (d Dice) Mean() float64 {
	return (float64(d.Min()) + float64(d.Max())) / 2.0
}

// Return the minimum roll for a set of dice.
func (d Dice) Min() int {
	min := 0

	for _, count := range d.ds {
		min += count
	}

	return min + d.mod
}

// Return the maximum roll for a set of dice..
func (d Dice) Max() int {
	max := 0

	for sides, count := range d.ds {
		max += count * sides
	}

	return max + d.mod
}

// Returns all of the dice types (by number of sides) in descending order. NOT
// trimmed for duplicates, so 3d6+1d4 will return [6,6,6,4]. Does not return the
// modifier.
func (d Dice) diceDescending() []int {
	dice := []int{}

	for sides, count := range d.ds {
		for i := 0; i < count; i+=1 {
			dice = append(dice, -sides)
		}
	}

	sort.Ints(dice)

	for i, d := range dice {
		dice[i] = -d
	}

	return dice
}

// Combine two dice sets.
func (d *Dice) merge(d2 Dice) {
	for die, count := range(d2.ds) {
		if oldCount, ok := d.ds[die]; ok {
			d.ds[die] = oldCount + count
		} else {
			d.ds[die] = count
		}
	}

	d.mod += d2.mod
}

func (d Dice) String() string {
	var b Buffer

	for die, count := range(d.ds) {
		fmt.Fprintf(&b, "%dd%d+", count, die)
	}
	fmt.Fprint(&b, d.mod)

	return b.String()
}

// Dice Spec Parsing

var parseDiceSpec = Kleene(mergeDice, parseChunk, plus)
var parseChunk    = OrdChoice(nth(0), parseDie, parseMod)
var parseDie      = And(toDie, Int(), d, Int())
var parseMod      = And(toMod, Int())

var plus = Token(`\+`, "PLUS")
var d    = Token(`d`, "D")

func mergeDice(ns []ParsecNode) ParsecNode {
	dice := Dice {
		ds: map[int]int{},
		mod: 0,
	}

	for _, n := range(ns) {
		dice.merge(n.(Dice))
	}

	return dice
}

func nth(i int) Nodify {
	return func(ns []ParsecNode) ParsecNode {
		return ns[i]
	}
}

func toDie(ns []ParsecNode) ParsecNode {
	count, err := ParseInt(ns[0].(*Terminal).Value, 10, 0)
	if err != nil {
		return nil
	}

	sides, err := ParseInt(ns[2].(*Terminal).Value, 10, 0)
	if err != nil {
		return nil
	}

	return Dice {
		ds: map[int]int{
			int(sides): int(count),
		},
		mod: 0,
	}
}

func toMod(ns []ParsecNode) ParsecNode {
	mod, err := ParseInt(ns[0].(*Terminal).Value, 10, 0)
	if err != nil {
		return nil
	}

	return Dice {
		ds: map[int]int{},
		mod: int(mod),
	}
}