package cryptopals

import (
	"testing"
)

func TestChallenge1(t *testing.T) {
	result := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")	
	if (result != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t") {
		t.Error("Incorrect string", result)
	}
}
