package config1

import (
	contracts "../contracts"
)

//this should initialize the wrapper and returns a behaviour (interface) based on some sort of Dependency Injection
//Factory Method
func GetInstance(p1, p2 []byte) (error, contracts.IMarshal) {

	return newPackageManager(p1, p2)

}