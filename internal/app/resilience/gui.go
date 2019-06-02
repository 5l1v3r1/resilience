/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/kaepora/go-autostart"
)

type guiTrayControls struct {
	status    *systray.MenuItem
	toggle    *systray.MenuItem
	update    *systray.MenuItem
	help      *systray.MenuItem
	about     *systray.MenuItem
	autostart *systray.MenuItem
	quit      *systray.MenuItem
}

var guiHelpURI = "https://resilienceblocker.info/#help"

func guiOnReady() {
	guiAutoStart := guiBuildAutoStart()
	guiTray := guiBuildTray(guiAutoStart)
	guiTrayMonitor(guiTray, guiAutoStart)
}

func guiBuildAutoStart() *autostart.App {
	var appExec []string
	exePath, _ := os.Executable()
	switch runtime.GOOS {
	case "windows":
		// Tested.
		appExec = []string{exePath}
	case "linux":
		// TODO: Untested.
		appExec = []string{exePath}
	case "darwin":
		// TODO: Untested.
		appExec = []string{exePath}
	}
	return &autostart.App{
		Name:        "Resilience",
		DisplayName: "Resilience",
		Exec:        appExec,
	}
}

func guiBuildTray(guiAutoStart *autostart.App) *guiTrayControls {
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
	guiTray.autostart = systray.AddMenuItem("Start on Login", "Automatically Start on Login.")
	systray.AddSeparator()
	guiTray.quit = systray.AddMenuItem("Quit", "Quit Resilience.")
	if guiAutoStart.IsEnabled() {
		guiTray.autostart.Check()
	}
	return &guiTray
}

func guiTrayMonitor(guiTray *guiTrayControls, guiAutoStart *autostart.App) {
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
			case <-guiTray.autostart.ClickedCh:
				guiTrayMonitorOnAutoStart(guiTray, guiAutoStart)
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
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", guiHelpURI).Start()
	case "linux":
		exec.Command("xdg-open", guiHelpURI).Start()
	case "darwin":
		exec.Command("open", guiHelpURI).Start()
	}
}

func guiTrayMonitorOnAbout(guiTray *guiTrayControls) {
	aboutInfo()
}

func guiTrayMonitorOnAutoStart(guiTray *guiTrayControls, guiAutoStart *autostart.App) {
	if guiAutoStart.IsEnabled() {
		guiAutoStart.Disable()
		guiTray.autostart.Uncheck()
	} else {
		guiAutoStart.Enable()
		guiTray.autostart.Check()
	}
}

func guiTrayMonitorOnQuit(guiTray *guiTrayControls) {
	systray.Quit()
}

func guiOnExit() {}
