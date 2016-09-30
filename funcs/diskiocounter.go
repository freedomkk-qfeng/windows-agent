package funcs

import (
	"github.com/StackExchange/wmi"
)

type Win32_PerfFormattedData struct {
	AvgDiskSecPerRead_Base  uint32
	AvgDiskSecPerWrite_Base uint32
	DiskReadBytesPerSec     uint64
	DiskReadsPerSec         uint32
	DiskWriteBytesPerSec    uint64
	DiskWritesPerSec        uint32
	Name                    string
}

type Win32_PerfFormattedData_IDLE struct {
	Name            string
	PercentIdleTime uint64
}

type diskIOCounter struct {
	Msec_Read      uint32
	Msec_Write     uint32
	Read_Bytes     uint64
	Read_Requests  uint32
	Write_Bytes    uint64
	Write_Requests uint32
	Util           uint64
}

func PerfFormattedData() ([]Win32_PerfFormattedData, error) {

	var dst []Win32_PerfFormattedData
	Query := `SELECT 
				AvgDiskSecPerRead_Base,
				AvgDiskSecPerWrite_Base,
				DiskReadBytesPerSec,
				DiskReadsPerSec,
				DiskWriteBytesPerSec,
				DiskWritesPerSec,
				Name
				FROM Win32_PerfRawData_PerfDisk_PhysicalDisk`
	err := wmi.Query(Query, &dst)

	return dst, err
}

func PerfFormattedData_IDLE() ([]Win32_PerfFormattedData_IDLE, error) {

	var dst []Win32_PerfFormattedData_IDLE

	err := wmi.Query("SELECT PercentIdleTime FROM Win32_PerfFormattedData_PerfDisk_PhysicalDisk ", &dst)

	return dst, err
}

func IOCounters() (map[string]diskIOCounter, error) {
	ret := make(map[string]diskIOCounter, 0)
	dst, err := PerfFormattedData()
	if err == nil {
		for _, d := range dst {
			if d.Name == "_Total" { // not get _Total
				continue
			}
			ret[d.Name] = diskIOCounter{
				Msec_Read:      d.AvgDiskSecPerRead_Base,
				Msec_Write:     d.AvgDiskSecPerWrite_Base,
				Read_Bytes:     d.DiskReadBytesPerSec,
				Read_Requests:  d.DiskReadsPerSec,
				Write_Bytes:    d.DiskWriteBytesPerSec,
				Write_Requests: d.DiskWritesPerSec,
				Util:           0,
			}
		}
	}
	dstIdle, err := PerfFormattedData_IDLE()
	if err == nil {
		for _, dd := range dstIdle {
			if dd.Name == "_Total" {
				continue
			}
			result := ret[dd.Name]
			result.Util = dd.PercentIdleTime
			ret[dd.Name] = result
		}
	}
	return ret, err
}
