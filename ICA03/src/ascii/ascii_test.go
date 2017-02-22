package ascii

import 	"testing"

func TestGreetingASCII(t *testing.T) {
	ascii := GreetingASCII()
	for i := 0; i < len(ascii); i++ {
		if (isASCII(string(ascii[i])) == false){
			t.Fail()
			break
		}
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