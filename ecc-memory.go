package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/romchi/zabbix.ecc-memory/goedac"
)

type devices struct {
	DeviceName string `json:"{#DEVICE.NAME}"`
}

func main() {
	discoveryCommand := flag.NewFlagSet("discover", flag.ExitOnError)
	statsCommand := flag.NewFlagSet("stats", flag.ExitOnError)

	discoveryDeviceType := discoveryCommand.String("type", "", "device type {mc, csrow} (Required)")

	statsDeviceType := statsCommand.String("type", "", "device type {mc, csrow} (Required)")
	statsDeviceName := statsCommand.String("name", "", `Device "name" to get stats (Required)`)

	if len(os.Args) < 2 {
		fmt.Println("[discovery, stats] - required one command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "discovery":
		discoveryCommand.Parse(os.Args[2:])
	case "stats":
		statsCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if discoveryCommand.Parsed() {
		metricChoices := map[string]bool{"mc": true, "csrow": true}
		if _, validChoice := metricChoices[*discoveryDeviceType]; !validChoice {
			discoveryCommand.PrintDefaults()
			os.Exit(1)
		}
		switch *discoveryDeviceType {
		case "mc":
			discoveryMemControllers()
		case "csrow":
			discoveryChipSelectRows()
		default:
			discoveryCommand.PrintDefaults()
			os.Exit(1)
		}
	}

	if statsCommand.Parsed() {
		metricChoices := map[string]bool{"mc": true, "csrow": true}
		if _, validChoice := metricChoices[*statsDeviceType]; !validChoice {
			statsCommand.PrintDefaults()
			os.Exit(0)
		}
		if len(*statsDeviceName) < 1 {
			statsCommand.PrintDefaults()
			os.Exit(0)
		}
		switch *statsDeviceType {
		case "mc":
			statsMemController(*statsDeviceName)
		case "csrow":
			statsChipSelectRow(*statsDeviceName)
		default:
			statsCommand.PrintDefaults()
			os.Exit(0)
		}
	}
}

func discoveryMemControllers() {
	result := []devices{}

	memControllers, err := goedac.MemoryControllersList()
	if err != nil {
		log.Fatal(err)
	}

	for _, controller := range memControllers {
		result = append(result, devices{DeviceName: controller})
	}

	r, _ := json.Marshal(result)

	fmt.Println(string(r))
}

func discoveryChipSelectRows() {
	result := []devices{}

	csrows, err := goedac.ChipSelectRowList()
	if err != nil {
		log.Fatal(err)
	}

	for _, csrow := range csrows {
		result = append(result, devices{DeviceName: csrow})
	}

	r, _ := json.Marshal(result)

	fmt.Println(string(r))
}

func statsMemController(mcName string) {
	mcs, err := goedac.MemoryControllers()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range mcs {
		if mcName == c.Name {
			i, err := c.Stats()
			if err != nil {
				log.Fatal(err)
			}
			ji, _ := json.Marshal(i)
			fmt.Print(string(ji))
		}
	}
}

func statsChipSelectRow(csrowName string) {
	csrows, err := goedac.ChipSelectRows()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range csrows {
		if csrowName == c.Name {
			i, err := c.Stats()
			if err != nil {
				log.Fatal(err)
			}
			ji, _ := json.Marshal(i)
			fmt.Print(string(ji))
		}
	}
}
