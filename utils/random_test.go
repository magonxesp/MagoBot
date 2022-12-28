package utils

import "testing"

func TestRandomInt(t *testing.T) {
	random := RandomInt(0, 10)

	if random < 0 {
		t.Fatalf("The random is less than 0, got %d", random)
	}

	if random > 10 {
		t.Fatalf("The random is greater than 10, got %d", random)
	}
}

func TestRandomInt_MultipleTimes(t *testing.T) {
	times := 1000000

	for i := 0; i < times; i++ {
		random := RandomInt(0, 10)

		if random < 0 {
			t.Fatalf("The random is less than 0, got %d", random)
		}

		if random > 10 {
			t.Fatalf("The random is greater than 10, got %d", random)
		}
	}
}
