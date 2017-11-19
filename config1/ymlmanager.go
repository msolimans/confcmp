package config1

//new wrapped class/struct, in case we need to change it later to use JSON for example all we need to do is to change that wrapper
import (
	. "../external"
	"gopkg.in/yaml.v2"
	"reflect"
	"strings"
	"strconv"
	"errors"
	"time"
)

//Package wrapped into a new struct (composition)
//Decorator Pattern to extend its functionality
type YmlManager struct {
	packages []*Package
}

func newPackageManager() *YmlManager {
	return &YmlManager{make([]*Package, 2)} //2 packages, in case we need more in future
}

func (self *YmlManager) Unmarshal(in []byte, out interface{}) (chan error) {
	ch := make(chan error)
	go func() {
		ch <- yaml.Unmarshal(in, out)
	}()

	return ch
}

func (self *YmlManager) Unmarshal1(in []byte, out interface{}) (chan error) {
	ch := make(chan error)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		ch <- yaml.Unmarshal(in, out)
	}()

	return ch
}

func (self *YmlManager) Marshal(in interface{}) (out []byte, err error) {
	return yaml.Marshal(in)
}

func (self *YmlManager) isEqual(ch chan bool) {
	go func() {
		ch <- reflect.DeepEqual(self.packages[0], self.packages[1])
	}()
}

func (self *YmlManager) updateVersions(done chan bool) {
	for _, p := range self.packages {
		if p.VersionString != "" {
			vs := strings.Split(p.VersionString, ".")
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
					p.Version[count] = vi
					count++
				}
			}
		}
	}

	done <- true
}

func (self *YmlManager) Compare(p1 []byte, p2 []byte) (error, bool) {

	ch1 := make(chan error, 1)
	ch2 := make(chan error, 1)

	//first operation
	ch1 = self.Unmarshal(p1, &(self.packages[0]))

	//second operation
	ch2 = self.Unmarshal1(p2, &(self.packages[1]))

	//handling errors during unmarshaling
	if err := <-ch1; err != nil {
		return err, false
	}

	if err := <-ch2; err != nil {
		return err, false
	}

	done := make(chan bool, 1) //only 1 val and should be closed
	go self.updateVersions(done)

	if update := <-done; update {
		eqCh := make(chan bool, 1)
		go self.isEqual(eqCh)
		return nil, <-eqCh //waiting for equa channel to populate and return its result
	}

	return errors.New("error during parsing version string or updating version!"), false

}
