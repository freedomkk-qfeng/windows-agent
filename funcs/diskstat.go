package funcs

import (
	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

func DiskIOMetrics() (L []*model.MetricValue) {

	disk_iocounter, err := IOCounters()
	if err != nil {
		g.Logger().Println(err)
		return
	}

	for device, ds := range disk_iocounter {

		device := "device=" + device
		L = append(L, CounterValue("disk.io.msec_read", ds.Msec_Read, device))
		L = append(L, CounterValue("disk.io.msec_write", ds.Msec_Write, device))
		L = append(L, CounterValue("disk.io.read_bytes", ds.Read_Bytes, device))
		L = append(L, CounterValue("disk.io.read_requests", ds.Read_Requests, device))
		L = append(L, CounterValue("disk.io.write_bytes", ds.Write_Bytes, device))
		L = append(L, CounterValue("disk.io.write_requests", ds.Write_Requests, device))
		L = append(L, GaugeValue("disk.io.util", 100-ds.Util, device))
	}
	return
}
