falcon-windows-agent
===

open-falcon 的 windows-agent, go 语言编写, 开箱即用
支持端口监控
支持进程监控
支持注册为 windows 服务后台运行
内置 IIs 监控 和 MsSQL(SqlServer) 监控。


#### 上报字段
Windows Metrics
--------------------------------
| Counters | Type |Tag| Notes|
|-----|------|------|------|
|agent.alive  |GAUGE|/|ailve |
|cpu.idle      |GAUGE|/|cpu idle|
|cpu.busy      |GAUGE|/|cpu busy|
|cpu.user      |GAUGE|/|cpu user|
|cpu.system      |GAUGE|/|cpu system|
|mem.memtotal      |GAUGE|/|mem total|
|mem.memused      |GAUGE|/|mem used|
|mem.memfree      |GAUGE|/|mem free|
|mem.memfree.percent      |GAUGE|/|memfree percent|
|mem.memused.percent      |GAUGE|/|memused percent|
|df.bytes.total      |GAUGE|mounts=Mountpoint,fstype=fstype|device bytes total|
|df.bytes.free      |GAUGE|mounts=Mountpoint,fstype=fstype|device bytes free|
|df.bytes.total      |GAUGE|mounts=Mountpoint,fstype=fstype|device bytes total|
|df.bytes.used.percent      |GAUGE|mounts=Mountpoint,fstype=fstype|device used percent|
|df.bytes.free.percent      |GAUGE|mounts=Mountpoint,fstype=fstype|device free percent|
|df.statistics.total      |GAUGE|mounts=Mountpoint,fstype=fstype|device statistics total|
|df.statistics.used      |GAUGE|mounts=Mountpoint,fstype=fstype|device statistics used|
|df.statistics.used.percent      |GAUGE|mounts=Mountpoint,fstype=fstype|device statistics used percent|
|disk.io.msec_read      |COUNTER|device=device|disk io msec read|
|disk.io.msec_write      |COUNTER|device=device|disk io msec write|
|disk.io.read_bytes      |COUNTER|device=device|disk io read bytes|
|disk.io.write bytes      |COUNTER|device=device|disk io write bytes|
|disk.io.read_requests      |COUNTER|device=device|disk io read requests|
|disk.io.write requests      |COUNTER|device=device|disk io write requests|
|disk.io.util      |COUNTER|device=device|disk io util|
|net.if.in.bytes      |COUNTER|iface=ifname|net if bytes recv|
|net.if.in.packets      |COUNTER|iface=ifname|net if packets recv|
|net.if.in.errors      |COUNTER|iface=ifname|net if errors recv|
|net.if.in.dropped      |COUNTER|iface=ifname|net if dropped recv|
|net.if.out.bytes      |COUNTER|iface=ifname|net if bytes sent|
|net.if.out.packets      |COUNTER|iface=ifname|net if packets sent|
|net.if.out.errors      |COUNTER|iface=ifname|net if errors sent|
|net.if.out.dropped      |COUNTER|iface=ifname|net if dropped sent|
|tcpip.confailures     |COUNTER|/|tcp connect failure|
|tcpip.conactive     |COUNTER|/|tcp connect active|
|tcpip.conpassive     |COUNTER|/|tcp connect passive|
|tcpip.conestablished     |GAUGE|/|tcp connect established |
|tcpip.conreset    |COUNTER|/|tcp connect reset|
|net.port.listen    |GAUGE|port=port|port alive|
|proc.num    |GAUGE|cmdline=cmdline,proc=proc|proc number|


IIs Metrics
--------------------------------
| Counters | Type |Tag| Notes|
|-----|------|------|------|
|iis.bytes.received  |COUNTER|site=site|Bytes Received/sec |
|iis.bytes.sent     |COUNTER|site=site|Total Bytes Sent/sec|
|iis.requests.cgi      |COUNTER|site=site|CGI Requests/sec|
|iis.requests.copy     |COUNTER|site=site|copy Requests/sec|
|iis.requests.delete     |COUNTER|site=site|delete Requests/sec|
|iis.requests.get     |COUNTER|site=site|get Requests/sec|
|iis.requests.head     |COUNTER|site=site|head Requests/sec|
|iis.requests.isapi     |COUNTER|site=site|isapi Requests/sec|
|iis.requests.lock     |COUNTER|site=site|lock Requests/sec|
|iis.requests.mkcol     |COUNTER|site=site|mkcol Requests/sec|
|iis.requests.move     |COUNTER|site=site|move Requests/sec|
|iis.requests.options     |COUNTER|site=site|options Requests/sec|
|iis.requests.post     |COUNTER|site=site|post Requests/sec|
|iis.requests.proppatch     |COUNTER|site=site|proppatch Requests/sec|
|iis.requests.propfind     |COUNTER|site=site|propfind Requests/sec|
|iis.requests.put     |COUNTER|site=site|put Requests/sec|
|iis.requests.search     |COUNTER|site=site|search Requests/sec|
|iis.requests.trace     |COUNTER|site=site|trace Requests/sec|
|iis.requests.unlock     |COUNTER|site=site|unlock Requests/sec|
|iis.errors.notfount     |COUNTER|site=site|notfound errors/sec|
|iis.errors.locked     |COUNTER|site=site|locked errors/sec|
|iis.connection.attempts    |COUNTER|site=site|conn attempts/sec|
|iis.connections    |GAUGE|site=site|connections|
|iis.service.uptime     |GAUGE|site=site|Service Uptime|

视版本和配置不同，采集到的 Metric 可能有所差别。

MsSQL
--------------------------------
| Counters | Type |Tag| Notes|
|-----|------|------|------|
|MsSQL.Lock_Waits/sec     |GAUGE|instance=instance|Lock_Waits/sec|
|MsSQL.Log_File(s)_Size_(KB)     |GAUGE|instance=instance|Log_File(s)_Size_(KB)|
|MsSQL.Log_File(s)_Used_Size_(KB)     |GAUGE|instance=instance|Log_File(s)_Used_Size_(KB)|
|MsSQL.Percent_Log_Used     |GAUGE|instance=instance|Log_File(s)_Used_Size_(KB)|
|MsSQL.Errors/sec     |GAUGE|error_type=error_type|Log_File(s)_Used_Size_(KB)|
|MsSQL.Batch_Requests/sec     |GAUGE|/|Batch_Requests/sec|
|MsSQL.Target_Server_Memory_(KB)     |GAUGE|/|Target_Server_Memory_(KB)|
|MsSQL.Total_Server_Memory_(KB)     |GAUGE|/|Total_Server_Memory_(KB)|
|MsSQL.IO_requests     |GAUGE|/|IO_requests|
|MsSQL.Connection     |GAUGE|/|Connections|
|MsSQL.Uptime    |GAUGE|/|Service Uptime|

视版本和配置不同，采集到的 Metric 可能有所差别。
其中Lock_Waits/sec …… Total_Server_Memory_(KB) 等通过查询 sys.dm_os_performance_counters 表获得，这需要服务器上开启性能计数器。

如果这部分指标缺失，请确认性能计数器是否正确开启。



#### 使用方式


配置文件请参照cfg.example.json，修改该文件名为cfg.json

```
{
    "debug": true,
	"logfile": "windows.log",  //日志的输出路径
    "hostname": "",
    "ip": "",
	"iis":{
		"enabled": false,
		"websites": [
	        "Default Web Site" //web 的站点，可以留空，默认会采集_Total的
	    ]
 	}, 
	"mssql":{
		"enabled": false,
		"addr":"127.0.0.1",
		"port":1433,
		"username":"sa",
		"password":"123456",
		"encrypt":"disable",
		//disable - 不加密
		//false - 除认证报文外不加密
		//true -加密
		//SQL Server 2008 和 SQL Server 2008 R2 必须选择 disable，否则无法正常认证。要修复这个问题，需要升级 SQL Server 2008 R2 SP2，或 SQL Server 2008 SP3
		"instance": [ //要采集数据库实例名
	        "test"
	    ]
 	}, 
    "heartbeat": {
        "enabled": true,
        "addr": "127.0.0.1:6030",
        "interval": 60,
        "timeout": 1000
    },
    "transfer": {
        "enabled": true,
        "addrs": [
            "127.0.0.1:8433"
        ],
        "interval": 60,
        "timeout": 1000
    },
    "http": {
        "enabled": true,
        "listen": ":1988",
        "backdoor": false
    },
    "collector": {
        "ifacePrefix": ["Intel"] //所采集的网卡描述信息关键词，例如Intel(R)PRO/1000 MT NetworkConnection
    },
    "ignore": {
        "cpu.busy": true,
    }
}

```

#### http 信息维护接口

```
curl http://127.0.0.1:1988/health
正常则返回 ok

curl http://127.0.0.1:1988/version
返回版本

curl http://127.0.0.1:1988/workdir
返回工作目录
 
curl http://127.0.0.1:1988/config
返回配置
```

#### http 转发接口
```
http://127.0.0.1:1988//v1/push
```
#### 源码安装

```
cd %GOPATH%/src/github.com/freedomkk-qfeng/windows-agent
go get ./...
go build -o windows-agent.exe

```

#### 运行
以下命令需在管理员模式下运行开 cmd 或 Powershell

先试运行一下
```
.\windows-agent.exe
2016/08/08 13:44:31 cfg.go:96: read config file: cfg.json successfully
2016/08/08 13:44:31 var.go:31: logging on windows.log
2016/08/08 13:44:31 http.go:64: listening :1988
```
等待1-2分钟，观察输出，确认运行正常
使用 [nssm](https://nssm.cc/) 注册为 Windows 服务。

```
.\nssm.exe install windows-agent
```
![](http://i.imgur.com/SOhBSBo.png)


启动服务
```
.\nssm.exe start windows-agent
```


#### TODO
增加完善 sqlserver 的监控项
