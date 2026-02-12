package models

type Process struct {
	PID  int32   `json:"pid"`
	Name string  `json:"name"`
	CPU  float64 `json:"cpu"`
	RAM  float32 `json:"ram"`
}

type Service struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Network struct {
	Sent uint64 `json:"sent"`
	Recv uint64 `json:"recv"`
}

type Anomaly struct {
	Status    string  `json:"status"`  // normal, warbing, critical
	Threshold float64 `json:"threshold"`
	Current   float64 `json:"current"`
}


type SystemStat struct {
	CPU       float64   `json:"cpu"`
	RAM       float64   `json:"ram"`
	Disk      float64   `json:"disk"`
	CPUAnom   Anomaly   `json:"cpu_anomaly"`
	Uptime    uint64    `json:"uptime"`
	Network   Network   `json:"network"`
	Processes []Process `json:"processes"`
	Services  []Service `json:"services"`
}
