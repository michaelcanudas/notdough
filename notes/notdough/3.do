import system

let main = void() {
    loop(10, void() {
        system.print("hello!")
    })
}

let loop = void(x: int, fn: void()) {
    if x <= 0 {
        return
    }

    fn()
    loop(x - 1, fn)
}