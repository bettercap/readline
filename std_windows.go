// +build windows

package readline

import "os"

func init() {
	Stdin = os.Stdin
	Stdout = NewANSIWriter(os.Stdout)
	Stderr = NewANSIWriter(os.Stderr)
}
