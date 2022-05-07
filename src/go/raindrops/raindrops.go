package raindrops

import (
	"fmt"
	"sort"
)

type IntToStringMap map[int]string

// Cannot define methods on named types :(

// FactorToStringMap cannot be const because reasons
var FactorToStringMap = IntToStringMap{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// keysOfMap is necessary for two reasons:
// * It just doesn't exist on maps
// * Iterating over a map with `range` has inconsistent ordering
func (m IntToStringMap) keysOfMap() []int {
	// Cannot do `keys := [len(m)]int{}`
	// because array bound length must be constant
	var keys []int
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func Convert(number int) (output string) {
	anyFactorFound := false
	for _, key := range FactorToStringMap.keysOfMap() {
		if number%key == 0 {
			output += FactorToStringMap[key]
			anyFactorFound = true
		}
	}
	if !anyFactorFound {
		output = fmt.Sprintf("%d", number)
	}
	return
}
