package cli_app

import (
	"log"
)

func main() {
	if err := CLI_App.cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
