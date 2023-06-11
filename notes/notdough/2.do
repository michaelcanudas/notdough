import system

let main = void() {
    let x = -10
    let y = abs(x)

    system.print(y)
}

let abs = int(x: int) {
    if x < 0 {
        return -x
    }

    return x
}