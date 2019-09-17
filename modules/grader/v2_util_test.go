package grader

import (
	"fmt"
	"testing"
)

func TestV2Payout(t *testing.T) {
	for i := -10; i < 100; i++ {
		t.Run(fmt.Sprintf("V2 Payout %d", i), func(t *testing.T) {
			got := V2Payout(i)
			if i >= 0 && i < 25 && got != 200e8 {
				t.Errorf("V2Payout() = %v, want %v", got, 200e8)
			}
			if (i < 0 || i >= 25) && got != 0 {
				t.Errorf("V2Payout() = %v, want %v", got, 0)
			}
		})
	}
}
