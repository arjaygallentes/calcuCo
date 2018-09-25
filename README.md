## CalcuCo

### Install

### Usage

You can use the application in two ways: Shell mode or API

### Shell Mode

The shell mode uses the `terminal` package provided by [`golang.org/x/crypto/ssh/terminal`](https://godoc.org/golang.org/x/crypto/ssh/terminal), which means it supports many of the shell features you know and love (like history, pasting, and the `exit` command).

```shell
> 1+1
2
> 3(5/(3-4))
-15
> 3pi^2
29.608813203268074
> @+1
30.608813203268074
> @@@*2
-30
> ln(-1)
NaN
```

### Supported functions, operators, and constants

CalcuCo supports the following operations: `+`, `-`, `*`, `/`
