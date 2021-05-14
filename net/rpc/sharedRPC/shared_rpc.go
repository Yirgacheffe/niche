package sharedRPC

type MyFloats struct {
	A1, A2 float64
}

type MyInterface interface {
	Multiply(args *MyFloats, reply *float64) error
	Power(args *MyFloats, reply *float64) error
}
