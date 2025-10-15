
// guile_embed.c
#include <libguile.h>
#include <stdlib.h>

// Correct bootstrap function with required parameters
void guile_main(void *closure, int argc, char **argv) {
    // Example Scheme code that prints a message
    scm_c_eval_string("(display \"Hello from Guile Scheme!\\n\")");
}

// Wrapper function that calls scm_boot_guile passing
// command line arguments and the bootstrap callback
int call_scm_boot_guile(int argc, char **argv) {
    scm_boot_guile(argc, argv, guile_main, NULL);
    return 0; // scm_boot_guile never returns, but this satisfies compiler
}
