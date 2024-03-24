# Tiny Path
A very simplistic and lightweight Go version of Python's Pathlib.

This is mostly a sandbox for me to learn Go with, but you can use anything you like.

The goal of this repo is to create a reliable path system from Go, that is OS-agnostic,
and does not rely on any third party libraries.

Currently only a handful of basic functions are implemented:

- Path() - Constructor
- path.Parent
- path.Parents
- path.Suffix
- path.Stem
- path.Name
- path.With_suffix
- path.With_stem
- path.With_name
- path.Drive

Over time this will grow.

### Example Code
Script:

```go
package main

import (
	"fmt"

	pathlib "github.com/nate-maxwell/TinyPath"
)

func main() {
	MyPath := pathlib.Path("T:/git/TinyPath/test_file.txt")
	fmt.Println("Original Path:", MyPath.AsPosix())
	fmt.Println()

	fmt.Println("Posixed:", MyPath.AsPosix())
	fmt.Println()

	fmt.Println("Parent:", MyPath.Parent().AsPosix())
	fmt.Println()

	fmt.Println("Parents: [")
	for _, path := range MyPath.Parents() {
		fmt.Println(path.AsPosix())
	}
	fmt.Println("]")
	fmt.Println()

	fmt.Println("Stem:", MyPath.Stem())
	fmt.Println("Suffix:", MyPath.Suffix())
	fmt.Println("Name:", MyPath.Name())
	fmt.Println()

	newStem := MyPath.WithStem("New")
	fmt.Println("With Stem:", newStem.AsPosix())

	newSuffix := MyPath.WithSuffix(".json")
	fmt.Println("With Suffix:", newSuffix.AsPosix())

	newName := MyPath.WithName("Newer.fbx")
	fmt.Println("With Name:", newName.AsPosix())
	fmt.Println()

	drive := MyPath.Drive()
	fmt.Println("Drive:", drive)
	fmt.Println()
}
```

Output:
```
Original Path: C:/git/GoPathlib/test_file.txt

Posixed: C:/git/GoPathlib/test_file.txt

Parent: C:/git/GoPathlib

Parents: [
C:/git/GoPathlib
C:/git
C:
]

Stem: test_file
Suffix: .txt
Name: test_file.txt

With Stem: C:/git/GoPathlib/New.txt
With Suffix: C:/git/GoPathlib/test_file.json
With Name: C:/git/GoPathlib/Newer.fbx

Drive: C:
```
