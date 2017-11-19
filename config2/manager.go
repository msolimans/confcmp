package config2

import (
//	"errors"
	contracts "../contracts"
)

type ConfigManager struct {
	configs []contracts.IConfig
	comparer contracts.ICompare//set strategy (which comparer should be used, easy to change in future)
}

func newManager(configs []contracts.IConfig, comparer contracts.ICompare) IManager {
	return &ConfigManager{configs,comparer }
}

func (self *ConfigManager) SetComparer(comparer contracts.ICompare){
	self.comparer = comparer
}

func (self *ConfigManager) Compare(p1 []byte, p2 []byte) (error, bool) {

	ch1 := make(chan error, 1)
	ch2 := make(chan error, 1)


	//first operation
	ch1 = self.configs[0].Unmarshal(p1)
	ch2 = self.configs[1].Unmarshal(p2)

	//handling errors during unmarshaling
	if err := <-ch1; err != nil {
		return err, false
	}

	if err := <-ch2; err != nil {
		return err, false
	}


	eq := make(chan bool)
	go self.comparer.Compare(self.configs[0], self.configs[1], eq)
	return nil, <- eq


	//return errors.New("error during parsing version string or updating version!"), false

}