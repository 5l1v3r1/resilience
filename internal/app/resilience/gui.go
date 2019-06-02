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
		Name:        stateX.locale.resilience,
		DisplayName: stateX.locale.resilience,
		Exec:        appExec,
	}
}

func guiBuildTray(guiAutoStart *autostart.App) *guiTrayControls {
	var guiTray guiTrayControls
	systray.SetIcon(iconData)
	if runtime.GOOS != "darwin" {
		systray.SetTitle(stateX.locale.resilience)
	}
	systray.SetTooltip(stateX.locale.resilience)
	guiTray.status = systray.AddMenuItem(stateX.locale.enabled, "")
	guiTray.status.Disable()
	guiTray.toggle = systray.AddMenuItem(stateX.locale.enable, "")
	systray.AddSeparator()
	guiTray.update = systray.AddMenuItem(stateX.locale.update, "")
	systray.AddSeparator()
	guiTray.help = systray.AddMenuItem(stateX.locale.gettingStarted, "")
	guiTray.about = systray.AddMenuItem(stateX.locale.about, "")
	systray.AddSeparator()
	guiTray.autostart = systray.AddMenuItem(stateX.locale.autoStart, "")
	systray.AddSeparator()
	guiTray.quit = systray.AddMenuItem(stateX.locale.quit, "")
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
	if stateX.enabled {
		err := togglerDisable()
		if err == nil {
			guiTray.status.SetTitle(stateX.locale.disabled)
			guiTray.toggle.SetTitle(stateX.locale.enable)
		}
	} else {
		err := togglerEnable()
		if err == nil {
			guiTray.status.SetTitle(stateX.locale.enabled)
			guiTray.toggle.SetTitle(stateX.locale.disabled)
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
