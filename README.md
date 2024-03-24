# Tiny Path
A very simplistic and lightweight Go version of Python's Pathlib.

This is mostly a sandbox for me to learn Go with, but you can use anything you like.

The goal of this repo is to create a reliable path system from Go, that is OS-agnostic,
and does not rely on any third party libraries.

Currently only a handful of basic functions are implemented:

- Path() - Constructor
- path.parent
- path.parents
- path.suffix
- path.stem
- path.name
- path.with_suffix
- path.with_stem
- path.with_name

Over time this will grow.

### Example Code

```go
package main

import (
	"fmt"

	pathlib "github.com/nate-maxwell/TinyPath"
)

func main() {
	MyPath := pathlib.Path("T:/git/GoPathlib/test_file.txt")
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
}
```
