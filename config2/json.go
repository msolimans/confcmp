package config2

import (
	. "../external"
	"errors"
)
type JsonConfig struct{
	p *Package
}


func (self *JsonConfig) unmarshal(in []byte) (chan error) {
	ch := make(chan error)
	ch <- errors.New("Not implemented")
	return ch
}

func (self *JsonConfig) marshal() (out []byte, err error) {
	return nil, errors.New("Not implemented")
}
