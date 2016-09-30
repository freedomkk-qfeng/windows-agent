package funcs

import (
	"fmt"

	"net"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

const (
	minTCPPort = 0
	maxTCPPort = 65535
)

func IsTCPPortUsed(port int64) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}

	conn, err := net.Listen("tcp", ":"+string(port))

	if err != nil {
		return true
	}
	conn.Close()
	return false
}

func PortMetrics() (L []*model.MetricValue) {
	reportPorts := g.ReportPorts()
	sz := len(reportPorts)
	if sz == 0 {
		return
	}

	for i := 0; i < sz; i++ {
		tags := fmt.Sprintf("port=%d", reportPorts[i])
		if IsTCPPortUsed(reportPorts[i]) {
			L = append(L, GaugeValue(g.NET_PORT_LISTEN, 1, tags))
		} else {
			L = append(L, GaugeValue(g.NET_PORT_LISTEN, 0, tags))
		}
	}

	return
}
