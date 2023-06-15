func $PRINT69
    push 69
    print
endfunc

# load can now reference globals (func/rsrc) rn it pushes the location of that instruction (so 0 here)
load $PRINT69

# call without a function now pops an value from the execution stack
# and jumps to the function at that instruction. if it's not a func,
# it panics.

call