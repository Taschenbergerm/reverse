package reverse

import (
	"testing"
)


func TestReverse(t *testing.T) {
	testInput := []string{"This", "is", "a", "Test"}
	expect := "tseT a si sihT"
	got := RevSlice(testInput)
	if got != expect {
		t.Errorf("Expected %s \n Got %s", expect,got)
	}
}
