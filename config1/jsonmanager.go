package config1

//new wrapped class/struct, in case we need to change it later to use JSON for example all we need to do is to change that wrapper
import (
	. "../external"
	"errors"
)

//Package wrapped into a new struct (composition)
//Decorator Pattern to extend its functionality
type JsonManager struct {
	packages []*Package
}

func newJsonManager() *JsonManager {
	return &JsonManager{make([]*Package, 2)} //2 packages, in case we need more in future
}

func (self *JsonManager) unmarshal(in []byte, out interface{}) (chan error) {
	ch := make(chan error)
	ch <- errors.New("Not Implemented")
	return ch
}


func (self *JsonManager) marshal(in interface{}) (out []byte, err error) {
	return nil, errors.New("Not implemented")
}

func (self *JsonManager) Compare(p1 []byte, p2 []byte) (bool, error) {

	 return false, errors.New("Not implemented yet")

}