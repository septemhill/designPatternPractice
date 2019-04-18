package singleton

import "testing"

func TestSingleton(t *testing.T) {
	inst := GetInstance()

	n1 := inst.AddOne()
	if n1 != 1 {
		t.Errorf("got errors: expect value = 1, but get %v", n1)
	}

	n2 := inst.AddOne()
	if n2 != 2 {
		t.Errorf("got errors: expect value = 2, but get %v", n2)
	}
}

func BenchmarkSingleton(b *testing.B) {
	inst := GetInstance()

	for i := 0; i < b.N; i++ {
		inst.AddOne()
	}
}
