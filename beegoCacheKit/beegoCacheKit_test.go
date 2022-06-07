package beegoCacheKit

import (
	"fmt"
	"testing"
	"time"
)

func TestA1(t *testing.T) {
	t1 := time.Now()
	t2 := t1.AddDate(0, 0, -40)
	NewBeegoCache("AA")
	fmt.Println(t1, t2)
}
