package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SystemData struct {
	UsedRAM       uint64 `json:"memory_used"`
	TotalRAM      uint64 `json:"memory_total"`
	UsedDisk      uint64 `json:"disk_used"`
	TotalDisk     uint64 `json:"disk_total"`
	CpuPercentage uint64 `json:"cpu_percentage"`
	OsName        string `json:"os"`
}

func SendData(data SystemData, apiKey string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://soliduptime.com/api/v1/servers/monitor", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned status: %d", resp.StatusCode)
	} else {
		fmt.Println("System data sent!")
	}

	return nil
}
