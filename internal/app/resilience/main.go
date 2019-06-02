/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"github.com/getlantern/systray"
)

func main() {
	go func() {
		updateHosts(false)
		denierHostsInit()
		denierProxyInit()
		tickersInit()
	}()
	systray.Run(guiOnReady, guiOnExit)
}
