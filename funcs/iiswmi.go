package funcs

import (
	"github.com/StackExchange/wmi"
)

type Win32_PerfRawData_W3SVC_WebService struct {
	BytesReceivedPersec          uint64
	BytesSentPersec              uint64
	CGIRequestsPersec            uint32
	ConnectionAttemptsPersec     uint32
	CopyRequestsPersec           uint32
	CurrentConnections           uint32
	DeleteRequestsPersec         uint32
	GetRequestsPersec            uint32
	HeadRequestsPersec           uint32
	ISAPIExtensionRequestsPersec uint32
	LockRequestsPersec           uint32
	LockedErrorsPersec           uint32
	MkcolRequestsPersec          uint32
	MoveRequestsPersec           uint32
	Name                         string
	NotFoundErrorsPersec         uint32
	OptionsRequestsPersec        uint32
	PostRequestsPersec           uint32
	PropfindRequestsPersec       uint32
	ProppatchRequestsPersec      uint32
	PutRequestsPersec            uint32
	SearchRequestsPersec         uint32
	TraceRequestsPersec          uint32
	UnlockRequestsPersec         uint32
	ServiceUptime                uint32
}

func IIsCounters() ([]Win32_PerfRawData_W3SVC_WebService, error) {
	var dst []Win32_PerfRawData_W3SVC_WebService
	err := wmi.Query("SELECT * FROM Win32_PerfRawData_W3SVC_WebService", &dst)
	if err != nil {
		return dst, err
	}
	return dst, nil
}
