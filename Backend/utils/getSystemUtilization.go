package utils

import (
	"log"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type MemoryStats struct {
	UsedMemory float64
	FreeMemory float64
}

type CPUStats struct {
	UsedCPU float64
	FreeCPU float64
}

type SystemStats struct {
	MemoryStats MemoryStats
	CPUStats    CPUStats
}

func getMemoryStats() MemoryStats {
	memory, err := memory.Get()
	if err != nil {
		log.Fatal(err)
	}

	// Log the memory used as a percentage
	log.Printf("memory used: %f %%\n", float64(memory.Used)/float64(memory.Total)*100)

	// Log the memory free as a percentage
	log.Printf("memory free: %f %%\n", 100-float64(memory.Used)/float64(memory.Total)*100)

	// Return the memory stats
	return MemoryStats{
		UsedMemory: float64(memory.Used) / float64(memory.Total) * 100,
		FreeMemory: 100 - float64(memory.Used)/float64(memory.Total)*100,
	}
}

func getCPUStats() CPUStats {
	before, err := cpu.Get()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		log.Fatal(err)
	}
	total := float64(after.Total - before.Total)
	log.Printf("cpu in use: %f %%\n", float64(after.User-before.User)/total*100+float64(after.System-before.System)/total*100)
	log.Printf("cpu free: %f %%\n", float64(after.Idle-before.Idle)/total*100)

	return CPUStats{
		UsedCPU: float64(after.User-before.User)/total*100 + float64(after.System-before.System)/total*100,
		FreeCPU: float64(after.Idle-before.Idle) / total * 100,
	}
}

func GetSystemUtilization() SystemStats {
	memoryStatistics := getMemoryStats()
	cpuStatistics := getCPUStats()

	return SystemStats{
		MemoryStats: memoryStatistics,
		CPUStats:    cpuStatistics,
	}
}
