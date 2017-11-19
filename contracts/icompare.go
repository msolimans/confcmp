package contracts

//Strategy pattern, different algorithms for comparisons
type ICompare interface{

	Compare(config1, config2 IConfig, ch chan bool)
	Diff(config1, config2 IConfig, ch chan string)

}
