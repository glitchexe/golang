package main

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/registry"
)

func getMachineGUID() string {
	var key string
	if runtime.GOARCH == "amd64" {
		log.Debug("Found 64bit OS...")
		key = `SOFTWARE\Microsoft\Cryptography`
	} else if runtime.GOARCH == "386" {
		log.Debug("Found 32bit OS...")
		key = `SOFTWARE\WOW6432Node\Microsoft\Cryptography`
	}
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("MachineGuid")
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func main() {
	// Print the Machine GUID from the Windows Registry
	fmt.Printf(getMachineGUID())
}
