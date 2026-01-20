//go:build windows

package shellcode

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"runtime"
	"syscall"
	"unsafe"
	"workshop3_dev/internals/models"

	"golang.org/x/sys/windows"
)

// TODO: Define PE Structures for Windows DLL loading
// These structures map to the Windows PE format headers

// IMAGE_DOS_HEADER - First structure in any PE file
// Hint: Contains Magic (uint16) for "MZ" signature and Lfanew (int32) pointing to NT headers
type IMAGE_DOS_HEADER struct {
	Magic  uint16
	_      [58]byte
	Lfanew int32
}

// TODO: Define IMAGE_FILE_HEADER
// Hint: Contains Machine, NumberOfSections, SizeOfOptionalHeader, Characteristics
type IMAGE_FILE_HEADER struct {
	// Add fields here
}

// IMAGE_DATA_DIRECTORY - Entry in the DataDirectory array
type IMAGE_DATA_DIRECTORY struct{ VirtualAddress, Size uint32 }

// TODO: Define IMAGE_OPTIONAL_HEADER64 for 64-bit PE files
// Hint: Contains Magic, ImageBase (uint64), SizeOfImage, SizeOfHeaders, AddressOfEntryPoint
// and DataDirectory array of 16 IMAGE_DATA_DIRECTORY entries
type IMAGE_OPTIONAL_HEADER64 struct {
	// Add fields here - this is a large structure with many fields
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

// TODO: Define IMAGE_SECTION_HEADER for PE sections
// Hint: Contains Name [8]byte, VirtualSize, VirtualAddress, SizeOfRawData, PointerToRawData
type IMAGE_SECTION_HEADER struct {
	// Add fields here
}

// IMAGE_BASE_RELOCATION - Base relocation block header
type IMAGE_BASE_RELOCATION struct{ VirtualAddress, SizeOfBlock uint32 }

// IMAGE_IMPORT_DESCRIPTOR - Import directory entry
type IMAGE_IMPORT_DESCRIPTOR struct{ OriginalFirstThunk, TimeDateStamp, ForwarderChain, Name, FirstThunk uint32 }

// TODO: Define IMAGE_EXPORT_DIRECTORY for finding exported functions
// Hint: Contains NumberOfFunctions, NumberOfNames, AddressOfFunctions, AddressOfNames, AddressOfNameOrdinals
type IMAGE_EXPORT_DIRECTORY struct {
	// Add fields here
}

// --- Constants for PE loading ---
const (
	IMAGE_DIRECTORY_ENTRY_EXPORT    = 0
	DLL_PROCESS_ATTACH              = 1
	IMAGE_DOS_SIGNATURE             = 0x5A4D     // "MZ"
	IMAGE_NT_SIGNATURE              = 0x00004550 // "PE\0\0"
	IMAGE_DIRECTORY_ENTRY_BASERELOC = 5
	IMAGE_DIRECTORY_ENTRY_IMPORT    = 1
	IMAGE_REL_BASED_DIR64           = 10
	IMAGE_REL_BASED_ABSOLUTE        = 0
	IMAGE_ORDINAL_FLAG64            = uintptr(1) << 63
	MEM_COMMIT                      = 0x00001000
	MEM_RESERVE                     = 0x00002000
	MEM_RELEASE                     = 0x8000
	PAGE_READWRITE                  = 0x04
	PAGE_EXECUTE_READWRITE          = 0x40
)

// Global Proc Address Loader
var (
	kernel32DLL        = windows.NewLazySystemDLL("kernel32.dll")
	procGetProcAddress = kernel32DLL.NewProc("GetProcAddress")
)

// Helper function to convert section name bytes to string
func sectionNameToString(nameBytes [8]byte) string {
	n := bytes.IndexByte(nameBytes[:], 0)
	if n == -1 {
		n = 8
	}
	return string(nameBytes[:n])
}

// windowsShellcode implements the CommandShellcode interface for Windows.
type windowsShellcode struct{}

// New is the constructor for our Windows-specific Shellcode command
func New() CommandShellcode {
	return &windowsShellcode{}
}

// DoShellcode loads and runs the given DLL bytes in the current process.
func (rl *windowsShellcode) DoShellcode(
	dllBytes []byte,
	exportName string,
) (models.ShellcodeResult, error) {

	fmt.Println("|✅ SHELLCODE DOER| The SHELLCODE command has been executed.")

	// Basic validation
	if runtime.GOOS != "windows" {
		return models.ShellcodeResult{Message: "Loader is Windows-only"}, fmt.Errorf("windowsReflectiveLoader called on non-Windows OS: %s", runtime.GOOS)
	}
	if len(dllBytes) == 0 {
		return models.ShellcodeResult{Message: "No DLL bytes provided"}, errors.New("empty DLL bytes")
	}
	if exportName == "" {
		return models.ShellcodeResult{Message: "Export name not specified"}, errors.New("export name required for DLL execution")
	}

	fmt.Printf("|📋 SHELLCODE DETAILS|\n-> Self-injecting DLL (%d bytes)\n-> Calling Function: '%s'\n",
		len(dllBytes), exportName)

	// TODO: Parse DOS Header
	// Hint: Use binary.Read with bytes.NewReader to read IMAGE_DOS_HEADER
	// Check dosHeader.Magic == IMAGE_DOS_SIGNATURE (0x5A4D = "MZ")
	reader := bytes.NewReader(dllBytes)
	var dosHeader IMAGE_DOS_HEADER
	// Read DOS header here

	// TODO: Seek to NT Headers using dosHeader.Lfanew and read PE signature
	// Hint: reader.Seek(int64(dosHeader.Lfanew), 0)
	// Then read uint32 PE signature and verify == IMAGE_NT_SIGNATURE

	// TODO: Read IMAGE_FILE_HEADER and IMAGE_OPTIONAL_HEADER64
	var fileHeader IMAGE_FILE_HEADER
	var optionalHeader IMAGE_OPTIONAL_HEADER64
	// Read headers here

	log.Println("|⚙️ SHELLCODE ACTION| [+] Parsed PE Headers successfully.")

	// TODO: Allocate memory for DLL using windows.VirtualAlloc
	// Hint: First try at preferred ImageBase, if that fails, try at address 0
	// allocBase, err := windows.VirtualAlloc(preferredBase, allocSize, windows.MEM_RESERVE|windows.MEM_COMMIT, windows.PAGE_EXECUTE_READWRITE)
	var allocBase uintptr
	allocSize := uintptr(optionalHeader.SizeOfImage)
	_ = allocSize // Remove when implemented
	// Allocate memory here

	// TODO: Copy PE headers to allocated memory
	// Hint: Use unsafe.Slice to get a byte slice view of allocated memory
	// memSlice := unsafe.Slice((*byte)(unsafe.Pointer(allocBase)), allocSize)
	// copy(memSlice[:headerSize], dllBytes[:headerSize])

	// TODO: Copy each section to its virtual address in allocated memory
	// Hint: Iterate through sections using fileHeader.NumberOfSections
	// For each section, copy SizeOfRawData bytes from PointerToRawData to VirtualAddress

	// TODO: Process base relocations if loaded at non-preferred address
	// Hint: If allocBase != optionalHeader.ImageBase, apply relocations
	// Use DataDirectory[IMAGE_DIRECTORY_ENTRY_BASERELOC] to find relocation table

	// TODO: Process Import Address Table (IAT)
	// Hint: Use DataDirectory[IMAGE_DIRECTORY_ENTRY_IMPORT] to find import directory
	// For each imported DLL: LoadLibrary, then resolve each function with GetProcAddress
	// Write resolved addresses to IAT

	// TODO: Call DLL entry point (DllMain) with DLL_PROCESS_ATTACH
	// Hint: entryPointAddr := allocBase + uintptr(optionalHeader.AddressOfEntryPoint)
	// syscall.SyscallN(entryPointAddr, allocBase, DLL_PROCESS_ATTACH, 0)

	// TODO: Find and call the exported function
	// Hint: Use DataDirectory[IMAGE_DIRECTORY_ENTRY_EXPORT] to find export directory
	// Search through export names to find exportName, get its address, and call it
	// syscall.SyscallN(targetFuncAddr)

	// Placeholder return - replace with actual result
	_ = allocBase
	_ = fileHeader
	_ = reader
	finalMsg := fmt.Sprintf("DLL loaded and export '%s' called successfully.", exportName)
	return models.ShellcodeResult{Message: finalMsg}, nil
}
