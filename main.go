package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/disk"
	// "github.com/shirou/gopsutil/v3/cpu"
	// "github.com/shirou/gopsutil/v3/docker"
	// "github.com/shirou/gopsutil/v3/host"
	// "github.com/shirou/gopsutil/v3/internal"
	// "github.com/shirou/gopsutil/v3/load"
	// "github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/v3/net"
	// "github.com/shirou/gopsutil/v3/process"
	// "github.com/shirou/gopsutil/v3/winservices"
)

func main() {

	memVirtualMemory, _ := mem.VirtualMemory()
	cpuInfos, _ := cpu.Info()
	//cpuCountsPhysical, _ := cpu.Counts(false)
	//cpuCountsLogical, _ := cpu.Counts(true)
	diskPartitions, _ := disk.Partitions(false)

	fmt.Printf("\n-RAM INFO-\n")
	fmt.Printf("Total RAM: %v MB\nFree RAM: %v MB\nUsedPercent: %v%%\n",
		memVirtualMemory.Total,
		memVirtualMemory.Free,
		memVirtualMemory.UsedPercent)

	fmt.Printf("\n-CPU INFO-\n")
	for _, cpuInfo := range cpuInfos {

		fmt.Printf("CPU Model: %v\nMhZ: %v\nCores: %v\n",
			cpuInfo.ModelName,
			cpuInfo.Mhz,
			cpuInfo.Cores)
	}

	fmt.Printf("\n-DISK INFO-\n")
	for _, diskPartidiskPartition := range diskPartitions {
		diskUsage, _ := disk.Usage(diskPartidiskPartition.Mountpoint)
		fmt.Printf("Device: %v\nFstype: %v\nMountpoint: %v\nTotal: %v MB\nFree: %v MB\n",
			diskPartidiskPartition.Device,
			diskPartidiskPartition.Fstype,
			diskPartidiskPartition.Mountpoint,
			diskUsage.Total,
			diskUsage.Free)
	}

}
