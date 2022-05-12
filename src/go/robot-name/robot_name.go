package robotname

import (
	"errors"
	"math"
	"math/rand"
)

// Initially implemented as a slice, but that was inefficient
var allocatedNames = make(map[string]bool)

// Makes checking "are all names allocated" easier
// In proper production, we'd make this an actual object to ensure invariants
var numAllocatedNames = 0

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		newName, err := generateName()
		if err != nil {
			return newName, err
		}
		r.name = newName
	}
	return r.name, nil

}

func (r *Robot) Reset() {
	r.name = ""

}

func allocateName(name string) {
	// Later product requirements would probably require looking up robot-by-name, so we'd change
	// the type of the map
	allocatedNames[name] = true
	numAllocatedNames++
}

func unallocateName(name string) {
	// No error handling if name is not present
	// Yet another thing that I'm surprised that the standard library doesn't provide
	allocatedNames[name] = false
	// Technically we should check that the name was already allocated before doing this decrement
	numAllocatedNames--
}

func generateName() (string, error) {
	if areAllNamesAllocated() {
		return "", errors.New("All names allocated")
	}
	candidateName := generateNameCandidate()
	for nameHasBeenAllocated(candidateName) {
		candidateName = generateNameCandidate()
	}
	allocateName(candidateName)
	return candidateName, nil
}

func randomLetter() string {
	return string("ABCDEFGHIJKLMNOPQRSTUVWXYZ"[rand.Intn(26)])
}

func randomNumber() string {
	return string("0123456789"[rand.Intn(10)])
}

func nameHasBeenAllocated(name string) bool {
	return allocatedNames[name] == true
}

// generateNameCandidate generates a _possible_ name, but does not check for if it has been
// allocated already (that responsibility lies with `generateName`)
func generateNameCandidate() string {
	return randomLetter() + randomLetter() + randomNumber() + randomNumber() + randomNumber()
}

func areAllNamesAllocated() bool {
	// Total number of possible names = 26^2 (2 letters) * 10^3 (3 numbers)
	// > should be impossible, but just in case!
	return numAllocatedNames >= int(math.Pow(26, 2)*math.Pow(10, 3))
}
