package strategy

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(lval, rval int) int {
	return o.Operator.Apply(lval, rval)
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}
