package services

import (
	"github.com/rifalafandi314/monitoring-server/models"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)


func GetSystemStat() models.SystemStat {
	cpuVal, _ := cpu.Percent(0, false)
	memVal, _ := mem.VirtualMemory()
	diskVal, _ := disk.Usage("/")

	return models.SystemStat{
		CPU: cpuVal[0],
		RAM: memVal.UsedPercent,
		Disk: diskVal.UsedPercent,
	}
}