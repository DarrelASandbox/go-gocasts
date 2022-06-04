## About The Project

- Go: The Complete Developer's Guide (Golang)
- Master the fundamentals and advanced features of the Go Programming Language (Golang)
- Tutorial for GoCasts
- [Original Repo: GoCasts](https://github.com/StephenGrider/GoCasts)
- [Stephen Grider](https://github.com/StephenGrider)

&nbsp;

---

&nbsp;

## Basics

```sh
go
go run main.go
go help packages
```

- <b>Type of packages:</b>
  - <b>Executable:</b> Generates a file that we can run (<i>"main" is special</i>)
  - <b>Reusable:</b> Code used as 'helpers'. Good place to put reusable logic
- <code>package main</code> Defines a package that can be compiled and then _executed_. <b>Must have a func called 'main'</b>
- <code>package calculator</code> Defines a package that can be used as a dependency (helper code)
- <code>package uploader:</code> Defines a package that can be used as a dependency (helper code)
- Variables can be initialized outside of a function, but cannot be assigned a value.

&nbsp;

---

&nbsp;

```sh
# https://stackoverflow.com/questions/58018729/go-linter-in-vs-code-not-working-for-packages-across-multiple-files
# Initialize new module in current directory
go mod init gocasts
```

```go
var card string = "Ace of Spades"
```

- [Linux File/Directory Permissions cheat sheet](https://www.thegeekdiary.com/linux-file-directory-permissions-cheat-sheet/)
- [What is the difference between concurrency, parallelism and asynchronous methods?](https://stackoverflow.com/questions/4844637/what-is-the-difference-between-concurrency-parallelism-and-asynchronous-methods)
- [Threading Tutorial #1 - Concurrency, Threading and Parallelism Explained](https://www.youtube.com/watch?v=olYdb0DdGtM)

&nbsp;

---

&nbsp;
