package main

import (
	"fmt"
	"log"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func main() {
	var key string
	if runtime.GOARCH == "amd64" {
		fmt.Printf("Found 64bit OS...\n")
		key = `SOFTWARE\Microsoft\Cryptography`
	} else if runtime.GOARCH == "386" {
		fmt.Printf("Found 32bit OS...")
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
	fmt.Printf("Windows system GUID is %q\n", s)
}
