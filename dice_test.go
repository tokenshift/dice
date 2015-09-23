package dice

import (
	. "reflect"
	. "testing"
)

func assertEquals(t *T, expected, actual interface{}) bool {
	if DeepEqual(expected, actual) {
		return true
	} else {
		t.Errorf("Expected %v, got %v", expected, actual)
		return false
	}
}

func assertWithin(t *T, low, high, actual int) bool {
	if actual < low {
		t.Errorf("Expected >= %d, got %d", low, actual)
		return false
	} else if actual > high {
		t.Errorf("Expected <= %d, got %d", high, actual)
		return false
	} else {
		return true
	}
}

func assertRollsWithin(t *T, times, low, high int, dice Dice) bool {
	for i := 0; i < times; i+=1 {
		roll := dice.RollAll()
		if !assertWithin(t, low, high, roll) {
			return false
		}
	}

	return true
}

func TestSingleDie(t *T) {
	dice, err := Parse("1d6")
	if err != nil {
		t.Error(err)
		return
	}

	assertEquals(t, 1, dice.Min())
	assertEquals(t, 6, dice.Max())
	assertEquals(t, 3.5, dice.Mean())
	assertRollsWithin(t, 1000, 1, 6, dice)

	rolls, mod := dice.RollEach()

	assertEquals(t, 0, mod)
	assertEquals(t, 1, len(rolls))
	assertEquals(t, 6, rolls[0].Sides)
	assertWithin(t, 1, 6, rolls[0].Result)
}

func TestMultipleDice(t *T) {
	dice, err := Parse("3d4")
	if err != nil {
		t.Error(err)
		return
	}

	assertEquals(t, 3, dice.Min())
	assertEquals(t, 12, dice.Max())
	assertEquals(t, 7.5, dice.Mean())
	assertRollsWithin(t, 1000, 3, 12, dice)

	rolls, mod := dice.RollEach()

	assertEquals(t, 0, mod)
	assertEquals(t, 3, len(rolls))
	assertEquals(t, 4, rolls[0].Sides)
	assertWithin(t, 1, 4, rolls[0].Result)
	assertEquals(t, 4, rolls[1].Sides)
	assertWithin(t, 1, 4, rolls[1].Result)
	assertEquals(t, 4, rolls[2].Sides)
	assertWithin(t, 1, 4, rolls[2].Result)
}

func TestWithModifier(t *T) {
	dice, err := Parse("1d6+4")
	if err != nil {
		t.Error(err)
		return
	}

	assertEquals(t, 5, dice.Min())
	assertEquals(t, 10, dice.Max())
	assertEquals(t, 7.5, dice.Mean())
	assertRollsWithin(t, 1000, 5, 10, dice)

	rolls, mod := dice.RollEach()

	assertEquals(t, 4, mod)
	assertEquals(t, 1, len(rolls))
	assertEquals(t, 6, rolls[0].Sides)
	assertWithin(t, 1, 6, rolls[0].Result)
}

func TestDuplicates(t *T) {
	dice, err := Parse("1d6+4+2d6+2")
	if err != nil {
		t.Error(err)
		return
	}

	assertEquals(t, 9, dice.Min())
	assertEquals(t, 24, dice.Max())
	assertEquals(t, 16.5, dice.Mean())
	assertRollsWithin(t, 1000, 9, 26, dice)

	rolls, mod := dice.RollEach()

	assertEquals(t, 6, mod)
	assertEquals(t, 3, len(rolls))
	assertEquals(t, 6, rolls[0].Sides)
	assertWithin(t, 1, 6, rolls[0].Result)
	assertEquals(t, 6, rolls[1].Sides)
	assertWithin(t, 1, 6, rolls[1].Result)
	assertEquals(t, 6, rolls[2].Sides)
	assertWithin(t, 1, 6, rolls[2].Result)
}

func TestMixedDice(t *T) {
	dice, err := Parse("1d20+2d12+3d6+4")
	if err != nil {
		t.Error(err)
		return
	}

	assertEquals(t, 10, dice.Min())
	assertEquals(t, 66, dice.Max())
	assertEquals(t, 38.0, dice.Mean())
	assertRollsWithin(t, 1000, 10, 66, dice)

	rolls, mod := dice.RollEach()

	assertEquals(t, 4, mod)
	assertEquals(t, 6, len(rolls))
	assertEquals(t, 20, rolls[0].Sides)
	assertWithin(t, 1, 20, rolls[0].Result)
	assertEquals(t, 12, rolls[1].Sides)
	assertWithin(t, 1, 12, rolls[1].Result)
	assertEquals(t, 12, rolls[2].Sides)
	assertWithin(t, 1, 12, rolls[2].Result)
	assertEquals(t, 6, rolls[3].Sides)
	assertWithin(t, 1, 6, rolls[3].Result)
	assertEquals(t, 6, rolls[4].Sides)
	assertWithin(t, 1, 6, rolls[4].Result)
	assertEquals(t, 6, rolls[5].Sides)
	assertWithin(t, 1, 6, rolls[5].Result)

	dice, err = Parse("2d12+4+1d20+3d6")
	if err != nil {
		t.Error(err)
		return
	}

	assertEquals(t, 10, dice.Min())
	assertEquals(t, 66, dice.Max())
	assertEquals(t, 38.0, dice.Mean())
	assertRollsWithin(t, 1000, 10, 66, dice)

	rolls, mod = dice.RollEach()

	assertEquals(t, 4, mod)
	assertEquals(t, 6, len(rolls))
	assertEquals(t, 20, rolls[0].Sides)
	assertWithin(t, 1, 20, rolls[0].Result)
	assertEquals(t, 12, rolls[1].Sides)
	assertWithin(t, 1, 12, rolls[1].Result)
	assertEquals(t, 12, rolls[2].Sides)
	assertWithin(t, 1, 12, rolls[2].Result)
	assertEquals(t, 6, rolls[3].Sides)
	assertWithin(t, 1, 6, rolls[3].Result)
	assertEquals(t, 6, rolls[4].Sides)
	assertWithin(t, 1, 6, rolls[4].Result)
	assertEquals(t, 6, rolls[5].Sides)
	assertWithin(t, 1, 6, rolls[5].Result)
}
