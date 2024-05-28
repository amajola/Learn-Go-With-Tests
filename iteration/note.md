# Iteration in Go

## Key Concepts

- **For loop:** The primary way to do repetitive tasks in Go. There are no `while`, `do`, or `until` keywords.
- **Benchmarking:** Go has built-in benchmarking capabilities, allowing you to measure the performance of your code.
- **Strings package:** Offers various functions for working with strings.

## Code Examples

Code snippet

```go
package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character // `+=` assignment operator.
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```


## Benchmarking

The `testing.B` gives you access to the cryptically named `b.N`.

When the benchmark code is executed, it runs `b.N` times and measures how long it takes.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

To run the benchmarks do `go test -bench=.` (or if you're in Windows Powershell `go test -bench="."`)


## Practice Exercises

- Modify the test to allow specifying the number of repetitions.
- Write an `ExampleRepeat` function to document the usage of `Repeat`.
- Explore the `strings` package and experiment with different string functions using tests.