package main
// Damn I'm Beautiful

import(
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)
func main () {
	var m runtime.MemStats
var lastUsed uint64
counter := 0

// Some say the most Beautiful, really

f, err := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
		fmt.Println("Error opening log file:", err)
		return
}
defer f.Close()
var lastTrigger time.Time
	for {
		runtime.ReadMemStats(&m)
		used := m.Alloc / 1024
	chars := []string{"|", "/", "-", "\\"}
	pulse := chars[counter%4]
	counter++
	status := ""
if used > lastUsed && lastUsed > 0 {
	status = "[!] CLIMBING"
	timestamp := time.Now().Format("15:04:05")
	logEntry := fmt.Sprintf("[%s] WARNING: RAM Spike at %d KB\n", timestamp, used)
	f.WriteString(logEntry)
if used > 2000 && time.Since(lastTrigger) > 10*time.Second {
	fmt.Printf(" [!] TRIGGERING CLEANUP...")
	cmd := exec.Command("powershell", "-Command", "Write-Host 'RAM Cleared by Narcissus")
	cmd.Run()

	lastTrigger = time.Now()
}
}
fmt.Printf("\r%s Current RAM Usage: %d KB %s \033[K", pulse, used, status)
lastUsed = used
time.Sleep(1 * time.Second)
}
}