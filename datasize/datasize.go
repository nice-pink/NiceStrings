package datasize

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Bytes string = "Bytes"
	KB    string = "KB"
	MB    string = "MB"
	GB    string = "GB"
	TB    string = "TB"
	PB    string = "PB"
)

const (
	MultiBytes int64 = 1
	MultiKB    int64 = 1024
	MultiMB    int64 = MultiKB * 1024
	MultiGB    int64 = MultiMB * 1024
	MultiTB    int64 = MultiGB * 1024
	MultiPB    int64 = MultiTB * 1024
)

type DataSize struct {
	Value int64
	Unit  string
}

func (d *DataSize) String() string {
	return fmt.Sprintf("%d%s", d.Value, d.Unit)
}

func (d *DataSize) ToBytes() int64 {
	if d.Unit == KB {
		return d.Value * MultiKB
	}
	if d.Unit == MB {
		return d.Value * MultiMB
	}
	if d.Unit == GB {
		return d.Value * MultiGB
	}
	if d.Unit == TB {
		return d.Value * MultiTB
	}
	if d.Unit == PB {
		return d.Value * MultiPB
	}
	return d.Value
}

func ToBytes(s string) (int64, error) {
	ds, err := FromString(s)
	if err != nil {
		return 0, err
	}
	return ds.ToBytes(), nil
}

func FromString(s string) (*DataSize, error) {
	if strings.HasSuffix(s, KB) {
		return ForSuffix(s, KB)
	}
	if strings.HasSuffix(s, MB) {
		return ForSuffix(s, MB)
	}
	if strings.HasSuffix(s, GB) {
		return ForSuffix(s, GB)
	}
	if strings.HasSuffix(s, TB) {
		return ForSuffix(s, TB)
	}
	if strings.HasSuffix(s, PB) {
		return ForSuffix(s, PB)
	}
	return nil, fmt.Errorf("invalid data size: %s", s)
}

func ForSuffix(s, suffix string) (*DataSize, error) {
	value := strings.TrimSuffix(s, suffix)
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return &DataSize{Value: intValue, Unit: suffix}, nil
}
