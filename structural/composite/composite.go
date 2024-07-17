package composite

import "fmt"

// INode 是节点接口，定义了节点操作的基本方法
type INode interface {
	Add(child INode) // 添加子节点
	Tree(space int)  // 显示节点树状结构
}

// Node 是抽象节点，包含节点公共属性
type Node struct {
	name string
}

// NewNode 创建一个新的抽象节点
func NewNode(name string) *Node {
	return &Node{name: name}
}

// Add 实现了节点接口的添加方法，抽象节点默认不添加子节点
func (n *Node) Add(child INode) {}

// Tree 实现了节点接口的展示方法，用于显示节点名称
func (n *Node) Tree(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(n.name)
}

// Folder 是目录类，实现了节点接口
type Folder struct {
	Node          // 继承抽象节点
	nodes []INode // 子节点列表
}

// NewFolder 创建一个新的目录
func NewFolder(name string) *Folder {
	return &Folder{
		Node: Node{name: name},
	}
}

// Add 实现了节点接口的添加方法，在目录中添加子节点
func (f *Folder) Add(children INode) {
	if len(f.nodes) == 0 {
		f.nodes = make([]INode, 0)
	}
	f.nodes = append(f.nodes, children)
}

// Tree 实现了节点接口的展示方法，显示目录下的子节点
func (f *Folder) Tree(space int) {
	f.Node.Tree(space)
	space++

	for _, node := range f.nodes {
		node.Tree(space)
	}
}

// File 是文件类，实现了节点接口
type File struct {
	Node // 继承抽象节点
}

// NewFile 创建一个新的文件
func NewFile(name string) *File {
	return &File{
		Node: Node{name: name},
	}
}

// Add 实现了节点接口的添加方法，文件类不允许添加子节点
func (f *File) Add(child INode) {
	fmt.Println("不允许添加子节点")
}

// Tree 实现了节点接口的展示方法，显示文件节点
func (f *File) Tree(space int) {
	f.Node.Tree(space)
}
