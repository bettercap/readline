// +build windows

package readline

import (
	"os"
	"os/exec"
)

func init() {
	//no buffering
	exec.Command("stty", "cbreak", "min", "1").Run()
	//no visible output
	exec.Command("stty", "-echo").Run()

	Stdin = os.Stdin
	Stdout = os.Stdout
	Stderr = os.Stderr
}
