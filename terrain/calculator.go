package main

type Calculator struct {
	operation OperationFunc
	v1 float64
	v2 float64
}

type OperationFunc func(float64,float64) float64

func Add(a float64, b float64) float64{
	return a+b
}
func (c Calculator) Calculate(operation OperationFunc) float64 {
	return operation(c.v1, c.v2)
}
func main(){
	var c Calculator
	c.v1=10
	c.v2=20
	c.Calculate(Add)
}