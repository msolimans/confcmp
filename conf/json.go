package conf

import (
	. "../external"
	"errors"
)

//implements IConfig
type JsonConfig struct{
	p *Package
}

func (self *JsonConfig) UpdateVersion(done chan bool) {

	go func() {
		done <- true
	}()
}

func (self *JsonConfig) Unmarshal(in []byte, ch chan error) {

	ch <- errors.New("Not implemented")

}

func (self *JsonConfig) Marshal() (out []byte, err error) {
	return nil, errors.New("Not implemented")
}
