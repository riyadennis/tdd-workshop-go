package number

import "fmt"

type RiyaInt struct {
	num int
}

func newRiya(num int) (*RiyaInt) {
	return &RiyaInt{
		num,
	}
}
func (r *RiyaInt) GetNative() (int) {
	return r.num
}

func (r *RiyaInt) IsEqualTo(integer MarksInteger) (bool) {
	return integer.GetNative() == r.num
}

func (r *RiyaInt) IsGreaterThan(integer MarksInteger) (bool) {
	return r.num < integer.GetNative()
}
func (r *RiyaInt)IsLessThan(integer MarksInteger)(bool){
	return r.num > integer.GetNative()
}

func (r *RiyaInt) Add (integer MarksInteger) (MarksInteger){
	r.num = r.num + integer.GetNative()
	return r
}

func (r *RiyaInt) Subtract(integer MarksInteger)(MarksInteger){
	r.num = r.num - integer.GetNative()
	return r
}

func (r *RiyaInt) Multiply(integer MarksInteger) (MarksInteger){
	r.num *= integer.GetNative()
	return r
}

func (r *RiyaInt) DivideBy(integer MarksInteger) (MarksInteger, error){
	if integer.GetNative() != 0 {
		r.num = r.num / integer.GetNative()
		return r, nil
	}
	err := fmt.Errorf("Cant divide by zero")
	rR := &RiyaInt{}
	return rR, err
}
