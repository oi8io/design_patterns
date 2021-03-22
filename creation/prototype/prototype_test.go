package creation

import (
	"testing"
)

func TestProtoType(t *testing.T) {
	file1 := &FileNode{Name: "file1"}
	file2 := &FileNode{Name: "file2"}
	file3 := &FileNode{Name: "file3"}
	folder1 := &Folder{
		Children: []DirNode{file1},
		Name:     "folder1",
	}
	folder2 := &Folder{
		Children: []DirNode{folder1,file2,file3},
		Name:     "folder2",
	}
	folder2.print(" ")
	folder3 := folder2.clone()
	folder3.print(" ")
}






