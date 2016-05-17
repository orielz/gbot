package helper

import (
	gconfig "gbot/config"
	"math/rand"
	"os/exec"
	"time"
)

// GetRandomName returns random name
func GetRandomName() string {
	return gconfig.Nicks[random(0, len(gconfig.Nicks))]
}

// GetRandomServer returns random server
func GetRandomServer() string {
	return gconfig.Servers[random(0, len(gconfig.Servers))]
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// IsAdmin determine if the source is admin
func IsAdmin(src string) bool {

	for _, admin := range gconfig.Admins {
		if admin == src {
			return true
		}
	}

	return false
}

// RemoveFromSlice removes string from strings slice
func RemoveFromSlice(slice []string, element string) []string {

	var index = 0
	var found = false

	for i, e := range slice {
		if element == e {
			found = true
			index = i
		}
	}

	if found {
		return append(slice[:index], slice[index+1:]...)
	}

	return slice

}

// Run cmd command
func Run(cmd string) {
	c := exec.Command("cmd", "/C", cmd)

	if err := c.Run(); err != nil {
		//fmt.Println("Error: ", err)
	}
}

// RandStringBytes return random string
func RandStringBytes(n int) string {

	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
