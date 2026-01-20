package shellcode

import "workshop3_dev/internals/models"

// TODO: Define CommandShellcode interface for shellcode execution
// This interface allows different implementations for different platforms (Windows, Mac, etc.)
// Hint: It should have one method:
//   DoShellcode(dllBytes []byte, exportName string) (models.ShellcodeResult, error)
type CommandShellcode interface {
	DoShellcode(dllBytes []byte, exportName string) (models.ShellcodeResult, error)
}
