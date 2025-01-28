package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// saveToFile сохраняет информацию в файл hw_info.txt
func saveToFile(data string) {
	file, err := os.Create("hw_info.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}
}

// executeCommand выполняет команду и возвращает результат
func executeCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return fmt.Sprintf("Ошибка выполнения команды: %s", err)
	}
	return out.String()
}

func main() {
	var report strings.Builder

	// CPU Information
	cpuInfo, _ := cpu.Info()
	report.WriteString("CPU Information:\n")
	for _, cpu := range cpuInfo {
		report.WriteString(fmt.Sprintf("  Model: %s\n  Cores: %d\n  Frequency: %.2f MHz\n",
			cpu.ModelName, cpu.Cores, cpu.Mhz))
	}

	// RAM Information
	vmStat, _ := mem.VirtualMemory()
	report.WriteString("\nRAM Information:\n")
	report.WriteString(fmt.Sprintf("  Total: %.2f GB\n  Used: %.2f GB\n  Free: %.2f GB\n",
		float64(vmStat.Total)/1e9, float64(vmStat.Used)/1e9, float64(vmStat.Free)/1e9))

	// GPU Information (using lspci)
	report.WriteString("\nGPU Information:\n")
	gpuInfo := executeCommand("lspci", "-nnk")
	report.WriteString(gpuInfo)

	// USB Devices (using lsusb)
	report.WriteString("\nUSB Devices:\n")
	usbInfo := executeCommand("lsusb")
	report.WriteString(usbInfo)

	// Network Adapters
	netInterfaces, _ := net.Interfaces()
	report.WriteString("\nNetwork Adapters:\n")
	for _, iface := range netInterfaces {
		status := "Inactive"
		if iface.Flags[0] == "up" {
			status = "Active"
		}
		report.WriteString(fmt.Sprintf("  Name: %s\n  Status: %s\n  MAC: %s\n",
			iface.Name, status, iface.HardwareAddr))
	}

	// Save to file
	saveToFile(report.String())

	// Display with dialog
	cmd := exec.Command("dialog", "--title", "Hardware Information", "--msgbox", report.String(), "30", "80")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Ошибка запуска dialog:", err)
	}
}
