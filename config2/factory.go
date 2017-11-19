package config2

import (
	"../contracts"
	"../comparison"
	. "../conf"
)

//Factory Method Pattern (Inject different Manager here)
func GetInstance(p1, p2 []byte) (error, contracts.IManager) {
	comparer := comparison.GetInstance()
	con := make([]contracts.IConfig, 2)
	con[0] = new(YamlConfig)//use or inject differnt configs here (JsonConfig for example)
	con[1] = new(YamlConfig)


	ch1 := make(chan error, 1)
	ch2 := make(chan error, 1)

	//first operation
	go con[0].Unmarshal(p1, ch1)
	go con[1].Unmarshal(p2, ch2)

	//handling errors during unmarshaling
	if err := <-ch1; err != nil {
		return err, nil
	}

	if err := <-ch2; err != nil {
		return err, nil
	}

	return nil, newManager(con, comparer)

}
