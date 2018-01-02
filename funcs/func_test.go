package funcs

import (
	"testing"
)

const (
	server   = "192.168.11.128"
	port     = 1433
	user     = "sa"
	password = "ecNU10@$"
	encrypt  = "disable"
)

func Test_cpu(t *testing.T) {
	times, err := CPUTimes(false)
	t.Log(times)
	t.Error(err)
}

func Test_mem_info(t *testing.T) {
	meminfo, err := mem_info()
	t.Log(meminfo)
	t.Error(err)
}

func Test_disk(t *testing.T) {
	diskPartitions, err := disk_Partitions()
	t.Log(diskPartitions)
	t.Error(err)
	diskUsage, err := disk_usage("E:")
	t.Log(diskUsage)
	t.Error(err)
}

func Test_net_status(t *testing.T) {
	var ifacePrefix = []string{"本地连接", "Loop"}
	netifs, err := net_status(ifacePrefix)
	t.Log(netifs)
	t.Error(err)
}

func Test_IsTCPPortUsed(t *testing.T) {
	res := CheckTCPPortUsed(1988)
	t.Log(res)
}

func Test_TestIOCounters(t *testing.T) {
	r, err := PerfFormattedData()
	t.Log(r)
	t.Error(err)
	ret, err := IOCounters()
	t.Log(ret)
	t.Error(err)
}

func Test_Process(t *testing.T) {
	p, err := Processes()
	t.Log(p)
	t.Error(err)
	cnt := 0
	m := map[int]string{
		1: "smss.exe",
	}
	for i := 0; i < len(p); i++ {
		if is_a(p[i], m) {
			cnt++
		}
	}
	t.Log(cnt)
}
func Test_tcpip(t *testing.T) {
	ret, _ := TcpipCounters()
	t.Log(ret)
}

func Test_iis_status(t *testing.T) {
	result, err := IIsCounters()
	t.Error(err)
	t.Log(result)
}

func Test_in_array(t *testing.T) {
	instance := []string{"test"}
	re := in_array("test", instance)
	t.Log(re)
}

func Test_performance_query(t *testing.T) {
	instance := []string{"_Total", "test"}
	db, err := mssql_conn(server, port, user, password, encrypt)
	if err != nil {
		t.Error(err)
	}
	result, err := performance_query(db, instance)
	t.Log(result)
	t.Error(err)
	db.Close()
}

func Test_io_req_query(t *testing.T) {
	db, err := mssql_conn(server, port, user, password, encrypt)
	if err != nil {
		t.Error(err)
	}
	result, err := io_req_query(db)
	t.Log(result)
	t.Error(err)
	db.Close()
}

func Test_conn_query(t *testing.T) {
	db, err := mssql_conn(server, port, user, password, encrypt)
	if err != nil {
		t.Error(err)
	}
	result, err := conn_query(db)
	t.Log(result)
	t.Error(err)
	db.Close()
}

func Test_uptime_query(t *testing.T) {
	db, err := mssql_conn(server, port, user, password, encrypt)
	if err != nil {
		t.Error(err)
	}
	result, err := uptime_query(db)
	t.Log(result)
	t.Error(err)
	db.Close()
}
