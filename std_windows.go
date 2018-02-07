// +build windows

package readline

import "os"

func init() {
	Stdin = os.Stdin
	Stdout = os.Stdout
	Stderr = os.Stderr
}
