package ascii

import 	"testing"

func TestGreetingASCII(t *testing.T) {
	ascii := GreetingASCII()
	if !(isASCII(ascii)){
		t.Fail()
	}
}

// Fra https://play.golang.org/p/hnZzfnbXeF
func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}
