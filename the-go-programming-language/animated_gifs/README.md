# Takeaways

You can still write to a file by having the os.Stdout as a io.Writer (you don't need to pass in the file).

```go
func main() {
	lissajous(os.Stdout)
}
```

You just need to call the compiled binary with:

```bash
./animated_gif > out.gif
```

This is the same as doing something like this:

```go
func main() {
	f, err := os.Create("out.gif")
	if err != nil {
		fmt.Println("cannot create out.gif", err)
		return
	}
	defer f.Close()
	lissajous(os.Stdout)
}
```

And after that executing the compiled binary with:
```bash
./animated_gif
```
