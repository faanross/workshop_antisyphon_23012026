//go:build darwin

package shellcode

import (
	"fmt"
	"workshop3_dev/internals/models"
)

// TODO: Implement macShellcode struct that implements CommandShellcode interface
// This is a platform-specific implementation for macOS
// Note the build tag //go:build darwin at the top - this file only compiles on macOS

// TODO: Implement New() constructor that returns CommandShellcode interface
func New() CommandShellcode {

}

// TODO: Implement DoShellcode for macOS
// On macOS, this feature is not implemented yet, so we just return a message
func (ms *macShellcode) DoShellcode(dllBytes []byte, exportName string) (models.ShellcodeResult, error) {
	fmt.Println("|SHELLCODE DOER MACOS| This feature has not yet been implemented for MacOS.")
	
	return result, nil
}
