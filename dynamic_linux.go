//go:build dynamic && linux

package main

import _ "unsafe"

// Force an ELF PT_INTERP entry so the scratch release tests the dynamic glibc
// metadata path without adding a runtime dependency to the static ARM build.
//go:cgo_import_dynamic libc_puts puts "libc.so.6"
