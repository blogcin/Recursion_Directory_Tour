package Util

import (
	"os"
	"runtime"
)

type Util struct { }

func (util *Util) CheckIsExist(file string) bool {
	status := false
	if _, err := os.Stat(file); os.IsNotExist(err) {
		status = false
	} else {
		status = true
	}

	return status

}

func (util *Util) SupportMultiPlatform() string {
	var temp string

	if runtime.GOOS == "linux" {
		temp = "/";
	} else if runtime.GOOS == "windows" {
		temp = "\\";
	} else {
		temp = "UnKnown"
	}

	return temp
}