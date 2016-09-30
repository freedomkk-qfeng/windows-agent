package funcs

import (
	"fmt"
)

func CheckCollector() {

	output := make(map[string]bool)

	output["df.bytes"] = len(DeviceMetrics()) > 0
	output["memory  "] = len(MemMetrics()) > 0

	for k, v := range output {
		status := "fail"
		if v {
			status = "ok"
		}
		fmt.Println(k, "...", status)
	}
}
