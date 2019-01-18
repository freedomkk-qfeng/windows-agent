package funcs

import (
	"fmt"
	"net"
	"strconv"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

const (
	minTCPPort = 0
	maxTCPPort = 65535
)

func IsTCPPortUsed(addr string, port int64) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	connString := addr + strconv.FormatInt(port, 10)
	conn, err := net.Listen("tcp", connString)
	if err != nil {
		//		log.Println(connString, conn, err)
		return true
	}
	conn.Close()
	return false
}

func IsTCPPortUsed2(ip string, port int64) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	connStr := ip + ":" + strconv.FormatInt(port, 10)
	tcpAddr, err := net.ResolveTCPAddr("tcp", connStr)
	if err != nil {
		return false
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

func CheckTCPPortUsed(port int64) bool {
	//if IsTCPPortUsed("0.0.0.0:", port) {
	//	return true
	//}
	//if IsTCPPortUsed("127.0.0.1:", port) {
	//	return true
	//}
	//if IsTCPPortUsed("[::1]:", port) {
	//	return true
	//}
	//if IsTCPPortUsed("[::]:", port) {
	//	return true
	//}
	//return false
	return IsTCPPortUsed2("127.0.0.1", port)
}

func PortMetrics() (L []*model.MetricValue) {
	reportPorts := g.ReportPorts()
	sz := len(reportPorts)
	if sz == 0 {
		return
	}

	for i := 0; i < sz; i++ {
		tags := fmt.Sprintf("port=%d", reportPorts[i])
		if CheckTCPPortUsed(reportPorts[i]) {
			L = append(L, GaugeValue(g.NET_PORT_LISTEN, 1, tags))
		} else {
			L = append(L, GaugeValue(g.NET_PORT_LISTEN, 0, tags))
		}
	}

	return
}
