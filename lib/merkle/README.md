# merkle

This is a port of OpenZeppelin's JS [merkle-tree](https://github.com/OpenZeppelin/merkle-tree/tree/master) library to Go.
It excludes the `StandardTree` wrapper types since our use-case doesn't require it.
It also moves all non-omni-required logic to the test package to decrease the surface area of the library.
