package main

import "fmt"

//ParseIni parses .ini-file.
func ParseIni() map[string]map[string]string {
	return map[string]map[string]string{"Map": {"Key": "Value"}}
}

func main() {
	fmt.Println(ParseIni())
	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
}
