package system

import (
	"log"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryUsage() [2]uint64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Error getting memory info: %v", err)
	}

	return [2]uint64{v.Used / 1024 / 1024, v.Total / 1024 / 1024}
}
