package funcs

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

type mssql struct {
	metric string
	value  float64
	Type   string
	Tag    string
}

func mssqlMetrics() (L []*model.MetricValue) {
	if !g.Config().MsSQL.Enabled {
		g.Logger().Println("MsSQL Monitor is disabled")
		return
	}

	server := g.Config().MsSQL.Addr
	port := g.Config().MsSQL.Port
	user := g.Config().MsSQL.Username
	password := g.Config().MsSQL.Password
	instance := g.Config().MsSQL.Instance
	encrypt := g.Config().MsSQL.Encrypt
	instance = append(instance, "_Total")

	db, err := mssql_conn(server, port, user, password, encrypt)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	uptime, err := uptime_query(db)
	if err == nil {
		L = append(L, GaugeValue("MsSQL.Uptime", uptime))
	} else {
		g.Logger().Println(err)
	}
	conn, err := conn_query(db)
	if err == nil {
		L = append(L, GaugeValue("MsSQL.Connection", conn))
	} else {
		g.Logger().Println(err)
	}
	io_req, err := io_req_query(db)
	if err == nil {
		L = append(L, GaugeValue("MsSQL.IO_requests", io_req))
	} else {
		g.Logger().Println(err)
	}
	mssql_performance, err := performance_query(db, instance)
	if err == nil {
		for _, perf := range mssql_performance {
			switch perf.Type {
			case "GUAGE":
				L = append(L, GaugeValue("MsSQL."+perf.metric, perf.value, perf.Tag))
			case "COUNTER":
				L = append(L, GaugeValue("MsSQL."+perf.metric, perf.value, perf.Tag))
			}
		}
	} else {
		g.Logger().Println(err)
	}
	return
}

func mssql_conn(server string, port int, user string, password string, encrypt string) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;encrypt=%s", server, user, password, port, encrypt)
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}
	return db, err

}

func uptime_query(db *sql.DB) (int64, error) {
	sql := "select DATEDIFF(second,(select crdate from master.dbo.sysdatabases where NAME = 'tempdb'),GETDATE())"
	rows, err := db.Query(sql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return 0, err
	}
	if cols == nil {
		return 0, err
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}
	var uptime int64
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			g.Logger().Println(err)
			continue
		}
		v := vals[0].(*interface{})
		uptime = (*v).(int64)

	}

	return uptime, nil
}

func io_req_query(db *sql.DB) (float64, error) {
	sql := "select count(*) as io_req from sys.dm_io_pending_io_requests"
	rows, err := db.Query(sql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return 0, err
	}
	if cols == nil {
		return 0, err
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}
	var value float64
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			g.Logger().Println(err)
			continue
		}
		v := vals[0].(*interface{})
		vv := (*v).(int64)
		value = float64(vv)
	}
	return value, nil
}

func conn_query(db *sql.DB) (float64, error) {
	sql := "SELECT COUNT(*) AS CONNECTIONS FROM sys.dm_exec_connections"
	rows, err := db.Query(sql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return 0, err
	}
	if cols == nil {
		return 0, err
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}
	var value float64
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			g.Logger().Println(err)
			continue
		}
		v := vals[0].(*interface{})
		vv := (*v).(int64)
		value = float64(vv)
	}
	return value, nil
}

func performance_query(db *sql.DB, instance []string) ([]mssql, error) {
	sql := "select * from sys.dm_os_performance_counters"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	if cols == nil {
		return nil, err
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}
	var result []mssql
	var result_value mssql
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			g.Logger().Println(err)
			continue
		}
		v1 := vals[1].(*interface{})
		v2 := vals[2].(*interface{})
		v3 := vals[3].(*interface{})
		counter_name := (*v1).(string)
		instance_name := (*v2).(string)
		value := (*v3).(int64)
		counter_name = format_mertic(counter_name)
		instance_name = format_mertic(instance_name)
		if Type, ok := MsSQL_Mertics_instance[counter_name]; ok {
			if in_array(instance_name, instance) {
				result_value.metric = counter_name
				result_value.Tag = "instance=" + instance_name
				result_value.Type = Type
				result_value.value = float64(value)
				result = append(result, result_value)
			}
		}
		if Type, ok := MsSQL_Mertics[counter_name]; ok {
			if counter_name == "Errors/sec" {
				result_value.metric = counter_name
				result_value.Tag = "error_type=" + instance_name
				result_value.Type = Type
				result_value.value = float64(value)
				result = append(result, result_value)
			} else {
				result_value.metric = counter_name
				result_value.Tag = ""
				result_value.Type = Type
				result_value.value = float64(value)
				result = append(result, result_value)
			}

		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
}

func format_mertic(metric string) string {
	result := strings.TrimSpace(metric)
	result = strings.Replace(result, " ", "_", -1)
	return result
}

func in_array(a string, array []string) bool {
	for _, v := range array {
		v = format_mertic(v)
		if a == v {
			return true
		}
	}
	return false
}
