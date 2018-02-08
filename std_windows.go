// +build windows

package readline

import "os"

func init() {
	Stdin = NewRawReader() // os.Stdin
	Stdout = os.Stdout
	Stderr = os.Stderr
}
