package animalshelter

import "testing"

func testNoAnimals(shelter *AnimalShelter) func(t *testing.T) {
	return func(t *testing.T) {
		var want *AnimalSpec
		var wantOk bool
		want, wantOk = nil, false
		got, gotOk := shelter.Dequeue()
		if gotOk != wantOk {
			t.Fatalf("want %t, got %t", wantOk, gotOk)
		}
		if got != want {
			t.Fatalf("want %p, got %p", want, got)
		}
	}
}

func testNoCats(shelter *AnimalShelter) func(t *testing.T) {
	return func(t *testing.T) {
		var want *AnimalSpec
		var wantOk bool
		want, wantOk = nil, false
		got, gotOk := shelter.DequeueCat()
		if gotOk != wantOk {
			t.Fatalf("want %t, got %t", wantOk, gotOk)
		}
		if got != want {
			t.Fatalf("want %p, got %p", want, got)
		}
	}
}
func testNoDogs(shelter *AnimalShelter) func(t *testing.T) {
	return func(t *testing.T) {
		var want *AnimalSpec
		var wantOk bool
		want, wantOk = nil, false
		got, gotOk := shelter.DequeueDog()
		if gotOk != wantOk {
			t.Fatalf("want %t, got %t", wantOk, gotOk)
		}
		if got != want {
			t.Fatalf("want %p, got %p", want, got)
		}
	}
}

func testDequeue(shelter *AnimalShelter, wantName string, wantSpecies Species) func(t *testing.T) {
	return func(t *testing.T) {
		wantOk := true

		got, gotOk := shelter.Dequeue()

		if gotOk != wantOk {
			t.Fatalf("want %t, got %t", wantOk, gotOk)
		}
		if wantName != got.Name {
			t.Fatalf("wanted %s, got %s", wantName, got.Name)
		}
		if wantSpecies != got.Species {
			t.Fatalf("wanted %s, got %s", wantSpecies.toString(), got.Species.toString())
		}
	}
}

func testDequeueCat(shelter *AnimalShelter, wantName string) func(t *testing.T) {
	return func(t *testing.T) {
		wantOk := true
		wantSpecies := CAT

		got, gotOk := shelter.DequeueCat()

		if gotOk != wantOk {
			t.Fatalf("want %t, got %t", wantOk, gotOk)
		}
		if wantName != got.Name {
			t.Fatalf("wanted %s, got %s", wantName, got.Name)
		}
		if wantSpecies != got.Species {
			t.Fatalf("wanted %s, got %s", wantSpecies.toString(), got.Species.toString())
		}
	}
}

func testDequeueDog(shelter *AnimalShelter, wantName string) func(t *testing.T) {
	return func(t *testing.T) {
		wantOk := true
		wantSpecies := DOG

		got, gotOk := shelter.DequeueDog()

		if gotOk != wantOk {
			t.Fatalf("want %t, got %t", wantOk, gotOk)
		}
		if wantName != got.Name {
			t.Fatalf("wanted %s, got %s", wantName, got.Name)
		}
		if wantSpecies != got.Species {
			t.Fatalf("wanted %s, got %s", wantSpecies.toString(), got.Species.toString())
		}
	}
}

func TestAnimalShelter(t *testing.T) {
	firstCat := AnimalSpec{Name: "alice", Species: CAT}
	firstDog := AnimalSpec{Name: "bob", Species: DOG}
	secondCat := AnimalSpec{Name: "catty", Species: CAT}
	secondDog := AnimalSpec{Name: "doug", Species: DOG}

	t.Run("empty shelter", func(t *testing.T) {
		shelter := AnimalShelter{}
		t.Run("no animals", testNoAnimals(&shelter))
		t.Run("no cats", testNoCats(&shelter))
		t.Run("no dogs", testNoDogs(&shelter))
	})

	t.Run("singleton shelters", func(t *testing.T) {
		t.Run("one cat", func(t *testing.T) {
			wantName := "alice"
			wantSpecies := CAT
			t.Run("dequeue animal is the cat", func(t *testing.T) {
				var shelter AnimalShelter
				shelter.Enqueue(firstCat)

				testDequeue(&shelter, wantName, wantSpecies)(t)
			})
			t.Run("dequeue cat is the cat", func(t *testing.T) {
				var shelter AnimalShelter
				shelter.Enqueue(firstCat)

				testDequeueCat(&shelter, wantName)(t)
			})
			t.Run("no dogs", func(t *testing.T) {
				var shelter AnimalShelter
				shelter.Enqueue(firstCat)

				testNoDogs(&shelter)(t)
			})
		})
		t.Run("one dog", func(t *testing.T) {
			wantName := "bob"
			wantSpecies := DOG
			t.Run("dequeue animal is the dog", func(t *testing.T) {
				var shelter AnimalShelter
				shelter.Enqueue(firstDog)

				testDequeue(&shelter, wantName, wantSpecies)(t)
			})
			t.Run("no cats", func(t *testing.T) {
				var shelter AnimalShelter
				shelter.Enqueue(firstDog)

				testNoCats(&shelter)(t)
			})
			t.Run("dequeue dog is the dog", func(t *testing.T) {
				var shelter AnimalShelter
				shelter.Enqueue(firstDog)

				testDequeueDog(&shelter, wantName)(t)
			})
		})
	})

	t.Run("cat then dog", func(t *testing.T) {
		t.Run("dequeue animal is the cat", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstCat)
			shelter.Enqueue(firstDog)

			testDequeue(&shelter, firstCat.Name, firstCat.Species)(t)
		})
		t.Run("dequeue cat is the cat", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstCat)
			shelter.Enqueue(firstDog)

			testDequeueCat(&shelter, firstCat.Name)(t)
		})
		t.Run("dequeue dog is the dog", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstCat)
			shelter.Enqueue(firstDog)

			testDequeueDog(&shelter, firstDog.Name)(t)
		})
	})

	t.Run("dog then cat", func(t *testing.T) {
		t.Run("dequeue animal is the dog", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstDog)
			shelter.Enqueue(firstCat)

			testDequeue(&shelter, firstDog.Name, firstDog.Species)(t)
		})
		t.Run("dequeue cat is the cat", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstDog)
			shelter.Enqueue(firstCat)

			testDequeueCat(&shelter, firstCat.Name)(t)
		})
		t.Run("dequeue dog is the dog", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstDog)
			shelter.Enqueue(firstCat)

			testDequeueDog(&shelter, firstDog.Name)(t)
		})
	})

	t.Run("multiple animals", func(t *testing.T) {
		t.Run("dequeue animal is the first cat", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstCat)
			shelter.Enqueue(firstDog)
			shelter.Enqueue(secondCat)
			shelter.Enqueue(secondDog)

			testDequeue(&shelter, firstCat.Name, firstCat.Species)(t)
		})
		t.Run("dequeue cat is the first cat", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstCat)
			shelter.Enqueue(firstDog)
			shelter.Enqueue(secondCat)
			shelter.Enqueue(secondDog)

			testDequeueCat(&shelter, firstCat.Name)(t)
		})
		t.Run("dequeue dog is the first dog", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstCat)
			shelter.Enqueue(firstDog)
			shelter.Enqueue(secondCat)
			shelter.Enqueue(secondDog)

			testDequeueDog(&shelter, firstDog.Name)(t)
		})
		t.Run("draining shelter", func(t *testing.T) {
			var shelter AnimalShelter
			shelter.Enqueue(firstCat)
			shelter.Enqueue(firstDog)
			shelter.Enqueue(secondCat)
			shelter.Enqueue(secondDog)

			testDequeue(&shelter, firstCat.Name, firstCat.Species)(t)
			testDequeue(&shelter, firstDog.Name, firstDog.Species)(t)
			testDequeue(&shelter, secondCat.Name, secondCat.Species)(t)
			testDequeue(&shelter, secondDog.Name, secondDog.Species)(t)
			testNoAnimals(&shelter)(t)
		})
	})
}
