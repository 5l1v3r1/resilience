/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"runtime/debug"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	go func() {
		for range time.NewTicker(24 * time.Hour).C {
			updateClient(false)
			updateHosts(false)
		}
	}()
	go func() {
		for range time.NewTicker(60 * time.Second).C {
			debug.FreeOSMemory()
		}
	}()
	go func() {
		updateHosts(false)
		denierHostsInit()
		denierProxyInit()
	}()
	systray.Run(guiOnReady, guiOnExit)
}
