package goedac

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	//"github.com/multiplay/go-edac/lib/edac"
)

var (
	mcPath = "/sys/devices/system/edac/mc"
)

// MemoryControllerStruct comment
type MemoryControllerStruct struct {
	Name string
}

// MemoryControllerStats comment
type MemoryControllerStats struct {
	SecondsSinceReset int64
	MCName            string
	SizeMB            int64
	UECount           int64
	UENoinfoCount     int64
	CECount           int64
	CENoinfoCount     int64
	SDRAMScrubRate    int64
	MaxLocation       string
}

// ChipSelectRowStruct comment
type ChipSelectRowStruct struct {
	Name string
}

// ChipSelectRowStats comment
type ChipSelectRowStats struct {
	Size                      int64  // size_mb
	EdacMode                  string // edac_mode
	DeviceType                string // dev_type
	MemoryType                string // mem_type
	CorrectableErrors         int64  // ce_count
	UncorrectableErrors       int64  // ue_count
	Chanel0CorrectableError   int64  // ch0_ce_count
	Chanel0UncorrectableError int64  // ch0_ue_count
	Chanel0DimLabel           string // ch0_dimm_label
	Chanel1CorrectableError   int64  // ch1_ce_count
	Chanel1UncorrectableError int64  // ch1_ue_count
	Chanel1DimLabel           string // ch1_dimm_label
}

func valueString(file string, addr *string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	*addr = strings.TrimSpace(string(d))
	return nil
}

func valueInt64(file string, addr *int64) error {
	var s string
	if err := valueString(file, &s); err != nil {
		return err
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*addr = v

	return nil
}
