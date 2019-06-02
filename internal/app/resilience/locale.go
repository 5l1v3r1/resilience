/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

import "strings"

type locale struct {
	resilience                    string
	aboutInfoTitle                string
	aboutInfoText                 string
	errorTitle                    string
	denierUpdateErrorText         string
	denierHostsErrorText          string
	enabled                       string
	disabled                      string
	enable                        string
	disable                       string
	update                        string
	gettingStarted                string
	about                         string
	autoStart                     string
	quit                          string
	updateInfoTitle               string
	updateHostsErrorText          string
	updateHostsNoUpdateInfoText   string
	updateClientErrorText         string
	updateClientHasUpdateInfoText string
	updateClientNoUpdateInfoText  string
}

type locales struct {
	en locale
}

var localeText = locales{
	en: locale{
		resilience:     "Resilience",
		aboutInfoTitle: "About Resilience",
		aboutInfoText: strings.Join([]string{
			"Resilience " + versionString + "\n",
			"Resilience is an easy to use content blocker for your computer.",
			"For news and information, please visit:",
			"https://resilienceblocker.info",
		}, "\n"),
		errorTitle:                  "Resilience Error",
		denierUpdateErrorText:       "Could not update your Resilience block list.",
		denierHostsErrorText:        "Could not read or write to your local Resilience block list.",
		enabled:                     "Resilience is Enabled",
		disabled:                    "Resilience is Disabled",
		enable:                      "Enable",
		disable:                     "Disable",
		update:                      "Update",
		gettingStarted:              "Getting Started",
		about:                       "About",
		autoStart:                   "Start on Login",
		quit:                        "Quit",
		updateInfoTitle:             "Resilience Update",
		updateHostsErrorText:        "Could not update your Resilience block list.",
		updateHostsNoUpdateInfoText: "No updates are available for your Resilience block list.",
		updateClientErrorText:       "Could not check for updates for Resilience.",
		updateClientHasUpdateInfoText: strings.Join([]string{
			"An update is available for your Resilience client.\n",
			"To download the latest version, please visit:",
			"https://resilienceblocker.info",
		}, "\n"),
		updateClientNoUpdateInfoText: "No updates are available for your Resilience client.",
	},
}
