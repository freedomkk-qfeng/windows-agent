package funcs

import (
	"log"

	"github.com/StackExchange/wmi"
	"github.com/open-falcon/common/model"
)

type Tcpipdatastat struct {
	ConFailures    uint64 `json:"confailures"`
	ConActive      uint64 `json:"conactive"`
	ConPassive     uint64 `json:"conpassive"`
	ConEstablished uint64 `json:"conestablished"`
	ConReset       uint64 `json:"conreset"`
}

type Win32_TCPPerfFormattedData struct {
	ConnectionFailures     uint64
	ConnectionsActive      uint64
	ConnectionsPassive     uint64
	ConnectionsEstablished uint64
	ConnectionsReset       uint64
}

func TcpipMetrics() (L []*model.MetricValue) {

	ds, err := TcpipCounters()
	if err != nil {
		log.Println("Get tcpip data fail: ", err)
		return
	}

	L = append(L, CounterValue("tcpip.confailures", ds[0].ConFailures))
	L = append(L, CounterValue("tcpip.conactive", ds[0].ConActive))
	L = append(L, CounterValue("tcpip.conpassive", ds[0].ConPassive))
	L = append(L, GaugeValue("tcpip.conestablished", ds[0].ConEstablished))
	L = append(L, CounterValue("tcpip.conreset", ds[0].ConReset))

	return
}

func TcpipCounters() ([]Tcpipdatastat, error) {
	ret := make([]Tcpipdatastat, 0)
	var dst []Win32_TCPPerfFormattedData
	err := wmi.Query("SELECT ConnectionFailures,ConnectionsActive,ConnectionsPassive,ConnectionsEstablished,ConnectionsReset FROM Win32_PerfRawData_Tcpip_TCPv4", &dst)
	if err != nil {
		return ret, err
	}

	for _, d := range dst {

		ret = append(ret, Tcpipdatastat{
			ConFailures:    uint64(d.ConnectionFailures),
			ConActive:      uint64(d.ConnectionsActive),
			ConPassive:     uint64(d.ConnectionsPassive),
			ConEstablished: uint64(d.ConnectionsEstablished),
			ConReset:       uint64(d.ConnectionsReset),
		})
	}

	return ret, nil
}
