package conf

import (
	. "../external"
	"../contracts"
	"gopkg.in/yaml.v2"
	"strings"
	"strconv"
)

//Implements IConfig
//Decorator Pattern to extend its functionality
type YamlConfig struct {
	Package *Package //favor composition over inheritance
}

func (self *YamlConfig) UpdateVersion(done chan bool) {

	go func() {
		if self.Package.VersionString != "" {
			vs := strings.Split(self.Package.VersionString, ".")
			count := 0
			for _, v := range vs {
				//Version is defiend as [3]int that is why I am taking the first 3 only
				if count > 2 { //0,1,2 only
					break
				}

				if vi, err := strconv.Atoi(v); err != nil {
					done <- false //err in val (not int)
					return        //return now
				} else {
					self.Package.Version[count] = vi
					count++
				}
			}
		}

		done <- true
	}()
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