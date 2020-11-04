//ALDO FUSTER TURPIN

package main

import (
	"testing"
)

func TestSockMerchant(t *testing.T) {
	n := int32(9)
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	result := sockMerchant(n, ar)

	expected := int32(3)

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}
