package config

import (
	"testing"
)

const pathtofile = "/users/ariefdarmawan/Temp/config.json"

func TestCfgWrite(t *testing.T) {
	SetConfigFile(pathtofile)
	Set("FullName", "Arief Darmawan")
	if e := Write(); e != nil {
		t.Error(e.Error())
	}
}

func TestCfgGet(t *testing.T) {
	SetConfigFile(pathtofile)
	s := Get("FullName").(string)
	if s != "Arief Darmawan" {
		t.Error("Unable to read value. Expected 'Arief Darmawan' got '" + s + "'")
	}
}
