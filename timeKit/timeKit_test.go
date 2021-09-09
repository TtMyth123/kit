package timeKit

import (
	"fmt"
	"testing"
	"time"
)

func TestGetNumList(t *testing.T) {
	t1 := time.Now()
	t2 := t1.AddDate(0,0,-40)
	fmt.Println(t1, t2)
}