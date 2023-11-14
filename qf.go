package main

import (
	"fmt"
	"os"

	"github.com/adrianguyareach/FIXlib/lib/tradeclient"
	"github.com/adrianguyareach/FIXlib/lib/utils"
)

func main() {

	err := tradeclient.Client(os.Args)
	if err != nil {

		message := fmt.Sprintf("Failure starting application: %s", err)
		utils.PrintBad(message)
	}
	select {}

}
