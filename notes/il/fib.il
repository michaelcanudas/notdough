rsrc $DESC text "enter number of digits to print:"

func $FIB
    store $COUNT
    push 0
    store $A
    push 1
    store $B

    load $COUNT
    push 0
    cmp ==
    if
        ret
    endif

    load $COUNT
    dec
    store $COUNT

    load $A
    print

    load $COUNT
    push 0
    cmp ==
    if
        ret
    endif

    load $COUNT
    dec
    store $COUNT

    load $B
    print

    loop
        load $A
        load $B
        add
        load $B
        store $A
        dup
        store $B
        print
        
        load $COUNT
        dec
        dup
        store $COUNT
        push 0
        cmp <
        if
            cont
        endif
    endloop
endfunc

load $DESC
printstr
read
call $FIB