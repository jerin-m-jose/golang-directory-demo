package inmemoryfs

import (
	"errors"
	"filesystemdemo/internal/filesystem"
	"fmt"
	"sort"
	"strings"
)

// New Create a new in memory filesystem
func New() filesystem.FileSystem {
	return &InMemoryFS{
		root: &Directory{name: ".", children: make(map[string]*Directory)},
	}
}

// InMemoryFS implements the FileSystem interface
type InMemoryFS struct {
	root *Directory
}

func (ifs *InMemoryFS) GetChildren() interface{} {
	return ifs.root.children
}

func (ifs *InMemoryFS) Create(name string) error {
	parts := strings.Split(name, "/")
	dir := ifs.root

	//create nested child directories as needed
	for _, part := range parts {
		if _, exists := dir.children[part]; !exists {
			dir.children[part] = &Directory{name: part, children: make(map[string]*Directory)}
		}
		dir = dir.children[part]
	}
	return nil
}

func (ifs *InMemoryFS) List(sort filesystem.Sort) {
	switch sort {
	case filesystem.AlphabeticalSort:
		ifs.listChildrenAlphabeticalSort(ifs.root, "")
	default:
		ifs.listChildrenAlphabeticalSort(ifs.root, "")
	}

}

func (ifs *InMemoryFS) listChildrenAlphabeticalSort(dir *Directory, indent string) {
	if len(dir.children) > 0 {
		// Step 1: Collect the children Directory names
		childrenNames := make([]string, 0, len(dir.children))
		for name := range dir.children {
			childrenNames = append(childrenNames, name)
		}

		// Step 2: Sort the Directory names alphabetically
		sort.Strings(childrenNames)

		for _, childName := range childrenNames {
			child := dir.children[childName]
			fmt.Println(indent + child.name)
			ifs.listChildrenAlphabeticalSort(child, indent+"  ")
		}
	}
}

// Delete delete the specified folder
func (ifs *InMemoryFS) Delete(name string) error {
	parts := strings.Split(name, "/")

	//delete nested child directories as needed
	dir := ifs.root
	parent := ifs.root
	for _, part := range parts {
		if _, exists := dir.children[part]; exists {
			parent = dir
			dir = dir.children[part]
		} else {
			return errors.New(part + " does not exist")
		}

	}
	delete(parent.children, dir.name)

	return nil
}

func (ifs *InMemoryFS) Move(src string, dest string) error {
	srcDirs := strings.Split(src, "/")
	srcParent := ifs.root

	//find src Directory to move
	var dirToMove *Directory
	var srcDir string
	var index int
	for index, srcDir = range srcDirs {
		if _, exists := srcParent.children[srcDir]; exists {
			if index == len(srcDirs)-1 {
				dirToMove = srcParent.children[srcDir]
			} else {
				srcParent = srcParent.children[srcDir]
			}
		} else {
			return errors.New(srcDir + " does not exist")
		}
	}

	//find dest Directory
	destDirs := strings.Split(dest, "/")
	destParent := ifs.root
	var destDir string
	for _, destDir = range destDirs {
		if _, exists := destParent.children[destDir]; exists {
			destParent = destParent.children[destDir]
		} else {
			return errors.New(destDir + " does not exist")
		}
	}

	if dirToMove != nil {
		//move to dest Directory
		destParent.children[srcDir] = dirToMove

		//delete from original parent
		delete(srcParent.children, srcDir)
	} else {
		return errors.New(srcDir + " does not exist")
	}

	return nil
}

// Directory in-memory file system
type Directory struct {
	name     string
	children map[string]*Directory
}
