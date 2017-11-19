package config2

import (
	"sync"
	"../contracts"
	"../comparison"
	. "../conf"
)

//this should initialize the wrapper and returns a behaviour (interface) based on some sort of Dependency Injection
var instance contracts.IManager
var mutex sync.Mutex

//Factory Method Pattern - Singleton Pattern
func GetInstance() contracts.IManager{

	//check-lock-check technique
	if instance == nil{
		//lock the executuion to make it thread-safe
		mutex.Lock()
		//don't forget to defere unlock call
		defer mutex.Unlock()

		//check again to make sure it is still nil inside lock block
		if instance == nil{

			c := comparison.GetInstance()
			con := make([]contracts.IConfig,2)
			con[0] = new(YamlConfig)
			con[1] = new(YamlConfig)
			instance = newManager(con, c)
		}
	}

	return instance
}