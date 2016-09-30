package funcs

var MsSQL_Mertics_instance = map[string]string{
	"Lock_Waits/sec":             "GUAGE",
	"Average_Wait_Time_(ms)":     "GUAGE",
	"Log_File(s)_Size_(KB)":      "GUAGE",
	"Log_File(s)_Used_Size_(KB)": "GUAGE",
	"Percent_Log_Used":           "GUAGE",
}

var MsSQL_Mertics = map[string]string{
	"Errors/sec":                "GUAGE",
	"Target_Server_Memory_(KB)": "GUAGE",
	"Total_Server_Memory_(KB)":  "GUAGE",
	"Batch_Requests/sec":        "GUAGE",
}
