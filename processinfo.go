package  systeminfo

import (
	"github.com/mitchellh/go-ps"
)

func GetProcessPid(ProcessName string) (ProcessId int) {
	p, _ := ps.Processes()
	var length int = len(p)
	var flag int
	for flag < length {
		if p[flag].Executable() == ProcessName {
			return p[flag].Pid()
		} else {
			flag++
		}
	}
	return 0
}

func GetProcessStatus(ProcessName string) (status bool) {
	p, _ := ps.Processes()
	var length int = len(p)
	var flag int = 0
	for flag < length {
		if p[flag].Executable() == ProcessName {
			return true
		} else {
			flag++
		}
	}
	return false
}

