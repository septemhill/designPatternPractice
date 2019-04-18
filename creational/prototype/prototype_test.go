package prototype

import "errors"

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return nil
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(n int) (ItemInfoGetter, error) {
	return nil, errors.New("not implemented yet")
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return ""
}

func GetShirtsCloner()  ShirtCloner {
	
}