package main

import (
	"fmt"
	"log"
	"./config1"
	"./config2"
)

var pData1 = `
kind: "st"
versionstring: "3.2.1"
authors:
- name: Judson
  email: judon@opentable.com
- name: sam
  email: sam@opentable.com
`

var pData2 = `
kind: "set"
versionstring: "3.2.1"
authors:
- name: Muhammad
  email: m.solimanz@hotmail.com
- name: sam
  email: sam@opentable.com
`

func main() {

	//Solution1
	if err, manager := config1.GetInstance([]byte(pData1), []byte(pData2)); err != nil{
		log.Fatalf("error occured during initializing manager %v", err)
	}else{
		if err, res := manager.Compare(); err != nil {
			log.Fatalf("error occured %v", err)
		} else {
			fmt.Printf("\nresult=%s", res)
		}

	}



	//Solution2
	if err, manager2 := config2.GetInstance([]byte(pData1), []byte(pData2)); err != nil{
		log.Fatalf("error occured %v", err)
	}else{
		if err, res := manager2.Compare(); err != nil {
			log.Fatalf("error occured %v", err)
		} else {
			fmt.Printf("\nresult=%s", res)
		}

		if err, res := manager2.Diff(); err != nil {
			log.Fatalf("error occured %v", err)
		} else {
			fmt.Printf("\ndifference=%s", res)
		}

	}


}
