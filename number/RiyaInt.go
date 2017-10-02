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
	if integer.GetNative() == 0 {
		r.num = 0
		return nil, fmt.Errorf("Error cant divide by %v", integer.GetNative() )
	}
	divide  := new(RiyaInt)
	divide.num = r.num / integer.GetNative()
	return divide, nil
}
