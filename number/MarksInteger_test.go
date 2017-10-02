package number

import (
	"testing"
)

func TestRiyaInt_GetNative(t *testing.T) {
	r := newRiya(1)
	one := r.GetNative()
	if one != 1 {
		t.Error("Failed")
	}
	r2 := newRiya(2)
	two := r2.GetNative()
	if two != 2 {
		t.Error("Failed")
	}
}

func TestRiyaInt_IsEqualTo(t *testing.T) {
	r := newRiya(10)
	r1 := newRiya(10)

	if r.IsEqualTo(r1) != true {
		t.Error("Failed")
	}

	r2 := newRiya(2)
	r3 := newRiya(2)
	if r2.IsEqualTo(r3) != true {
		t.Error("Failed")
	}
}

func TestRiyaInt_IsGreaterThan(t *testing.T) {
	fourtyFive := newRiya(45)
	thirtySix := newRiya(36)

	if fourtyFive.IsGreaterThan(thirtySix) == true {
		t.Error("Failed")
	}
}
func TestRiyaInt_IsLessThan(t *testing.T) {
	twentyThree := newRiya(23)
	thirteen := newRiya(13)

	if twentyThree.IsLessThan(thirteen) != true {
		t.Error("Failed")
	}
}

func TestRiyaInt_Add(t *testing.T) {
	thiryFour := newRiya(34)
	if thiryFour.GetNative() != 34 {
		t.Error("Failed")
	}
	ten := newRiya(10)
	if ten.GetNative() != 10 {
		t.Error("Failed")
	}
	fourtyFour := thiryFour.Add(ten)
	if fourtyFour.GetNative() != 44 {
		t.Error("Failed")
	}
}

func TestRiyaInt_Subtract(t *testing.T) {
	thirtyFive := newRiya(35)
	ten := newRiya(10)
	twentyFive := thirtyFive.Subtract(ten)

	if ten.GetNative() != 10 {
		t.Error("Failed")
	}
	if thirtyFive.GetNative() != 35 {
		t.Error("Failed")
	}

	if twentyFive.GetNative() != 25 {
		t.Error("Failed")
	}
}
func TestRiyaInt_Multiply(t *testing.T) {
	thirtyFive := newRiya(35)
	if thirtyFive.GetNative() != 35 {
		t.Error("Failed")
	}

	ten := newRiya(10)
	if ten.GetNative() != 10 {
		t.Error("Failed")
	}
	threeHundredfifty := thirtyFive.Multiply(ten)
	if threeHundredfifty.GetNative() != 350 {
		t.Error("Failed")
	}
}
//
//func TestRiyaInt_DivideBy(t *testing.T) {
//	thirtyFive := newRiya(35)
//	if thirtyFive.num != 35 {
//		t.Error("Failed")
//	}
//	ten := newRiya(10)
//	if ten.num != 10 {
//		t.Error("Failed")
//	}
//	three, err := thirtyFive.DivideBy(ten)
//	if err != nil {
//		t.Error(err.Error())
//	}
//	if three.num != 3 {
//		t.Error("Failed")
//	}
//
//	fiftyFive := newRiya(55)
//	if fiftyFive.num != 55 {
//		t.Error("Failed")
//	}
//	zero := newRiya(0)
//	if zero.num != 0 {
//		t.Error("Failed")
//	}
//	_, err = fiftyFive.DivideBy(zero)
//	if err == nil {
//		t.Error("Expects an error")
//	}
//
//}
