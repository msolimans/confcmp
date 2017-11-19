package comparison

import (
	"../contracts"
	"sync"
)

var instance contracts.ICompare
var mutex sync.Mutex

//Factory-Method, Singleton
func GetInstance() contracts.ICompare {

	//check-lock-check technique
	if instance == nil{
		//lock the executuion to make it thread-safe
		mutex.Lock()
		//don't forget to defere unlock call
		defer mutex.Unlock()

		//check again to make sure it is still nil inside lock block
		if instance == nil{
			instance = newGoCompare()
		}
	}

	return instance

}