package number

import (
	"testing"
)

func TestRiyaInt_GetNative(t *testing.T) {
	r := newRiya(1)
	one, err := r.GetNative()
	if err != nil {
		t.Error(err)
	}
	if one != 1 {
		t.Error("No pass")
	}
	r2 := newRiya(2)
	two, err := r2.GetNative()
	if err != nil {
		t.Error(err)
	}
	if two != 2 {
		t.Error("no pass")
	}
}

func TestRiyaInt_IsEqualTo(t *testing.T) {
	r := newRiya(10)
	if r.IsEqualTo(10) != true {
		t.Error("Failed")
	}

	r2 := newRiya(2)
	if r2.IsEqualTo(2) != true {
		t.Error("Failed")
	}
}

func TestRiyaInt_IsGreaterThan(t *testing.T) {
	n := newRiya(45)
	if n.IsGreaterThan(36) == true {
		t.Error("Failed")
	}
}
func TestRiyaInt_IsLessThan(t *testing.T) {
	n := newRiya(23)
	if n.IsLessThan(13) != true {
		t.Error("Failed")
	}
}

func TestRiyaInt_Add(t *testing.T) {
	thiryFour := newRiya(34)
	if thiryFour.num != 34 {
		t.Error("Failed")
	}
	ten := newRiya(10)
	if ten.num != 10 {
		t.Error("Failed")
	}
	fourtyFour := thiryFour.Add(ten)
	if fourtyFour.num != 44 {
		t.Error("Failed")
	}
}

func TestRiyaInt_Subtract(t *testing.T) {
	thirtyFive := newRiya(35)
	ten := newRiya(10)
	twentyFive := thirtyFive.Subtract(ten)

	if ten.num != 10 {
		t.Error("Failed")
	}
	if thirtyFive.num != 35 {
		t.Error("Failed")
	}

	if twentyFive.num != 25 {
		t.Error("Failed")
	}
}
func TestRiyaInt_Multiply(t *testing.T) {
	thirtyFive := newRiya(35)
	if thirtyFive.num != 35 {
		t.Error("Failed")
	}

	ten := newRiya(10)
	if ten.num != 10 {
		t.Error("Failed")
	}
	threeHundredfifty := thirtyFive.Multiply(ten)
	if threeHundredfifty.num != 350 {
		t.Error("Failed")
	}
}

func TestRiyaInt_DivideBy(t *testing.T) {
	thirtyFive := newRiya(35)
	if thirtyFive.num != 35 {
		t.Error("Failed")
	}
	ten := newRiya(10)
	if ten.num != 10 {
		t.Error("Failed")
	}
	three, err := thirtyFive.DivideBy(ten)
	if err!= nil {
		t.Error(err.Error())
	}
	if three.num != 3 {
		t.Error("Failed")
	}

	fiftyFive := newRiya(55)
	if fiftyFive.num != 55 {
		t.Error("Failed")
	}
	zero := newRiya(0)
	if zero.num != 0 {
		t.Error("Failed")
	}
	_, err = fiftyFive.DivideBy(zero)
	if err == nil {
		t.Error("Expects an error")
	}

}
