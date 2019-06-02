/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sqweek/dialog"
)

type updateData struct {
	Latest int
}

var updateHostsB2sumURI = "https://resilienceblocker.info/data/blockList.b2sum"
var updateHostsURI = "https://resilienceblocker.info/data/blockList"
var updateClientJSONURI = "https://resilienceblocker.info/data/updateClient.json"

func updateHosts(explicit bool) error {
	var httpClient = &http.Client{Timeout: 60 * time.Second}
	var err error
	r, err := httpClient.Get(updateHostsB2sumURI)
	if err != nil {
		updateHostsError()
		return err
	}
	defer r.Body.Close()
	b2sum, err := ioutil.ReadAll(r.Body)
	if err != nil {
		updateHostsError()
		return err
	}
	hostsHash := strings.Trim(string(b2sum), "\r\n ")
	if stateState.hostsHash == hostsHash {
		if explicit {
			updateHostsNoUpdateInfo()
		}
		return err
	}
	r, err = httpClient.Get(updateHostsURI)
	if err != nil {
		updateHostsError()
		return err
	}
	defer r.Body.Close()
	hosts, err := ioutil.ReadAll(r.Body)
	return denierUpdate(hosts, true)
}

func updateClient(explicit bool) error {
	var updateResult updateData
	var httpClient = &http.Client{Timeout: 20 * time.Second}
	var err error
	r, err := httpClient.Get(updateClientJSONURI)
	if err != nil {
		updateClientError()
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		updateClientError()
		return err
	}
	err = json.Unmarshal(body, &updateResult)
	if err != nil {
		updateClientError()
		return err
	}
	if updateResult.Latest > versionBuild {
		updateClientHasUpdateInfo()
	} else {
		if explicit {
			updateClientNoUpdateInfo()
		}
	}
	return err
}

func updateHostsError() {
	dialog.Message(
		"Could not update your Resilience block list.",
	).Title("Resilience Update").Error()
}

func updateHostsNoUpdateInfo() {
	dialog.Message(
		"No updates are available for your Resilience block list.",
	).Title("Resilience Update").Info()
}

func updateClientError() {
	dialog.Message(
		"Could not check for updates for Resilience.",
	).Title("Resilience Update").Error()
}

func updateClientHasUpdateInfo() {
	dialog.Message(strings.Join([]string{
		"An update is available for your Resilience client.\n",
		"To download the latest version, please visit:",
		"https://resilienceblocker.info",
	}, "\n")).Title("Resilience Update").Info()
}

func updateClientNoUpdateInfo() {
	dialog.Message(
		"No updates are available for your Resilience client.",
	).Title("Resilience Update").Info()
}
