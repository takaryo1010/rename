package rename

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

func RenameAll(root string, NewName string) {
    var i int = 1
    filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }
        newPath := filepath.Join(root, info.Name())
        newFilePath := filepath.Join(root, NewName+strconv.Itoa(i)+filepath.Ext(path))
        err = os.Rename(newPath, newFilePath)
        if err != nil {
            return nil
        }
		i++
        return nil
    })
}

