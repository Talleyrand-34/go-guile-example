
package main

/*
#cgo pkg-config: guile-3.0
#include <libguile.h>
#include <stdlib.h>

// Correct bootstrap function with required parameters
void guile_main(void *closure, int argc, char **argv) {
    // Example Scheme code to print a message
    scm_c_eval_string("(display \"Hello from Guile Scheme!\\n\")");
}

// Wrapper function that calls scm_boot_guile passing
// command line arguments and the bootstrap callback
int call_scm_boot_guile(int argc, char **argv) {
    scm_boot_guile(argc, argv, guile_main, NULL);
    return 0; // scm_boot_guile never returns, but just to satisfy compiler
}
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
