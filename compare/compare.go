// Package compare provides funcs to compare directories
package compare

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/udhos/equalfile"
	"path/filepath"
)

// CompareFilesInDirs :
// compare list of files in two directories
func CompareFilesInDirs(dir1 string, dir2 string, filesToCompare []string) {
	for _, file := range filesToCompare {
		cmp := equalfile.New(nil, equalfile.Options{})
		isSame, err := cmp.CompareFile(filepath.Join(dir1, file), filepath.Join(dir2, file))

		if err != nil {
			color.Red("%v\n", err)
		}
		if isSame {
			fmt.Printf("%s - %s => %s\n", filepath.Join(dir1, file), filepath.Join(dir2, file), color.GreenString("Same"))
		} else {
			fmt.Printf("%s - %s => %s\n", filepath.Join(dir1, file), filepath.Join(dir2, file), color.RedString("Diff"))
		}
	}
}
