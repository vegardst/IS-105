package iso

import 	(
	"testing"
	"../iso"
)
// Feiler da ikke alle symboler er fra extended ascii

func TestGreetingExtendedASCII(t *testing.T) {
	ascii := iso.GreetingExtendedASCII()
	if !(isExtendedASCII(ascii)){
		t.Fail()
	}
}

func isExtendedASCII(s string) bool {
	for _, c := range s {
		if c > 255 {
			return false
		}
		if c < 127 {
			return false
		}
	}
	return true
}
