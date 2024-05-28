## Go "Hello, World" Tutorial with TDD

## Basic "Hello, World"

- Create `hello.go`
- Write code:

Go

```
package main // Package declaration

import "fmt" // Import the fmt package for printing

func main() {  // Main function where the program starts
    fmt.Println("Hello, world") // Print "Hello, world"
}
```

- Run with `go run hello.go`

## Test-Driven Development (TDD) Cycle:

- **Write a failing test:** Create `test.go` file and write a test function that asserts the expected output.
- **Make the compiler pass:** Modify the code in `hello.go` to compile successfully.
- **Run the test and observe failure:** The test should fail initially, indicating the code doesn't meet the requirement.
- **Write code to pass the test:** Implement the logic in `hello.go` to make the test pass.
- **Refactor:** Improve the code's structure and readability while keeping the tests passing.

## Key Go Concepts:

- **Functions, arguments, and return types:**
```go
func Hello(name string) string { // Function with argument and return type
    return "Hello, " + name 
}

func getPrefix(language string) (prefix string) { // prefix gets initalized here and prefix is what will get returned

	switch language {
	case zulu:
		prefix = "Sawubona"
	case spanish:
		prefix = "Hola, "
	case french:
		prefix = "Bonjour"
	default:
		prefix = "Hello, "
	}

	return
}
```

- **`if`, `const`, and `switch` statements:**
```go
if name == "" { // If statement
    name = "World"
}

const englishHelloPrefix = "Hello, " // Constant declaration

switch language { // Switch statement
    case "Spanish":
        return "Hola, " + name
    default:
        return englishHelloPrefix + name
}
```

- **Variables and constants:**

```go
var message string = "Hello, world" // Variable declaration and initialization
const pi = 3.14159 // Constant declaration
```

- **Package imports:**
```go
import (
    "fmt"
    "testing"
)
```

- **Test functions:**
```go
func TestHello(t *testing.T) { // Test function
    // Test assertions
}
```

## Benefits of TDD:

- Ensures code correctness and functionality
- Improves code design and structure
- Facilitates refactoring and maintenance
- Enables incremental development with **confidence**