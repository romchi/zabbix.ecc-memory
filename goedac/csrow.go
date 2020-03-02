package goedac

import (
	"os"
	"path/filepath"
	"strings"
)

// ChipSelectRowList comment
func ChipSelectRowList() ([]string, error) {

	mc, err := MemoryControllersList()
	if err != nil {
		return nil, err
	}

	chipSelectRows := []string{}

	for _, name := range mc {
		// TODO: find regexp which match csrow123
		files, err := filepath.Glob(filepath.Join(mcPath, name, "csrow[0-9]"))
		if err != nil {
			return nil, err
		}

		for _, csname := range files {
			chipSelectRows = append(chipSelectRows, name+"."+filepath.Base(csname))
		}
	}

	return chipSelectRows, nil
}

// ChipSelectRows comment
func ChipSelectRows() ([]ChipSelectRowStruct, error) {
	csrList, err := ChipSelectRowList()
	if err != nil {
		return nil, err
	}

	chipSelectRows := make([]ChipSelectRowStruct, len(csrList))
	for ctr, num := range csrList {
		chipSelectRows[ctr] = ChipSelectRowStruct{Name: filepath.Base(num)}
	}

	return chipSelectRows, nil
}

// Stats comment
func (csr ChipSelectRowStruct) Stats() (*ChipSelectRowStats, error) {
	csrDir := strings.Replace(csr.Name, ".", "/", -1)
	dir := filepath.Join(mcPath, csrDir)

	f, err := os.Open(dir)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	var files []string
	if files, err = f.Readdirnames(0); err != nil {
		return nil, err
	}

	csrStats := &ChipSelectRowStats{}

	for _, f := range files {
		p := filepath.Join(dir, f)
		switch f {
		case "size_mb":
			if err = valueInt64(p, &csrStats.Size); err != nil {
				return nil, err
			}
		case "edac_mode":
			if err = valueString(p, &csrStats.EdacMode); err != nil {
				return nil, err
			}
		case "dev_type":
			if err = valueString(p, &csrStats.DeviceType); err != nil {
				return nil, err
			}
		case "mem_type":
			if err = valueString(p, &csrStats.MemoryType); err != nil {
				return nil, err
			}
		case "ce_count":
			if err = valueInt64(p, &csrStats.CorrectableErrors); err != nil {
				return nil, err
			}
		case "ue_count":
			if err = valueInt64(p, &csrStats.UncorrectableErrors); err != nil {
				return nil, err
			}
		}
	}
	return csrStats, nil
}
