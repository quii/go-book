package tempcov

import "testing"

func TestItKelvins(t *testing.T) {
	// Given
	c := FreezingC

	// When
	k := CToK(c)

	// Then
	if k != FreezingK {
		t.Error("I expected", FreezingK, "but i got", k)
	}
}
