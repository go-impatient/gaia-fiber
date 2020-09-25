package sd

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	_  = iota             // 0
	KB = 1 << (10 * iota) // 1 << (10 * 1)
	MB                    // 1 << (10 * 2)
	GB                    // 1 << (10 * 3)
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *fiber.Ctx) error {
	message := "OK"
	return c.SendString(message)
}

// DiskCheck checks the disk usage.
func DiskCheck(c *fiber.Ctx) error {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)

	return c.Status(status).SendString(message)
}

// OSCheck OS
func OSCheck(c *fiber.Ctx) error {
	status := http.StatusOK
	text := "OK"
	goOs := runtime.GOOS
	compiler := runtime.Compiler
	numCpu := runtime.NumCPU()
	version := runtime.Version()
	numGoroutine := runtime.NumGoroutine()
	message := fmt.Sprintf("%s - %s| %s | %s | %s| %s", text, goOs, compiler, numCpu, version, numGoroutine)
	return c.Status(status).SendString(message)
}

// CPUCheck checks the cpu usage.
func CPUCheck(c *fiber.Ctx) error {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)
	return c.Status(status).SendString(message)
}

// RAMCheck checks the disk usage.
func RAMCheck(c *fiber.Ctx) error {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	return c.Status(status).SendString(message)
}
