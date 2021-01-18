package puzzle

import (
//	"log"
)

type Bottle struct {
	Pills        int
	PillWeight   float64
	WeightExcess float64
}

func NewBottle(pills int, weight float64) *Bottle {
	var bottle = Bottle{
		PillWeight:   weight,
		Pills:        pills,
		WeightExcess: float64(pills) * 0.1,
	}
	return &bottle
}

func NewBottles(numberOfBottles, heaviestBottle int) map[int]*Bottle {
	var bottles = make(map[int]*Bottle)
	for i := 0; i < numberOfBottles; i++ {
		var pills = i + 1
		var weight float64 = 1.0
		if heaviestBottle == pills {
			weight = 1.1
		}
		var bottle = NewBottle(pills, weight)
		bottles[pills] = bottle
	}
	return bottles
}

func getRealTotalWeight(bottles map[int]*Bottle) float64 {
	var totalWeight float64

	for _, bottle := range bottles {
		var bottleWeight = bottle.PillWeight * float64(bottle.Pills)
		totalWeight = totalWeight + bottleWeight
	}

	return totalWeight
}

func getHeavierBottle(numberOfBottles, heaviestBottle int) *Bottle {
	if numberOfBottles == 0 {
		var emptyBottle = NewBottle(0, 0.0)
		return emptyBottle
	}

	if numberOfBottles == 1 {
		var singleBottle = NewBottle(1, 1.0)
		return singleBottle
	}

	var bottles = NewBottles(numberOfBottles, heaviestBottle)

	var realTotalWeight = getRealTotalWeight(bottles)
	var totalWeight = int(realTotalWeight)
	var diff = realTotalWeight - float64(totalWeight)

	var heavierBottle *Bottle
	for _, bottle := range bottles {
		if int(bottle.WeightExcess*100.0) == int(diff*100.0) {
			heavierBottle = bottle
			break
		}
	}

	return heavierBottle
}
