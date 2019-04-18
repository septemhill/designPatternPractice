package strategy

import "testing"

func TestStrategy(t *testing.T) {
	add := Operation{Addition{}}
	res := add.Operate(2, 4)
	if res != 6 {
		t.Errorf("got errors: expect 6 but :%v", res)
	}

	mul := Operation{Multiplication{}}

	res = mul.Operate(5, 6)
	if res != 30 {
		t.Errorf("got errors: expect 30, but: %v", res)
	}
}
