/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

func togglerEnable() error {
	stateX.enabled = true
	return nil
}

func togglerDisable() error {
	stateX.enabled = false
	return nil
}
