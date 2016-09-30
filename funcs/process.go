package funcs

import (
	"strings"
	"time"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
	"github.com/shirou/gopsutil/process"
)

type P struct {
	name    string
	cmdline string
}

func ProcMetrics() (L []*model.MetricValue) {

	reportProcs := g.ReportProcs()
	sz := len(reportProcs)
	if sz == 0 {
		return
	}
	startTime := time.Now()
	ps, err := Processes()
	if err != nil {
		g.Logger().Println(err)
		return
	}

	pslen := len(ps)

	for tags, m := range reportProcs {
		cnt := 0
		for i := 0; i < pslen; i++ {
			if is_a(ps[i], m) {
				cnt++
			}
		}

		L = append(L, GaugeValue(g.PROC_NUM, cnt, tags))
	}
	endTime := time.Now()
	g.Logger().Printf("UpdateProcessStats complete. Process time %s. Number of Process is %d", endTime.Sub(startTime), pslen)
	return
}

func is_a(p P, m map[int]string) bool {
	// only one kv pair
	for key, val := range m {
		if key == 1 {
			// name
			if val != p.name {
				return false
			}
		} else if key == 2 {
			// cmdline
			if !strings.Contains(p.cmdline, val) {
				return false
			}
		}
	}
	return true
}

func Processes() ([]P, error) {
	var processes = []P{}
	var PROCESS P
	pids, err := process.Pids()
	if err != nil {
		return processes, err
	}
	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err == nil {
			pname, err := p.Name()
			pcmdline, err := p.Cmdline()
			if err == nil {
				PROCESS.name = pname
				PROCESS.cmdline = pcmdline
				processes = append(processes, PROCESS)
			}
		}
	}
	return processes, err
}
