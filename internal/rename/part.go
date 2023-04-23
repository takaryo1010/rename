package rename

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func RenamePart(root string, NewName string) {

	Names:=strings.Split(NewName, ",")
	var (
		FromName string = Names[0]
		ToName string = Names[1]
	)
	if len(Names)!=2{
		return 
	}
	
    filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }
        newPath := filepath.Join(root, info.Name())
		newFilePath := strings.Replace(path, FromName,ToName, -1)
        err = os.Rename(newPath, newFilePath)
        if err != nil {
            return nil
        }
        return nil
    })
}