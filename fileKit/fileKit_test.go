package fileKit

import (
	"fmt"
	"testing"
)
func TestCreateFile(t *testing.T) {
	bb := make([]byte, 0)
	CreateFile(bb, `d:\\aa.jpg`)
}

func TestMkdir(t *testing.T) {
	e:=CreateMutiDir(`d:\\ccc\cc\c`)
	fmt.Println(e)
}

