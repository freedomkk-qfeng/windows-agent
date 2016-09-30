package funcs

import (
	"syscall"

	"unsafe"
)

var (
	Modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	ModNt       = syscall.NewLazyDLL("ntdll.dll")
	ModPdh      = syscall.NewLazyDLL("pdh.dll")

	ProcGetSystemTimes           = Modkernel32.NewProc("GetSystemTimes")
	ProcNtQuerySystemInformation = ModNt.NewProc("NtQuerySystemInformation")
	PdhOpenQuery                 = ModPdh.NewProc("PdhOpenQuery")
	PdhAddCounter                = ModPdh.NewProc("PdhAddCounterW")
	PdhCollectQueryData          = ModPdh.NewProc("PdhCollectQueryData")
	PdhGetFormattedCounterValue  = ModPdh.NewProc("PdhGetFormattedCounterValue")
	PdhCloseQuery                = ModPdh.NewProc("PdhCloseQuery")
)

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type CPUTimesStat struct {
	User   float64 `json:"user"`
	System float64 `json:"system"`
	Idle   float64 `json:"idle"`
	Total  float64 `json:"total"`
}

type Win32_Processor struct {
	LoadPercentage            *uint16
	Family                    uint16
	Manufacturer              string
	Name                      string
	NumberOfLogicalProcessors uint32
	ProcessorId               *string
	Stepping                  *string
	MaxClockSpeed             uint32
}

func CPUTimes(percpu bool) ([]CPUTimesStat, error) {
	var ret []CPUTimesStat

	var lpIdleTime FILETIME
	var lpKernelTime FILETIME
	var lpUserTime FILETIME
	r, _, _ := ProcGetSystemTimes.Call(
		uintptr(unsafe.Pointer(&lpIdleTime)),
		uintptr(unsafe.Pointer(&lpKernelTime)),
		uintptr(unsafe.Pointer(&lpUserTime)))
	if r == 0 {
		return ret, syscall.GetLastError()
	}

	LOT := float64(0.0000001)
	HIT := (LOT * 4294967296.0)
	idle := ((HIT * float64(lpIdleTime.DwHighDateTime)) + (LOT * float64(lpIdleTime.DwLowDateTime)))
	user := ((HIT * float64(lpUserTime.DwHighDateTime)) + (LOT * float64(lpUserTime.DwLowDateTime)))
	kernel := ((HIT * float64(lpKernelTime.DwHighDateTime)) + (LOT * float64(lpKernelTime.DwLowDateTime)))
	system := (kernel - idle)

	ret = append(ret, CPUTimesStat{
		Idle:   float64(idle),
		User:   float64(user),
		System: float64(system),
		Total:  float64(idle) + float64(user) + float64(system),
	})
	return ret, nil
}
