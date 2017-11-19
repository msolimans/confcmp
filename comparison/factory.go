package comparison

import "../contracts"
//Factory-Method
func GetInstance() contracts.ICompare {
	return newGoCompare()
}