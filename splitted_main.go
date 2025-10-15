
package main

/*
#cgo pkg-config: guile-3.0
#include "csrc/guile_embed.h"
#include <stdlib.h>
*/
import "C"

import (
    "fmt"
    "os"
    "unsafe"
)

func main() {
    fmt.Println("Hello from Go!")

    argc := C.int(len(os.Args))
    argv := make([]*C.char, len(os.Args)+1)

    for i, s := range os.Args {
        argv[i] = C.CString(s)
        defer C.free(unsafe.Pointer(argv[i]))
    }
    argv[len(os.Args)] = nil

    // Call Guile initialization with wrapper
    _ = C.call_scm_boot_guile(argc, &argv[0])
}
