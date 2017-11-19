package config1

//new wrapped class/struct, in case we need to change it later to use JSON for example all we need to do is to change that wrapper
import (
	. "../conf"
	"../contracts"
	"../comparison"
	"reflect"
	"errors"
)

//Package wrapped into a new struct (composition)
//Decorator Pattern to extend its functionality
type YmlManager struct {
	packages []contracts.IConfig
	comparer contracts.ICompare //set strategy (which comparer should be used, easy to change in future)
}

func newPackageManager() *YmlManager {
	con := make([]contracts.IConfig, 2)
	con[0] = new(YamlConfig)
	con[1] = new(YamlConfig)
	return &YmlManager{con, comparison.GetInstance()} //2 packages, in case we need more in future

}

func (self *YmlManager) isEqual(ch chan bool) {
	go func() {
		ch <- reflect.DeepEqual(self.packages[0], self.packages[1])
	}()
}

func (self *YmlManager) UpdateVersions(done chan bool) {

	idone := make(chan bool)

	for _, p := range self.packages {
		go p.UpdateVersion(idone)
	}

	//wait for all of them here
	for range self.packages{
		<- idone
	}

	//emit final done here
	done <- true
}

func (self *YmlManager) Compare(p1 []byte, p2 []byte) (error, bool) {

	ch1 := make(chan error, 1)
	ch2 := make(chan error, 1)

	//first operation
	go self.packages[0].Unmarshal(p1, ch1)
	go self.packages[1].Unmarshal(p2, ch2)

	//handling errors during unmarshaling
	if err := <-ch1; err != nil {
		return err, false
	}

	if err := <-ch2; err != nil {
		return err, false
	}

	done := make(chan bool, 1) //only 1 val and should be closed
	go self.UpdateVersions(done)

	if <-done {
		cdone := make(chan bool, 1)
		go self.comparer.Compare(self.packages[0], self.packages[1], cdone)

		return nil, <-cdone
	}

	return errors.New("error during parsing version string or updating version!"), false
}

func (self *YmlManager) Diff(p1 []byte, p2 []byte) (error, string) {
	ch1 := make(chan error, 1)
	ch2 := make(chan error, 1)

	//first operation
	go self.packages[0].Unmarshal(p1, ch1)
	go self.packages[1].Unmarshal(p2, ch2)

	//handling errors during unmarshaling
	if err := <-ch1; err != nil {
		return err, ""
	}

	if err := <-ch2; err != nil {
		return err, ""
	}

	diff := make(chan string)
	go self.comparer.Diff(self.packages[0], self.packages[1], diff)

	return nil, <-diff

}
