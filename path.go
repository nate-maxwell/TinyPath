package pathlib

import (
	"runtime"
	"strings"
)

var defaultSlash string

func init() {
	switch runtime.GOOS {
	case "windows":
		defaultSlash = "\\"
	default:
		defaultSlash = "/"
	}
}

/* ----------------------------------------------------------------------------------------------------------
Path object core & constructor
---------------------------------------------------------------------------------------------------------- */

type PathStruct struct {
	path string
	sep  string
}

func Path(args ...any) *PathStruct {
	var params []string
	sep := defaultSlash

	for _, obj := range args {
		switch v := obj.(type) {
		case string:
			// Convert '\' to '/' and split to get all tokens adn then append to params
			forwarded := strings.Replace(v, "\\", "/", -1)
			tokens := strings.Split(forwarded, "/")
			params = append(params, tokens...)
		case PathStruct:
			forwarded := v.AsPosix()
			tokens := strings.Split(forwarded, "/")
			params = append(params, tokens...)
		}

	}

	path := strings.Join(params, sep)
	return &PathStruct{path, sep}
}

/* ----------------------------------------------------------------------------------------------------------
Path functions
---------------------------------------------------------------------------------------------------------- */

// ----------Path Conversion----------

// Converts PathStruct path to a posix path.
func (p PathStruct) AsPosix() string {
	return strings.Replace(p.path, "\\", "/", -1)
}

// ----------Parents----------

// Returns the parent directory to the PathStruct Path.
// Else returns the PathStruct path if there is no parent.
func (p PathStruct) Parent() *PathStruct {
	posixed := p.AsPosix()

	lastIndex := strings.LastIndex(posixed, "/")
	if lastIndex == -1 {
		return &p
	}

	posParent := posixed[:lastIndex]
	vanilla := strings.Replace(posParent, "/", "\\", -1)
	newVal := vanilla

	return Path(newVal)
}

// Returns an array of PathStructs for each parent directory of the given PathStruct.
func (p PathStruct) Parents() []*PathStruct {
	parents := []*PathStruct{}
	parentCount := len(strings.Split(p.AsPosix(), "/"))
	if parentCount == -1 {
		parents = []*PathStruct{}
	}

	curPath := p

	for i := 0; i < parentCount-1; i++ {
		if curPath.Parent().path == "" {
			continue
		}
		parent := *curPath.Parent()
		parents = append(parents, &parent)
		curPath = parent
	}

	return parents
}

// ----------Suffix, Stem, Name----------

// Returns the file suffix for the given path, e.g. "C:/dir/file.txt" will return ".txt".
// Else returns an empty string if no suffix could be found.
func (p PathStruct) Suffix() string {
	posixed := p.AsPosix()
	lastIndex := strings.LastIndex(posixed, ".")
	if lastIndex == -1 {
		return ""
	}

	return posixed[lastIndex:]
}

// Returns the file name for the given path, e.g. "C:/dir/file.txt" will return "file.txt".
// Else returns an empty string if no name could be found.
func (p PathStruct) Name() string {
	posixed := p.AsPosix()
	lastIndex := strings.LastIndex(posixed, "/")
	if lastIndex == -1 {
		return ""
	}

	return posixed[lastIndex+1:]
}

// Returns the fiel stem fo rthe given path, e.g. "C:/dir/file.txt" will return "file".
// Else returns an empty string if no stem could be found.
func (p PathStruct) Stem() string {
	name := p.Name()
	if name == "" {
		return ""
	}

	return strings.Split(name, ".")[0]
}

// Returns a PathStruct using the given suffix.
func (p PathStruct) WithSuffix(suffix string) *PathStruct {
	curSuffix := p.Suffix()
	newPath := ""
	if curSuffix == "" {
		newPath = strings.Join([]string{p.AsPosix(), suffix}, "")
	} else {
		newPath = strings.Replace(p.AsPosix(), curSuffix, suffix, -1)
	}

	return Path(newPath)
}

// Returns a PathStruct using the given name.
func (p PathStruct) WithName(name string) *PathStruct {
	curName := p.Name()
	newPath := ""
	if curName == "" {
		newPath = strings.Join([]string{p.AsPosix(), name}, "")
	} else {
		newPath = strings.Replace(p.AsPosix(), curName, name, -1)
	}

	return Path(newPath)
}

// Returns a PathStruct using the given stem.
func (p PathStruct) WithStem(stem string) *PathStruct {
	curStem := p.Stem()
	curName := p.Name()
	newName := ""
	if curStem == "" || curName == "" {
		newName = stem
	} else {
		newName = strings.Replace(curName, curStem, stem, -1)
	}

	return p.WithName(newName)
}

// ----------Drive, Root----------

// Returns the drive letter for the given PathStruct, e.g. "C:"
func (p PathStruct) Drive() string {
	posixed := p.AsPosix()
	temp := strings.Split(posixed, ":")[0]
	if temp == posixed {
		return ""
	}
	drive := strings.Join([]string{temp, ":"}, "")
	return drive
}
