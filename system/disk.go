package system

import (
	"log"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskUsage() [2]uint64 {
	usage, err := disk.Usage("/")
	if err != nil {
		log.Fatalf("Error getting disk usage: %v", err)
	}

	return [2]uint64{usage.Used / 1024 / 1024, usage.Total / 1024 / 1024}
}
