package config2

import (
	. "../external"
	"../contracts"
	"gopkg.in/yaml.v2"
)


//implements IConfig
//Decorator Pattern to extend its functionality
type YamlConfig struct{
	p *Package //favor composition over inheritance
}


func (self *YamlConfig) Unmarshal(in []byte) (chan error) {
	ch := make(chan error)
	go func() {
		if self.p ==  nil{
			self.p = &Package{}
		}
		ch <- yaml.Unmarshal(in, self.p)
	}()

	return ch
}

func (self *YamlConfig) Marshal() (out []byte, err error) {
	return yaml.Marshal(self.p)
}

func (self *YamlConfig) Get() contracts.IConfig{
	return self
}