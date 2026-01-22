//go:build darwin

package shellcode

import (
	"fmt"
	"workshop3_dev/internals/models"
)

// Note the build tag //go:build darwin at the top - this file only compiles on macOS

// TODO: Implement macShellcode struct that implements CommandShellcode interface

// TODO: Implement New() constructor that returns CommandShellcode interface
func New() CommandShellcode {

}

// On macOS, this feature is not implemented yet, so we just return a message
func (ms *macShellcode) DoShellcode(dllBytes []byte, exportName string) (models.ShellcodeResult, error) {
	fmt.Println("|SHELLCODE DOER MACOS| This feature has not yet been implemented for MacOS.")

	// TODO: create result struct with failure Message
	return result, nil
}
