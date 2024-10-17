# Takeaways

flags module is used for flags parsing. It is necessary to call parse function as the first line in order to read the flags.

```go
func main() {
    flags.Parse()
}
```

You can see what flags are available by running the binary with -h or -help flag

```bash
./program -help
./program -h
```
