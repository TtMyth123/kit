package fileKit

import "testing"

func TestCreateFile(t *testing.T) {
	bb := make([]byte, 0)
	CreateFile(bb, `d:\\aa.jpg`)
}
