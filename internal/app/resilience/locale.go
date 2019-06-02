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
	fr locale
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
	fr: locale{
		resilience:     "Resilience",
		aboutInfoTitle: "À propos de Resilience",
		aboutInfoText: strings.Join([]string{
			"Resilience " + versionString + "\n",
			"Resilience est un bloqueur de contenu facile à utiliser pour votre ordinateur.",
			"Pour des nouvelles et de l'information, veuillez visiter le site :",
			"https://resilienceblocker.info",
		}, "\n"),
		errorTitle:                  "Erreur Resilience",
		denierUpdateErrorText:       "Impossible de mettre à jour votre liste de blocage Resilience.",
		denierHostsErrorText:        "Impossible de lire ou d'écrire dans votre liste de blocage de Resilience locale.",
		enabled:                     "Resilience Activé",
		disabled:                    "Resilience Désactivé",
		enable:                      "Activer",
		disable:                     "Désactiver",
		update:                      "Mettre a Jour",
		gettingStarted:              "Aide",
		about:                       "À Propos",
		autoStart:                   "Démarrer Automatiquement",
		quit:                        "Arrêt",
		updateInfoTitle:             "Mise à Jour Resilience",
		updateHostsErrorText:        "Impossible de mettre à jour votre liste de blocage Resilience.",
		updateHostsNoUpdateInfoText: "Aucune mise à jour n'est disponible pour votre liste de blocage Resilience.",
		updateClientErrorText:       "Impossible de vérifier les mises à jour pour Resilience.",
		updateClientHasUpdateInfoText: strings.Join([]string{
			"Une mise à jour est disponible pour votre client Resilience.\n",
			"Pour télécharger la dernière version, veuillez visiter le site :",
			"https://resilienceblocker.info",
		}, "\n"),
		updateClientNoUpdateInfoText: "Aucune mise à jour n'est disponible pour votre client Resilience.",
	},
	de: locale{
		resilience:     "Resilience",
		aboutInfoTitle: "Über Resilience",
		aboutInfoText: strings.Join([]string{
			"Resilience " + versionString + "\n",
			"Resilience ist ein einfach zu benutzender Content Blocker für Ihren Computer.",
			"Für Neues und Informationen besuchen Sie bitte:",
			"https://resilienceblocker.info",
		}, "\n"),
		errorTitle:                  "Resilience Error",
		denierUpdateErrorText:       "Konnte Ihre Resilience Block-Liste nicht aktualisieren.",
		denierHostsErrorText:        "Konnte nicht auf Ihre lokale Resilience Block-List zugreifen.",
		enabled:                     "Resilience ist aktiviert",
		disabled:                    "Resilience ist deaktiviert",
		enable:                      "Aktivieren",
		disable:                     "Deaktivieren",
		update:                      "Aktualisieren",
		gettingStarted:              "Erste Schritte",
		about:                       "Über",
		autoStart:                   "Beim Anmelden starten",
		quit:                        "Beenden",
		updateInfoTitle:             "Resilience Update",
		updateHostsErrorText:        "Konnte die Resilience Block-Liste nicht aktualisieren.",
		updateHostsNoUpdateInfoText: "Für die Resilience Block-Liste sind keine Updates verfügbar.",
		updateClientErrorText:       "Konnte nicht nach Updates für Resilience suchen.",
		updateClientHasUpdateInfoText: strings.Join([]string{
			"An update is available for your Resilience client.\n",
			"To download the latest version, please visit:",
			"https://resilienceblocker.info",
		}, "\n"),
		updateClientNoUpdateInfoText: "Für Resilience sind keine Updates verfügbar.",
	},
}
