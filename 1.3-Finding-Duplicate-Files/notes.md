# Notes

- Map provides constant time ops for get/set
- `bufio` makes io efficient and convienient. `Scanner` reads input and breaks into line or words which is often the easiest way to process input

`input := bufio.NewScanner(os.Stdin)`

Then you can  `input.Scan()` to read the next line, returning `true` if there is input. `input.Text()` returns the text.

Both `os.Stdin` and the return type of `os.Open(filepath)` are `os.File` and as such can be used by `bufio.Scanner`.

*However*, if you look at the [docs for `bufio.Scanner`](https://golang.org/pkg/bufio/#NewScanner) then you can see that it takes an io.Reader which is much easier to create for testing. 

A `map` is a reference to the underlying structure so you can safely pass it to functions without it being copied.
