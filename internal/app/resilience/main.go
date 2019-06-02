/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

//go:generate goversioninfo -64=true -icon=../../../assets/icon/icon.ico
package main

import (
	"github.com/getlantern/systray"
)

func main() {
	tickersInit()
	go func() {
		updateHosts(false)
		denierHostsInit()
		denierProxyInit()
	}()
	systray.Run(guiOnReady, guiOnExit)
}
