/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"strings"

	"github.com/sqweek/dialog"
)

func aboutInfo() {
	var aboutText = strings.Join([]string{
		"Resilience " + versionString + "\n",
		"Resilience is an easy to use content blocker for your computer.",
		"For news and information, please visit:",
		"https://resilienceblocker.info",
	}, "\n")
	dialog.Message(aboutText).Title("About Resilience").Info()
}
