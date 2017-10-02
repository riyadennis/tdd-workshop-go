package number

import "fmt"

type RiyaInt struct {
	num int
}

func newRiya(num int) (r RiyaInt) {
	r.num = num
	return r
}
func (r RiyaInt) GetNative() (int, error) {
	return r.num, nil
}

func (r RiyaInt) IsEqualTo(num int) (bool) {
	return num == r.num
}

func (r RiyaInt) IsGreaterThan(num int) (bool) {
	return r.num < num
}
func (r RiyaInt)IsLessThan(num int)(bool){
	return r.num > num
}

func (r RiyaInt) Add (num RiyaInt) (RiyaInt){
	r.num = r.num + num.num
	return r
}

func (r RiyaInt) Subtract(num RiyaInt)(RiyaInt){
	r.num = r.num - num.num
	return r
}

func (r RiyaInt) Multiply(num RiyaInt) (RiyaInt){
	r.num *= num.num
	return r
}

func (r RiyaInt) DivideBy(num RiyaInt) (RiyaInt, error){
	if num.num != 0 {
		r.num = r.num / num.num
		return r, nil
	}
	err := fmt.Errorf("Cant divide by zero")
	rR := RiyaInt{}
	return rR, err
}