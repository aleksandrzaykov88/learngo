package main

import (
	"fmt"
	"strings"
)

//ParseIni parses .ini-file.
func ParseIni(s []string) map[string]map[string]string {
	result := make(map[string]map[string]string)
	key := ""
	for _, line := range s {
		if strings.HasPrefix(line, ";") || line == "" {
			continue
		}
		if strings.HasPrefix(line, "[") {
			key = line[1 : len(line)-1]
			result[key] = make(map[string]string)
			continue
		}
		sLine := strings.Split(line, "=")
		result[key][sLine[0]] = sLine[1]
	}
	return result
}

func main() {
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
	fmt.Println(ParseIni(iniData))
}
