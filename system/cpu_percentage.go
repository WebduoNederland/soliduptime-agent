package system

import (
	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCPUPercentage() uint64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatalf("Error getting CPU percentage info: %v", err)
	}

	return uint64(math.Round(percent[0]))
}
