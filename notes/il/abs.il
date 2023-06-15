func $abs
    store $x

    load $x
    push 0
    cmp >

    if
        load $x
        neg
    else
        load $x
    endif

    ret
endfunc

read 
call $abs
print