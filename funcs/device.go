package funcs

import (
	"fmt"

	"github.com/freedomkk-qfeng/windows-agent/g"

	"github.com/open-falcon/common/model"
	"github.com/shirou/gopsutil/disk"
)

func disk_usage(path string) (*disk.UsageStat, error) {
	disk_usage, err := disk.Usage(path)
	return disk_usage, err
}

func disk_Partitions() ([]disk.PartitionStat, error) {
	disk_Partitions, err := disk.Partitions(true)
	return disk_Partitions, err
}
func DeviceMetrics() (L []*model.MetricValue) {
	diskPartitions, err := disk_Partitions()

	if err != nil {
		g.Logger().Println(err)
		return
	}

	var diskTotal uint64 = 0
	var diskUsed uint64 = 0

	for _, device := range diskPartitions {
		du, err := disk_usage(device.Mountpoint)
		if err != nil {
			g.Logger().Println(err)
			continue
		}

		diskTotal += du.Total
		diskUsed += du.Used

		tags := fmt.Sprintf("mount=%s,fstype=%s", device.Mountpoint, device.Fstype)
		L = append(L, GaugeValue("df.bytes.total", du.Total, tags))
		L = append(L, GaugeValue("df.bytes.used", du.Used, tags))
		L = append(L, GaugeValue("df.bytes.free", du.Free, tags))
		L = append(L, GaugeValue("df.bytes.used.percent", du.UsedPercent, tags))
		L = append(L, GaugeValue("df.bytes.free.percent", 100-du.UsedPercent, tags))
	}

	if len(L) > 0 && diskTotal > 0 {
		L = append(L, GaugeValue("df.statistics.total", float64(diskTotal)))
		L = append(L, GaugeValue("df.statistics.used", float64(diskUsed)))
		L = append(L, GaugeValue("df.statistics.used.percent", float64(diskUsed)*100.0/float64(diskTotal)))
	}

	return
}
