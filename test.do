{
    il {
        func $abs
            store $x

            load $x
            push 0
            cmp >

            if
                load $x
                neg
                ret
            endif

            load $x
        endfunc
    }

    -69
    print abs()

    let abs: int() {
    }
}