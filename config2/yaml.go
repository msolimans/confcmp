package config2

import (
	. "../external"
	"../contracts"
	"gopkg.in/yaml.v2"
)

//Implements IConfig
//Decorator Pattern to extend its functionality
type YamlConfig struct {
	Package *Package //favor composition over inheritance
}


func (self *YamlConfig) Unmarshal(in []byte, ch chan error) {

	go func() {
		if self.Package ==  nil{
			self.Package = &Package{}
		}
		ch <- yaml.Unmarshal(in, self.Package)
	}()

}

func (self *YamlConfig) Marshal() (out []byte, err error) {
	return yaml.Marshal(self.Package)
}

func (self *YamlConfig) Get() contracts.IConfig {
	return self
}