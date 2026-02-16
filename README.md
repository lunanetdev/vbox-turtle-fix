# VirtualBox Turtle Fix (Windows)

Fix the 🐢 **VirtualBox turtle icon** by automatically disabling Hyper-V, Virtualization-Based Security (VBS), and related Windows features.

This restores **full native hardware virtualization performance** (VT-x / AMD-V) for Oracle VM VirtualBox.

---

## Why this exists

When Hyper-V or VBS is enabled, VirtualBox runs under the Windows hypervisor instead of directly on the CPU.

This causes:

* 🐢 Turtle icon in VirtualBox
* Very slow virtual machines
* High CPU usage
* Poor responsiveness

This tool fixes the problem automatically.

---

## Features

* Automatic Administrator elevation (no right-click required)
* Fully automatic operation
* Disables Hyper-V
* Disables VBS / Credential Guard
* Removes WSL
* Restores native virtualization performance
* No installation required
* Portable single `.exe`
* No external dependencies

---

## Download

Download the latest version from:

**Releases → vbox-turtle-fix.exe**


---

## Usage

1. Download `vbox-turtle-fix.exe`
2. Double-click to run
3. Accept the Administrator prompt
4. Wait for completion
5. Reboot Windows

Done.

---

## Important

Administrator rights are required.

The application will request elevation automatically if needed.

---

## What this tool disables

This tool disables:

* Hyper-V
* Virtual Machine Platform
* Windows Subsystem for Linux (WSL)
* Hypervisor Platform
* Windows Sandbox
* Credential Guard
* Virtualization-Based Security

---

## Warning

After running this tool, these features will not work:

* WSL
* Docker Desktop (WSL backend)
* Hyper-V virtual machines
* Windows Sandbox

They can be restored manually if needed.

---

## Safety

This tool:

* Does NOT install anything
* Does NOT download anything
* Does NOT connect to internet
* Uses only built-in Windows tools

Fully open source.

---

## Reboot required

Changes take effect after reboot.

---

## Disclaimer

Use at your own risk.

This tool modifies Windows system configuration.

---
