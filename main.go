package main

import (
	"fmt"
	"log"
	"./config1"
	"./config2"
)

var pData1 = `
kind: "set"
versionstring: "3.2.1"
authors:
- name: Muhammad
  email: m.solimanz@hotmail.com
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
	manager := config1.GetInstance()
	if err, res := manager.Compare([]byte(pData1), []byte(pData2)); err != nil {
		log.Fatalf("error occured %v", err)
	} else {
		fmt.Printf("result=%s", res)
	}


	//Solution2

	manager2 := config2.GetInstance()
	if err, res := manager2.Compare([]byte(pData1), []byte(pData2)); err != nil {
		log.Fatalf("error occured %v", err)
	} else {
		fmt.Printf("result=%s", res)
	}




}
