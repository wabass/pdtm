package runner

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
    updateutils "github.com/projectdiscovery/utils/update"
)

const version = "v0.0.4"

var banner = fmt.Sprintf(`
                ____          
     ____  ____/ / /_____ ___ 
    / __ \/ __  / __/ __ __  \
   / /_/ / /_/ / /_/ / / / / /
  / .___/\__,_/\__/_/ /_/ /_/ 
 /_/                          %s
`, version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tprojectdiscovery.io\n\n")
}

// GetUpdateCallback returns a callback function that updates pdtm
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("pdtm", version)()
	}
}
