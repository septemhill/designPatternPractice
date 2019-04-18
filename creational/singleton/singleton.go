package singleton

import (
	"sync"
)

type singleton struct {
	count int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
