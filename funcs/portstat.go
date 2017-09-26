package funcs

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

const (
	minTCPPort = 0
	maxTCPPort = 65535
)

func IsTCPPortV4Used(port int64) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	connString := "127.0.0.1:" + strconv.FormatInt(port, 10)
	conn, err := net.Listen("tcp", connString)
	log.Println(connString, conn, err)
	if err != nil {
		return true
	}
	conn.Close()
	return false
}

func IsTCPPortV6Used(port int64) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	connString := "[::1]:" + strconv.FormatInt(port, 10)
	conn, err := net.Listen("tcp", connString)
	log.Println(connString, conn, err)
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
		if IsTCPPortV4Used(reportPorts[i]) || IsTCPPortV6Used(reportPorts[i]) {
			L = append(L, GaugeValue(g.NET_PORT_LISTEN, 1, tags))
		} else {
			L = append(L, GaugeValue(g.NET_PORT_LISTEN, 0, tags))
		}
	}

	return
}
