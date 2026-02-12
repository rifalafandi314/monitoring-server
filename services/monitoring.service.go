package services

import (
	"os/exec"
	"strings"

	"github.com/rifalafandi314/monitoring-server/models"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)


func GetSystemStat() models.SystemStat {
	cpuVal, _ := cpu.Percent(0, false)
	memVal, _ := mem.VirtualMemory()
	diskVal, _ := disk.Usage("/")
	uptime, _ := host.Uptime()

	status, threshold := AdaptiveThreshold(cpuVal[0])

	netIO, _ := net.IOCounters(false)

	processes, _ := process.Processes()
	procList := []models.Process{}

	for _, p := range processes {
		name, _ := p.Name()
		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryPercent()

		procList = append(procList, models.Process{
			PID: p.Pid,
			Name: name,
			CPU: cpu,
			RAM: mem,
		})

		if len(procList) > 20 {
			break
		}
	} 

	serviceList := []models.Service{}
	out, err := exec.Command("systemctl", "list-units", "--type=service", "--state=running").Output()
	if err != nil {
		lines := strings.Split(string(out), "\n")
		for _, l := range lines {
			if strings.Contains(l, ".service") {
				fields := strings.Fields(l)
				serviceList = append(serviceList, models.Service{
					Name: fields[0],
					Status: "Running",
				})
			}
		}
	}

	return models.SystemStat{
		CPU: cpuVal[0],
		RAM: memVal.UsedPercent,
		Disk: diskVal.UsedPercent,
		Uptime: uptime,
		Network: models.Network{
			Sent: netIO[0].BytesSent,
			Recv: netIO[0].BytesRecv,
		},
		CPUAnom: models.Anomaly{
			Status:    status,
			Threshold: threshold,
			Current:   cpuVal[0],
		},
		Processes: procList,
		Services:  serviceList,
	}
}