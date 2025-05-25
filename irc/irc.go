package irc

import (
	"fmt"
	gconfig "gbot/config"
	ghelper "gbot/helper"
	"strings"

	irc "github.com/fluffle/goirc/client"
)
// Start returns random name
func Start() {

	for {

		// Create config struct and init defaults
		config := initConfig()
		client := irc.Client(config)

		quit := make(chan bool)

		client.HandleFunc(irc.QUIT, func(conn *irc.Conn, line *irc.Line) { quit <- true })
		client.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) { quit <- true })
		client.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
			gconfig.Admins = gconfig.Admins[:0]
			conn.Join(gconfig.Channel)
		})
		client.HandleFunc(irc.PART, func(conn *irc.Conn, line *irc.Line) {
			if ghelper.IsAdmin(line.Src) {
				gconfig.Admins = ghelper.RemoveFromSlice(gconfig.Admins, line.Src)
			}
		})

		client.HandleFunc(irc.KICK, func(conn *irc.Conn, line *irc.Line) {

			var s = strings.Split(line.Raw, " ")
			var row = s[len(s)-1]
			var kicked = strings.Replace(strings.Split(row, "[")[1], "]", "", -1)

			if ghelper.IsAdmin(kicked) {
				gconfig.Admins = ghelper.RemoveFromSlice(gconfig.Admins, kicked)
			}

		})

		client.HandleFunc(irc.NICK, func(conn *irc.Conn, line *irc.Line) {

			if ghelper.IsAdmin(line.Src) {

				s := strings.Split(line.Raw, ":")
				var newNick = s[len(s)-1]

				var oldSrc = strings.Replace(s[1], " NICK", " ", -1)
				oldSrc = strings.Replace(oldSrc, " ", "", -1)

				var newSrc = newNick + "!" + strings.Split(oldSrc, "!")[1]
				newSrc = strings.Replace(newSrc, " ", "", -1)

				gconfig.Admins = ghelper.RemoveFromSlice(gconfig.Admins, oldSrc)
				gconfig.Admins = append(gconfig.Admins, newSrc)
			}

		})

		RegisterCommands(client)

		if err := client.Connect(); err != nil {
			fmt.Printf("Connection error: %s\n", err.Error())
		}

		<-quit

	}

}

func initConfig() *irc.Config {
	config := irc.NewConfig(ghelper.GetRandomName() + ghelper.RandStringBytes(1))
	config.Server = ghelper.GetRandomServer()
	config.NewNick = func(n string) string { return n + "^" }
	config.Me.Ident = ghelper.GetRandomName()
	config.Me.Name = ghelper.GetRandomName()

	return config
}
