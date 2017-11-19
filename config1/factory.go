package config1

import (
	"sync"
	contracts "../contracts"
)

//this should initialize the wrapper and returns a behaviour (interface) based on some sort of Dependency Injection
var instance contracts.IMarshal
var mutex sync.Mutex

//Factory Method with Singleton Pattern
func GetInstance() contracts.IMarshal{

	//check-lock-check technique
	if instance == nil{
		//lock the executuion to make it thread-safe
		mutex.Lock()
		//don't forget to defere unlock call
		defer mutex.Unlock()

		//check again to make sure it is still nil inside lock block
		if instance == nil{
			instance = newPackageManager()
		}
	}

	return instance
}