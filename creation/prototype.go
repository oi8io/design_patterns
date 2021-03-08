package creation

import "fmt"

type DirNode interface {
	print(string)
	clone() DirNode
}

type FileNode struct {
	Name string
}

func (f *FileNode) print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *FileNode) clone() DirNode {
	return &FileNode{Name: f.Name + "_clone"}
}

type Folder struct {
	Children []DirNode
	Name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i2 := range f.Children {
		i2.print(indentation+indentation)
	}
}

func (f *Folder) clone() DirNode {
	var tmp []DirNode
	for _, child := range f.Children {
		copy := child.clone()
		tmp = append(tmp, copy)
	}
	return &Folder{
		Children: tmp,
		Name:     f.Name + "_clone",
	}
}
