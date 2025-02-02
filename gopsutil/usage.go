package gopsutil

import (
	"go-profiler/database"

	"github.com/shirou/gopsutil/v3/process"
)

func GetProcessesInfo() ([]database.Process, error) {

	// get process info
	p, _ := process.Processes()
	process := make([]database.Process, 0)
	for _, proc := range p {
		name, err1 := proc.Name()
		cpuUsage, err2 := proc.CPUPercent()
		memUsage, err3 := proc.MemoryPercent()
		createTime, err4 := proc.CreateTime()
		if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
			process = append(process, database.Process{
				Name:       name,
				CPUUsage:   cpuUsage,
				Memory:     memUsage,
				ProcessId:  uint32(proc.Pid),
				CreateTime: createTime,
			})
		}
	}
	return process, nil

}
