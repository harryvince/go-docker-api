package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/harryvince/go-docker-api/utils"
)

// Endpoint to get system stats
func SystemStats(c *gin.Context) {
	var systemStatistics utils.SystemStats = utils.GetSystemUtilization()
	c.JSON(200, gin.H{
		"CPU": gin.H{
			"Used": systemStatistics.CPUStats.UsedCPU,
			"Free": systemStatistics.CPUStats.FreeCPU,
		},
		"Memory": gin.H{
			"Used": systemStatistics.MemoryStats.UsedMemory,
			"Free": systemStatistics.MemoryStats.FreeMemory,
		},
	})
}
