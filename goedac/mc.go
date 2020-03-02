package goedac

import (
	"os"
	"path/filepath"
)

// MemoryControllersList list of memory controllers
func MemoryControllersList() ([]string, error) {
	// TODO: find regexp which match mc123
	files, err := filepath.Glob(filepath.Join(mcPath, "mc[0-9]"))

	if _, err = os.Stat(mcPath); os.IsNotExist(err) {
		return []string{}, nil
	}

	if err != nil {
		return nil, err
	}

	memoryControllers := make([]string, len(files))
	for ctr, name := range files {
		memoryControllers[ctr] = filepath.Base(name)
	}

	return memoryControllers, nil
}

// MemoryControllers comment
func MemoryControllers() ([]MemoryControllerStruct, error) {
	controllerList, err := MemoryControllersList()
	if err != nil {
		return nil, err
	}

	memoryControllers := make([]MemoryControllerStruct, len(controllerList))
	for ctr, num := range controllerList {
		memoryControllers[ctr] = MemoryControllerStruct{Name: filepath.Base(num)}
	}

	return memoryControllers, nil
}

// Stats comment
func (mc MemoryControllerStruct) Stats() (*MemoryControllerStats, error) {
	dir := filepath.Join(mcPath, mc.Name)

	f, err := os.Open(dir)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	var files []string
	if files, err = f.Readdirnames(0); err != nil {
		return nil, err
	}

	mcStats := &MemoryControllerStats{}

	for _, f := range files {
		p := filepath.Join(dir, f)
		switch f {
		case "seconds_since_reset":
			if err = valueInt64(p, &mcStats.SecondsSinceReset); err != nil {
				return nil, err
			}
		case "ue_count":
			if err = valueInt64(p, &mcStats.UncorrectableErrors); err != nil {
				return nil, err
			}
		case "ue_noinfo_count":
			if err = valueInt64(p, &mcStats.UncorrectableErrorsNoinfo); err != nil {
				return nil, err
			}
		case "ce_count":
			if err = valueInt64(p, &mcStats.CorrectableErrors); err != nil {
				return nil, err
			}
		case "ce_noinfo_count":
			if err = valueInt64(p, &mcStats.CorrectableErrorsNoinfo); err != nil {
				return nil, err
			}
		case "size_mb":
			if err = valueInt64(p, &mcStats.SizeMB); err != nil {
				return nil, err
			}
		case "sdram_scrub_rate":
			if err = valueInt64(p, &mcStats.SDRAMScrubRate); err != nil {
				return nil, err
			}
		case "mc_name":
			if err = valueString(p, &mcStats.MCName); err != nil {
				return nil, err
			}
		case "max_location":
			if err = valueString(p, &mcStats.MaxLocation); err != nil {
				return nil, err
			}
		}
	}
	return mcStats, nil
}
