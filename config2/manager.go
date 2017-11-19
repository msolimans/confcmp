package config2

import (
	contracts "../contracts"
	"errors"
)

type ConfigManager struct {
	configs []contracts.IConfig
	comparer contracts.ICompare//set strategy (which comparer should be used, easy to change in future)
}

func newManager(configs []contracts.IConfig, comparer contracts.ICompare) contracts.IManager {
	return &ConfigManager{configs,comparer }
}

func (self *ConfigManager) SetComparer(comparer contracts.ICompare){
	self.comparer = comparer
}


func (self *ConfigManager) UpdateVersions(done chan bool) {

	idone := make(chan bool)

	for _, p := range self.configs {
		go p.UpdateVersion(idone)
	}

	//wait for all of them here
	for range self.configs{
		<- idone
	}

	//emit final done here
	done <- true
}



func (self *ConfigManager) Compare() (error, bool) {


	done := make(chan bool, 1) //only 1 val and should be closed
	go self.UpdateVersions(done)

	if <-done {

		eq := make(chan bool)
		go self.comparer.Compare(self.configs[0], self.configs[1], eq)
		return nil, <-eq
	}

	return errors.New("error during parsing version string or updating version!"), false
}


func (self *ConfigManager) Diff() ( error, string){
	done := make(chan bool, 1) //only 1 val and should be closed
	go self.UpdateVersions(done)

	if <-done {

		diff := make(chan string)
		go self.comparer.Diff(self.configs[0], self.configs[1], diff)

		return nil, <-diff
	}

	return  errors.New("error during parsing version string or updating version!"), ""
}