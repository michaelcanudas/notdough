# WHAT
let main: void() = {}
let main = void() {}

# the above code has implications (which i like):
# - saying {} declares a anonymous function, leaving it's signature to be inferenced
# - saying void() {} declares a typed anonymous function
# - all functions are anonymous (woah)

# i like var=mutable and let=immutable
# so for all intents and purposes declaring an anonymous function with let in the
# global scope would be like straight up declaring a normal func in another language