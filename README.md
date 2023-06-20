# Effortless Golang

<a href=""><img src="./effortless-golang-cover.png" alt="Effortless Golang" height="256px" align="right"></a>

I considered writing a book about Go, but decided not to. 

Instead I've compiled a list of curated resources to make learning the language as effortless as possible. Hince the name "Effortless Go".

You can **Follow** me on Twitter [@joshduffney](https://twitter.com/joshduffney).

Learn the Language Fundamentals
- [Introducing Go](https://www.oreilly.com/library/view/introducing-go/9781491941997/) by Caleb Doxsey
- [Go in Action](https://www.manning.com/books/go-in-action-second-edition) by William Kennedy, Brian Ketelsen, and Erik St.Martin
- [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/) by Jon Bodner

Web Development
- [Building Go Web Services and Applications](https://app.pluralsight.com/library/courses/go-building-web-services-applications) by me
- [Let's Go](https://lets-go.alexedwards.net/) by Alex Edwards
- [Let's Go Further](https://lets-go-further.alexedwards.net/) by Alex Edwards 

Building CLIs
- [Master Go CLI and TUI](https://leanpub.com/go-cli-tui) By Ravikanth Chaganti
- [Building Modern CLI Applications in Go](https://www.packtpub.com/product/building-modern-cli-applications-in-go) by Marian Montagnino

---

## Table of Contents

> This is the outline from the book proposal I never sent. But if I had to start over learning Go, this is how I'd go about it.

1. Before you begin
    1. Who should read this book.
    2. Go as your first programming language.
    3. [The Go Playground](https://play.golang.com/)
2. Built-in types
    1. [Data structures](https://www.geeksforgeeks.org/data-structures/)
    2. [Staticaly typed vs dynamic](https://stackoverflow.com/questions/1517582/what-is-the-difference-between-statically-typed-and-dynamically-typed-languages)
    3. [Primitive types](https://medium.com/golang-jedi-knight/primitive-data-types-in-golang-35a291df3bbe)
        - Numbers
        - Booleans
        - Bytes and Runes
    4. Composite types
        - Arrays
        - Slices
        - Strings
        - Maps
    5. Using fmt to determine the type.
    6. The Zero Value
3. Variables, Constants, and Declarations
    1. Naming variables
    2. Defining muliple variables
    3. Using the short variable delcaration :=
    4. Variable scopes
        - blocks
    5. Constants
        - const vs var
    6. Type and untyped declarations
        - Explicit type conversion
4. Control Structures
    1. What is a control structure?
    2. The if Statement.
        - Avoid else at all costs.
    3. The switch statements.
    4. The for Statement
        - Range
5. Functions
    1. Declaring and calling functions
        - Named and optional parameters
    2. Variadic input parameters
    3. Return values.
    4. Ignoring returned values
    5. Anonymous functions
    6. Closures
6. Pointers
    1. What are pointers?
    2. Understanding memory addresses and values
    3. Why are pointers important?
    4. Declaring, referencing, and dereferencing pointers.
    5. Avoiding nil pointers and null pointer dereferences. 
7. Go’s Type System
    1. Structs
        - User defined types
        - Type embedding
        - Anonymous structs
        - Struct literals
    2. Methods
        - Receivers
        - Method sets.
    3. Interfaces
        - Polymorphism
8. Leaving the playground
    1. Installing Go
    2. Go developer survey data results for Go dev environments.
    3. VS code setup
    4. Neovim setup
    5. Using the Go tools
        - Go run
        - Go build
9. Packages, modules, and repositories
    1. Creating a package
        - Go mod
    2. Flat structure
    3. Within a directory
    4. Withing a pkg folder
    5. Imports and Exports
        - Capitals for exports
        - Local imports
    6. 3rd party imports
    7. Is that’s a package, what’s a module?
    8. Library \ dependency management
    9. Comment-based documentation
10. The Standard Library
    1. Navigating pkg.go.dev
    2. Core packages
        - Reading data from files (io/ioutil,)
        - Create an http server.
        - Sorting with algorithms
11. Create a RESTful web service
    1. Setting up the project structure
    2. Understanding the net/http package
    3. Creating a healthcheck endpoint
    4. Making a user-defined application type with methods
    5. Adding routes and handlers
12. Build a command line app
    1. Adding a cmd app to the project repo
    2. Understanding the flags package
    3. Designing use cases
    4. The anatomy of a commandline app
    5. Implementing commands
    6. Adding flags
    7. Interacting with APIs
13. Writing Tests in Go
    1. Identifying test cases
    2. The test package.
    3. What are the different types of tests?
    4. Writing unit tests to validate code.
    5. Testing your API
    6. Testing your CLI app
