package comparison

import (
	"reflect"
	contracts "../contracts"
)

type SimpleCompare struct{

}

func newSimpleCompare() *SimpleCompare {
	return &SimpleCompare{}
}

func (self *SimpleCompare) Compare(config1, config2 contracts.IConfig, ch chan bool) {
	go func(){
		ch <- reflect.DeepEqual(config1, config2)
	}()
}
