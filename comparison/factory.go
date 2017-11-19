package comparison

//Factory-Method
func GetInstance() *SimpleCompare{
	return newSimpleCompare()
}