{
    let add: int() = {
        il {
            store $x
        }
        return x + 1
    }

    print add(5)
}