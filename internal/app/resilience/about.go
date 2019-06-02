/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

import (
	"github.com/sqweek/dialog"
)

func aboutInfo() {
	dialog.Message(stateX.locale.aboutInfoText).Title(stateX.locale.aboutInfoTitle).Info()
}
