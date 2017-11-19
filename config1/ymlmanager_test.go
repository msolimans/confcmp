package config1

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var similar1 = `
kind: "set"
versionstring: "3.2.1"
authors:
- name: Muhammad
  email: m.solimanz@hotmail.com
- name: sam
  email: sam@opentable.com
`

var similar2 = `
kind: "set"
versionstring: "3.2.1"
authors:
- name: Muhammad
  email: m.solimanz@hotmail.com
- name: sam
  email: sam@opentable.com
`

func TestYmlManager_Compare(t *testing.T) {

	if err, res := GetInstance().Compare([]byte(similar1), []byte(similar2)); err != nil {
		assert.Fail(t, "Error while comparing yaml data")
	} else {
		fmt.Printf("result=%s", res)
		assert.True(t, res == true, "Test Failure, both %s and %s are similar however it resulted in false", similar1, similar2)
	}
}