package main

import (
	"flag"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/webduonederland/soliduptime-agent/api"
	"github.com/webduonederland/soliduptime-agent/system"
)

var apiKey string

func init() {
	flag.StringVar(&apiKey, "api-key", "", "Your SolidUptime server API key")
}

func main() {
	flag.Parse()

	if apiKey == "" {
		log.Fatal("API key is required. Please provide a valid API key using the --api-key flag.")
	}

	scheduler := gocron.NewScheduler(time.Local)

	_, err := scheduler.Every(1).Minute().Do(sendSystemData)
	if err != nil {
		log.Fatalf("Failed to schedule job: %v", err)
	}

	scheduler.StartBlocking()
}

func sendSystemData() {
	memory := system.GetMemoryUsage()
	disk := system.GetDiskUsage()
	cpuPercentage := system.GetCPUPercentage()
	osName := system.GetOSName()

	data := api.SystemData{
		UsedRAM:       memory[0],
		TotalRAM:      memory[1],
		UsedDisk:      disk[0],
		TotalDisk:     disk[1],
		CpuPercentage: cpuPercentage,
		OsName:        osName,
	}

	err := api.SendData(data, apiKey)
	if err != nil {
		log.Fatalf("Failed to send data: %v", err)
	}
}
