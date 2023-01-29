package swapchar

import "testing"

func TestSwapCharactersInSentence (t *testing.T) {
	result := SwapCharactersInSentence("This is Divyansh!")
	if result != "tHIS IS dIVYANSH!" {
		t.Error("Expected: tHIS IS dIVYANSH!  got : ", result)
	}
}