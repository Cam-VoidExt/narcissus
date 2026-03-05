# Narcissus.exe 🪞

**A Go-based study in digital self-obsession.**

Narcissus is a lightweight systems engineering tool designed with one goal: to monitor itself with unyielding devotion. While other programs worry about system health or "Service Host" hogs, Narcissus only has eyes for its own internal heap.

## 🛠 Features
* **Self-Instrumentation:** Uses `runtime.MemStats` to monitor its own memory footprint in real-time.
* **Heuristic Feedback Loop:** If Narcissus detects it is "growing" (RAM spike), it triggers an autonomous cleanup via PowerShell.
* **Persistent Vanity:** Logs every "growth spurt" to a local `audit.log` for future generations to admire.
* **Ghost Killer UI:** Implements ANSI escape sequences (`\033[K`) to ensure the terminal output remains as flawless as its reflection.

## 🧠 Lessons Learned
* **The Power of Go:** Built to run as a single, static binary with a footprint under 1MB.
* **Feedback Loops:** Solved the "cleaner bringing in trash" problem by implementing a time-based debounce (cooldown) on system triggers.
* **Logic vs. Appearance:** Learned that terminal buffers are messy; professional tools require surgical cursor control.

## ⚠️ Safety Warning
Do not run Narcissus near a mirror. If the program discovers its own source code, it may fall in love and refuse to exit. This is a feature, not a bug.

---
*Created by Void-Ext*
