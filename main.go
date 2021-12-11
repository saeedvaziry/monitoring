package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/viper"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(frequency()).Minutes().Do(sendStats)
	<-s.Start()
}

// get stats url
func frequency() uint64 {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	configError := viper.ReadInConfig()
	if configError != nil {
		panic(fmt.Errorf("Cannot read the config.json file"))
	}
	return viper.GetUint64("frequency")
}

// get stats url
func statsUrl() string {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	configError := viper.ReadInConfig()
	if configError != nil {
		panic(fmt.Errorf("Cannot read the config.json file"))
	}
	return viper.GetString("url")
}

// get memory usage
func memoryUsage() string {
	m, _ := mem.VirtualMemory()
	return fmt.Sprintf("%f", m.UsedPercent)
}

// get cpu usage
func cpuUsage() string {
	c, _ := cpu.Percent(0, false)
	return fmt.Sprintf("%f", c)
}

// get disk usage
func diskUsage() string {
	d, _ := disk.Usage("/")
	return fmt.Sprintf("%f", d.UsedPercent)
}

// send stats to server
func sendStats() {
	values := map[string]string{"memory": memoryUsage(), "cpu": cpuUsage(), "disk": diskUsage()}
	jsonValue, _ := json.Marshal(values)
	_, err := http.Post(statsUrl(), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(fmt.Errorf("Send stats error"))
	}
}

func test() {
	fmt.Println("test")
}
