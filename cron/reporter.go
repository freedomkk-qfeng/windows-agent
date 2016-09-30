package cron

import (
	"fmt"

	"time"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

func ReportAgentStatus() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go reportAgentStatus(time.Duration(g.Config().Heartbeat.Interval) * time.Second)
	}
}

func reportAgentStatus(interval time.Duration) {
	for {
		hostname, err := g.Hostname()
		if err != nil {
			hostname = fmt.Sprintf("error:%s", err.Error())
		}

		req := model.AgentReportRequest{
			Hostname:     hostname,
			IP:           g.IP(),
			AgentVersion: g.VERSION,
		}

		var resp model.SimpleRpcResponse
		err = g.HbsClient.Call("Agent.ReportStatus", req, &resp)
		if err != nil || resp.Code != 0 {
			g.Logger().Println("call Agent.ReportStatus fail:", err, "Request:", req, "Response:", resp)
		}

		time.Sleep(interval)
	}
}
