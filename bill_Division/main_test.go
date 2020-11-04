//ALDO FUSTER TURPIN

package main

import (
	"strconv"
	"testing"
)

func TestDidBrianOverChargedAnnaInput0(t *testing.T) {
	bill := []int32{3, 10, 2, 9}
	k, b := int32(1), int32(12)

	wasOverCharged, amountOvercharged := didBrianOverChargedAnna(bill, k, b)

	wasOverChargedExpected, amountOverchargedExpected := true, int32(5)

	if wasOverCharged != wasOverChargedExpected {
		t.Error("Expected", strconv.FormatBool(wasOverChargedExpected), "but got", wasOverCharged)
	}

	if amountOvercharged != amountOverchargedExpected {
		t.Error("Expected", strconv.FormatBool(wasOverChargedExpected), "but got", wasOverCharged)
	}
}

func TestDidBrianOverChargedAnnaInput1(t *testing.T) {
	bill := []int32{3, 10, 2, 9}
	k, b := int32(1), int32(7)

	wasOverCharged, amountOvercharged := didBrianOverChargedAnna(bill, k, b)

	wasOverChargedExpected, amountOverchargedExpected := false, int32(0)

	if wasOverCharged != wasOverChargedExpected {
		t.Error("Expected", strconv.FormatBool(wasOverChargedExpected), "but got", wasOverCharged)
	}

	if amountOvercharged != amountOverchargedExpected {
		t.Error("Expected", strconv.FormatBool(wasOverChargedExpected), "but got", wasOverCharged)
	}
}
