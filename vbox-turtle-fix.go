package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func isAdmin() bool {

	cmd := exec.Command("net", "session")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err := cmd.Run()

	return err == nil
}

func elevate() {

	fmt.Println("Requesting Administrator privileges...")

	exe, _ := os.Executable()

	cmd := exec.Command("powershell",
		"Start-Process",
		"\""+exe+"\"",
		"-Verb",
		"RunAs")

	cmd.Run()

	os.Exit(0)
}

func run(name string, args ...string) {

	fmt.Println(">", name, strings.Join(args, " "))

	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err := cmd.Run()

	if err != nil {

		fmt.Println("Error:", err)
	}
}

func runHidden(name string, args ...string) {

	cmd := exec.Command(name, args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	cmd.Run()
}

func runPowerShell(command string) string {

	cmd := exec.Command("powershell",
		"-Command",
		command)

	out, _ := cmd.Output()

	return string(out)
}

func main() {

	if !isAdmin() {

		elevate()
		return
	}

	fmt.Println("================================")
	fmt.Println(" VirtualBox Turtle Fix Tool")
	fmt.Println("================================")

	fmt.Println()

	// Stop services

	runHidden("sc", "stop", "vmms")
	runHidden("sc", "stop", "vmcompute")
	runHidden("sc", "stop", "hvhost")
	runHidden("sc", "stop", "vmmem")

	// Disable hypervisor

	run("bcdedit", "/set", "hypervisorlaunchtype", "off")

	run("bcdedit", "/set", "vsmlaunchtype", "off")

	// Disable features

	features := []string{

		"Microsoft-Hyper-V-All",
		"Microsoft-Hyper-V",
		"Microsoft-Hyper-V-Hypervisor",
		"VirtualMachinePlatform",
		"Microsoft-Windows-Subsystem-Linux",
		"HypervisorPlatform",
		"Containers",
		"Windows-Defender-ApplicationGuard",
	}

	for _, f := range features {

		run("dism",
			"/online",
			"/disable-feature",
			"/featurename:"+f,
			"/norestart")
	}

	// Registry fixes

	run("reg", "add",
		`HKLM\SYSTEM\CurrentControlSet\Control\DeviceGuard`,
		"/v",
		"EnableVirtualizationBasedSecurity",
		"/t",
		"REG_DWORD",
		"/d",
		"0",
		"/f")

	run("reg", "add",
		`HKLM\SYSTEM\CurrentControlSet\Control\DeviceGuard`,
		"/v",
		"RequirePlatformSecurityFeatures",
		"/t",
		"REG_DWORD",
		"/d",
		"0",
		"/f")

	run("reg", "add",
		`HKLM\SYSTEM\CurrentControlSet\Control\DeviceGuard\Scenarios\HypervisorEnforcedCodeIntegrity`,
		"/v",
		"Enabled",
		"/t",
		"REG_DWORD",
		"/d",
		"0",
		"/f")

	run("reg", "add",
		`HKLM\SYSTEM\CurrentControlSet\Control\Lsa`,
		"/v",
		"LsaCfgFlags",
		"/t",
		"REG_DWORD",
		"/d",
		"0",
		"/f")

	// Remove WSL distros

	run("wsl", "--shutdown")

	list := runPowerShell("wsl -l -q")

	scanner := bufio.NewScanner(strings.NewReader(list))

	for scanner.Scan() {

		d := strings.TrimSpace(scanner.Text())

		if d != "" {

			run("wsl", "--unregister", d)
		}
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println(" FIX COMPLETE")
	fmt.Println(" REBOOT REQUIRED")
	fmt.Println("================================")

	fmt.Println()
	fmt.Println("Press ENTER to exit")

	fmt.Scanln()
}
