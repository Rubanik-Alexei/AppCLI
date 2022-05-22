package cli_app

import (
	"log"

	"github.com/Rubanik-Alexei/CLI_App/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
