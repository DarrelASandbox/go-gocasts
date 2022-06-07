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
- [Package fmt verbs](https://pkg.go.dev/fmt)

&nbsp;

---

&nbsp;

## Structs & Maps

|  Type  | Zero Value |
| :----: | :--------: |
| string |     ""     |
|  int   |     0      |
| float  |     0      |
|  bool  |   false    |

- GO is pass by value
- <b>RAM</b>

| Address |            Value            |     |
| :-----: | :-------------------------: | :-: |
|  0000   |                             |     |
|  0001   | person{firstName: "Jim"...} | jim |
|  0002   | person{firstName: "Jim"...} |  p  |
|  0003   |                             |     |

```go
// &variable : Give me the memory address of the value this variable is pointing at
// *pointer : Give me the value this memory address is pointing at

// *person : This is a type description - it means we're working with a pointer to a person
// *pointerToPerson : This is an operator - it means we want to manipulate the value the pointer is referencing

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Turn address into value with *address
// Turn value into address with &value
```

&nbsp;

---

&nbsp;

> <b>Dipesh: </b> Gotchas With Pointers
>
> This is not actually gotcha of pointer. To understand this we need to understand memory type and immutability. As you said in previous videos, go make copy of data to make it change, it is because of go's characteristic of immutability. Now, any variable stored on memory has two options, it could be added in stack(when it has fix size) or in heap memory(when it need to allocate more size then it will actually consume). Here slice is not fixed size element, as you said in previous videos, it can be shrink or expand. Pointer can change value by reference when it is in stack, but if any variable is stored in heap memory, it can be change directly as it is stored on memory which has greater amount of memory already allocated.

|          Arrays          |                   Slices                   |
| :----------------------: | :----------------------------------------: |
| Primitive data structure |            Can grow and shrink             |
|     Can't be resized     | Used 99% of the time for lists of elements |
|   Rarely used directly   |                                            |

| Value Types | Reference Types |
| :---------: | :-------------: |
|     int     |     slices      |
|    float    |      maps       |
|   string    |    channels     |
|    bool     |    pointers     |
|   structs   |    functions    |

- When we create a slice, Go will automatically create which two data structures?
  - An array and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array

&nbsp;

---

&nbsp;

|                         Map                         |                            Struct                             |
| :-------------------------------------------------: | :-----------------------------------------------------------: |
|       All keys & values must be the same type       |                Values can be of different type                |
|     Keys are indexed - we can iterate over them     |                  Keys don't support indexing                  |
|                   Reference Type                    |                          Value Type                           |
| Use to represent a collection of related properties | Use to represent a "thing" with a lot of different properties |
|   Don't need to know all the keys at compile time   |   You need to know all the different fields at compile time   |

- Map is good for closely related items liek color and hex codes

&nbsp;

---

&nbsp;

## Interfaces

- We know that...
  - Every value has a type
  - Every function has to specify the type of its arguments
- So does that mean...
  - Every function we ever write has to be rewritten to accommodate different types even if the logic in it is identical?

```go
type bot interface {
  getGreeting(string, int)(string, error)
}

// Function name
// List of argument types
// List of return types
```

| Concrete Type | Interface Type |
| :-----------: | :------------: |
|      map      |      bot       |
|    struct     |                |
|      int      |                |
|    string     |                |
|  englishBot   |                |

|                                                                 |                                                                                                                      |
| :-------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------------------: |
|                Interfaces are not generic types                 |                            Other languages have 'generic' types - go (famously) does not.                            |
|                    Interfaces are 'implicit'                    |                     We don't manually have to say that our custom type satisfies some interface.                     |
|        Interfaces are a contract to help us manage types        | GARBAGE IN -> GARBAGE OUT. If our custom type's implementation of a function is broken then interfaces wont help us! |
| Interfaces are tough. Step #1 is understanding how to read them | Understand how to read interfaces in the standard lib. Writing your own interfaces is tough and requires experience  |

![response-struct](./diagrams/diagrams-014-response-struct.png)

- [Reader Interface](https://pkg.go.dev/io#Reader)

&nbsp;

---

&nbsp;
