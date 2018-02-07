// +build windows

package readline

import "os"

func init() {
	Stdin = os.Stdin
	Stdout = NewANSIWriter(Stdout)
	Stderr = NewANSIWriter(Stderr)
}
