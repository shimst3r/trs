/*Package main provides the entry point to all (sub-) commands.

trs is modeled using a very simple state machine:

0. init state
1. start state
2. stop state

The init state is created when calling `trs init` for the first time.

To go from 0 to 1, a user has to call `trs start`, which starts the timer. Calling `trs start`
again while in start state, nothing will happen.

To go from 1 to 2, as user has to call `trs stop`, which stops the timer. Calling `trs stop`
again while in stop state, nothing will happen.

To go from 2 to 1 again, a user has to call `trs start`, which starts the timer again.
*/
package main

import (
	"fmt"
	"os"

	"github.com/shimst3r/trs/trs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: trs needs a subcommand, run \"trs help\" for more information.")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "help":
		trs.Help()
	case "init":
		trs.Init()
	case "start":
		trs.Start()
	case "stop":
		trs.Stop()
	case "today":
		trs.Today()
	default:
		fmt.Printf("Error: trs does not support subcommand %s, run \"trs help\" for more information.\n", os.Args[1])
		os.Exit(1)
	}
}
