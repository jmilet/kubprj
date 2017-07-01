package randomstring

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	total := 20
	res := RandStringRunes(total)

	if len(res) != total {
		t.Fatalf("Expected %d and got %d", total, len(res))
	}

	total = 0
	res = RandStringRunes(total)

	if len(res) != total {
		t.Fatalf("Expected %d and got %d", total, len(res))
	}
}

// Silly test, just to ease a run...
func TestRandomStringHighVolume(t *testing.T) {
	d := make(map[string]int)

	total := 1000000

	for i := 0; i < total; i++ {
		d[RandStringRunes(20)] += 1
	}

	if len(d) != total {
		t.Fatalf("Expected %d and got %d", total, len(d))
	}
}
