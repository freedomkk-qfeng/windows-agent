package funcs

import (
	"github.com/freedomkk-qfeng/windows-agent/g"

	"github.com/open-falcon/common/model"
	"github.com/shirou/gopsutil/mem"
)

func mem_info() (*mem.VirtualMemoryStat, error) {
	meminfo, err := mem.VirtualMemory()
	return meminfo, err
}

func MemMetrics() []*model.MetricValue {
	meminfo, err := mem_info()
	if err != nil {
		g.Logger().Println(err)
		return []*model.MetricValue{}
	}
	memTotal := meminfo.Total
	memUsed := meminfo.Used
	memFree := meminfo.Available
	pmemUsed := 100 * float64(memUsed) / float64(memTotal)
	pmemFree := 100 * float64(memFree) / float64(memTotal)

	return []*model.MetricValue{
		GaugeValue("mem.memtotal", memTotal),
		GaugeValue("mem.memused", memUsed),
		GaugeValue("mem.memfree", memFree),
		GaugeValue("mem.memfree.percent", pmemFree),
		GaugeValue("mem.memused.percent", pmemUsed),
	}

}
