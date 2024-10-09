package prototype

import "fmt"

type Folder struct {
	Children []Inode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	newFolder := &Folder{Name: f.Name + "_clone"}
	var temChildren []Inode
	for _, i := range f.Children {
		copy := i.Clone()
		temChildren = append(temChildren, copy)
	}
	newFolder.Children = temChildren
	return newFolder

}
