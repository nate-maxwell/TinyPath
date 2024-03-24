package pathlib

import (
	"fmt"
)

/* A basic non-unit test for making sure basic functions work for now.
This will get replaced with an actual unit test in the future. */

func test() {
	MyPath := Path("T:/git/TinyPath/test_file.txt")
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
