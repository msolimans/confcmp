package comparison

import (
	"github.com/google/go-cmp/cmp"
	contracts "../contracts"
	"reflect"
)

type GoCompare struct{

}


func newGoCompare() *GoCompare {
	return &GoCompare{}
}


func (self *GoCompare) Compare(config1, config2 contracts.IConfig, ch chan bool) {
	go func(){
		ch <- reflect.DeepEqual(config1, config2)
	}()
}


func (self *GoCompare) Diff(config1, config2 contracts.IConfig, ch chan string) {
	go func(){
		ch <- cmp.Diff(config1, config2)
	}()
}
