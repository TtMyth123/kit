package kit
import (
	"fmt"
	"reflect"
	"testing"
)
func TestGetGuid(t *testing.T) {
	bb := GetGuid()
	fmt.Println(bb)
}

func TestGetGuid11(t *testing.T) {
	type AA struct {
		A1 int
		A2 int64
		A3 string
	}

	aa := AA{A1: 1, A2: 2, A3: "3"}
	fmt.Println(aa)
	getType := reflect.TypeOf(aa)

	getType.FieldByNameFunc(func(s string) bool {
		sf,ok :=getType.FieldByName(s)
		if ok {
			fmt.Println("Name:",sf.Name)

		}
		return true
	});
}



func TestGetGuid2(t *testing.T) {
	 aa := GetGuidEx()
	 fmt.Println(aa)
}
