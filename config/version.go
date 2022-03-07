package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Paramsflag struct {
	Version bool `json:"version"`
}

var Pflag Paramsflag

func DoParams() {
	pflag.BoolVarP(&Pflag.Version, "version", "v", false, "show version")
	pflag.Parse()
	if Pflag.Version {
		fmt.Printf("Version: %s", Version)
		os.Exit(0)
	}
}
