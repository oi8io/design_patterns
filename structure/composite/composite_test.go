package structure

import "testing"

func TestComponent_Add(t *testing.T) {

	file1 := &Leaf{name: "File1"}
	file2 := &Leaf{name: "File2"}
	file3 := &Leaf{name: "File3"}

	folder1 := &Component{
		name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &Component{
		name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.process("---")
}