# Writing An Interpreter in Go

This is the interpreter I implemented when reading [Writing An Interpreter in Go](https://interpreterbook.com/),
mainly for the purpose of rewriting [Univer](https://github.com/dream-num/univer)'s formula engine.

Made some changes to the original code:

1. Use go.mod to manage the module and packages.
2. Add position info to tokens.
