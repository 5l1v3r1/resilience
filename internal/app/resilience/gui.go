/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
)

type guiTrayControls struct {
	status *systray.MenuItem
	toggle *systray.MenuItem
	update *systray.MenuItem
	help   *systray.MenuItem
	about  *systray.MenuItem
	quit   *systray.MenuItem
}

func guiOnReady() {
	guiTray := guiBuild()
	guiTrayMonitor(&guiTray)
}

func guiBuild() guiTrayControls {
	var guiTray guiTrayControls
	systray.SetIcon(iconData)
	if runtime.GOOS != "darwin" {
		systray.SetTitle("Resilience")
	}
	systray.SetTooltip("Resilience")
	guiTray.status = systray.AddMenuItem("Resilience is Enabled", "")
	guiTray.status.Disable()
	guiTray.toggle = systray.AddMenuItem("Disable", "")
	systray.AddSeparator()
	guiTray.update = systray.AddMenuItem("Update", "Check for Updates.")
	systray.AddSeparator()
	guiTray.help = systray.AddMenuItem("Getting Started", "Help with Resilience.")
	guiTray.about = systray.AddMenuItem("About", "About Resilience.")
	systray.AddSeparator()
	guiTray.quit = systray.AddMenuItem("Quit", "Quit Resilience.")
	return guiTray
}

func guiTrayMonitor(guiTray *guiTrayControls) {
	go func() {
		for {
			select {
			case <-guiTray.toggle.ClickedCh:
				guiTrayMonitorOnToggle(guiTray)
			case <-guiTray.update.ClickedCh:
				guiTrayMonitorOnUpdate(guiTray)
			case <-guiTray.help.ClickedCh:
				guiTrayMonitorOnHelp(guiTray)
			case <-guiTray.about.ClickedCh:
				guiTrayMonitorOnAbout(guiTray)
			case <-guiTray.quit.ClickedCh:
				guiTrayMonitorOnQuit(guiTray)
			}
		}
	}()
}

func guiTrayMonitorOnToggle(guiTray *guiTrayControls) {
	if stateState.enabled {
		err := togglerDisable()
		if err == nil {
			guiTray.status.SetTitle("Resilience is Disabled")
			guiTray.toggle.SetTitle("Enable")
		}
	} else {
		err := togglerEnable()
		if err == nil {
			guiTray.status.SetTitle("Resilience is Enabled")
			guiTray.toggle.SetTitle("Disable")
		}
	}
}

func guiTrayMonitorOnUpdate(guiTray *guiTrayControls) {
	go func() {
		updateHosts(false)
		updateClient(true)
	}()
}

func guiTrayMonitorOnHelp(guiTray *guiTrayControls) {
	helpPage := "https://resilienceblocker.info/#help"
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", helpPage).Start()
	case "linux":
		exec.Command("xdg-open", helpPage).Start()
	case "darwin":
		exec.Command("open", helpPage).Start()
	}
}

func guiTrayMonitorOnAbout(guiTray *guiTrayControls) {
	aboutInfo()
}

func guiTrayMonitorOnQuit(guiTray *guiTrayControls) {
	systray.Quit()
}

func guiOnExit() {}
