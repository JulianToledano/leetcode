package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */

func main() {
	obj := Constructor()
	obj.Mkdir("/goowmfn")
	fmt.Println(obj.Ls("/goowmfn"))
	fmt.Println(obj.Ls("/"))
	obj.Mkdir("/z")
	fmt.Println(obj.Ls("/"))
	fmt.Println(obj.Ls("/"))
	obj.AddContentToFile("/goowmfn/c", "shetopcy")
	fmt.Println(obj.Ls("/z"))
	fmt.Println(obj.Ls("/goowmfn/c"))
	fmt.Println(obj.Ls("/goowmfn"))
}

type Node struct {
	isDir     bool
	Name      string
	Father    *Node
	Childrens []*Node
	Contents  string
}

type FileSystem struct {
	Root            *Node
	CurrentPosition *Node
}

func Constructor() FileSystem {
	root := &Node{
		true,
		"/",
		nil,
		[]*Node{},
		"",
	}
	return FileSystem{
		root,
		root,
	}
}
func ParsePath(path string) []string {
	p := strings.Split(path, "/")
	return p[1:]
}

func (this *FileSystem) Ls(path string) []string {
	if path != "/" {
		p := ParsePath(path)
		this.goTo(p)
	}
	r := []string{}
	if this.CurrentPosition.isDir {
		for _, v := range this.CurrentPosition.Childrens {
			r = append(r, v.Name)
		}
		this.CurrentPosition = this.Root
		sort.Strings(r)
		return r
	}
	fileName := this.CurrentPosition.Name
	this.CurrentPosition = this.Root
	return []string{fileName}
}

func (this *FileSystem) Mkdir(path string) {
	p := ParsePath(path)
	this.mkdir(p)
	this.CurrentPosition = this.Root
}
func (this *FileSystem) mkdir(p []string) {
	for _, v := range p {
		f := find(v, this.CurrentPosition.Childrens)
		if f == nil {
			temp := &Node{
				true,
				v,
				this.CurrentPosition,
				[]*Node{},
				"",
			}
			this.CurrentPosition.Childrens = append(this.CurrentPosition.Childrens, temp)
			this.CurrentPosition = temp
		} else {
			this.CurrentPosition = f
		}
	}
}
func find(name string, files []*Node) *Node {
	for _, f := range files {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (this *FileSystem) goTo(path []string) {

	this.CurrentPosition = this.Root
	for _, v := range path {
		for _, n := range this.CurrentPosition.Childrens {
			if n.Name == v {
				this.CurrentPosition = n
				break
			}
		}
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	p := ParsePath(filePath)
	fileName := p[len(p)-1]
	p = p[:len(p)-1]
	this.mkdir(p)
	f := this.findFile(fileName)
	this.goTo(p)
	if f == nil {
		temp := &Node{
			false,
			fileName,
			this.CurrentPosition,
			[]*Node{},
			content,
		}
		this.CurrentPosition.Childrens = append(this.CurrentPosition.Childrens, temp)
	} else {
		f.Contents += content
	}
	this.CurrentPosition = this.Root
}

func (this *FileSystem) findFile(name string) *Node {
	for _, f := range this.CurrentPosition.Childrens {
		if name == f.Name && !f.isDir {
			return f
		}
	}
	return nil
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	p := ParsePath(filePath)
	fileName := p[len(p)-1]
	p = p[:len(p)-1]
	this.goTo(p)
	f := this.findFile(fileName)
	r := ""
	if f != nil {
		r = f.Contents
	}
	this.CurrentPosition = this.Root
	return r
}
