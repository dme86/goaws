package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/dme86/goaws/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		if viper.GetBool("debug") {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		os.Exit(1)
	}
}
