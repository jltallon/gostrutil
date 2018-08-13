package cmdvec


import (
	"fmt"
	"testing"
)

var data = []string{ "this", "%%", "this", "is", "%%", "this", "is", "a", "%%", "this", "is", "a", "test" }


func TestSplitMulti(t *testing.T) {

	res := SplitMulti(data, "%%")
	fmt.Println(res)


//	t.Fail()
}
