package g

import (
	"time"
)

// changelog:
// 1.0.0 windows-agent
// 1.0.1 ifstat use ifname instead ifdescription
// 1.0.2 fix net.listen.port bug 
const (
	VERSION          = "1.0.2"
	COLLECT_INTERVAL = time.Second
	NET_PORT_LISTEN  = "net.port.listen"
	DU_BS            = "du.bs"
	PROC_NUM         = "proc.num"
)
