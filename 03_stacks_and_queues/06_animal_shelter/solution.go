package animalshelter

import (
	"sync"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type Species uint8

const (
	DOG Species = iota
	CAT
)

func (spec Species) toString() string {
	switch spec {
	case DOG:
		return "dog"
	case CAT:
		return "cat"
	default:
		panic("unhandled animal")
	}
}

type animal struct {
	name  string
	order uint64
}

type AnimalShelter struct {
	mu sync.Mutex

	nextOrder uint64

	cats lib.ArrayQueue[*animal]
	dogs lib.ArrayQueue[*animal]
}

type AnimalSpec struct {
	Name    string
	Species Species
}

func (shelter *AnimalShelter) Enqueue(pet AnimalSpec) {
	shelter.mu.Lock()
	defer shelter.mu.Unlock()

	ele := animal{
		name:  pet.Name,
		order: shelter.nextOrder,
	}
	shelter.nextOrder++

	switch pet.Species {
	case CAT:
		shelter.cats.Add(&ele)
	case DOG:
		shelter.dogs.Add(&ele)
	default:
		panic("unhandled animal")
	}
}

func (shelter *AnimalShelter) Dequeue() (*AnimalSpec, bool) {
	shelter.mu.Lock()
	defer shelter.mu.Unlock()

	eldestCat, catExists := shelter.cats.Peek()
	eldestDog, dogExists := shelter.dogs.Peek()

	var ret *animal
	var retOk bool
	var species Species
	if !catExists || (dogExists && eldestDog.order < eldestCat.order) {
		ret, retOk = shelter.dogs.Remove()
		species = DOG
	} else {
		ret, retOk = shelter.cats.Remove()
		species = CAT
	}

	if !retOk {
		return nil, retOk
	}

	return &AnimalSpec{Name: ret.name, Species: species}, retOk

}

func (shelter *AnimalShelter) DequeueDog() (*AnimalSpec, bool) {
	shelter.mu.Lock()
	defer shelter.mu.Unlock()

	ret, retOk := shelter.dogs.Remove()
	if !retOk {
		return nil, retOk
	}

	return &AnimalSpec{Name: ret.name, Species: DOG}, retOk
}

func (shelter *AnimalShelter) DequeueCat() (*AnimalSpec, bool) {
	shelter.mu.Lock()
	defer shelter.mu.Unlock()

	ret, retOk := shelter.cats.Remove()
	if !retOk {
		return nil, retOk
	}

	return &AnimalSpec{Name: ret.name, Species: CAT}, retOk
}
