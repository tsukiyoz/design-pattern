package composite

func ExampleComposite() {
	driveD := NewFolder("D盘")

	// 创建文档目录及文件
	docDir := NewFolder("文档")
	docDir.Add(NewFile("简历.doc"))
	docDir.Add(NewFile("项目介绍.ppt"))
	driveD.Add(docDir)

	// 创建音乐目录及文件
	musicDir := NewFolder("音乐")

	jayDir := NewFolder("周杰伦")
	jayDir.Add(NewFile("发如雪.mp3"))
	jayDir.Add(NewFile("青花瓷.mp3"))

	easonDir := NewFolder("陈奕迅")
	easonDir.Add(NewFile("十年.mp3"))
	easonDir.Add(NewFile("富士山下.mp3"))

	musicDir.Add(jayDir)
	musicDir.Add(easonDir)

	driveD.Add(musicDir)

	driveD.Tree(0) // 显示整个节点树
	// Output:
	// D盘
	//  文档
	//   简历.doc
	//   项目介绍.ppt
	//  音乐
	//   周杰伦
	//    发如雪.mp3
	//    青花瓷.mp3
	//   陈奕迅
	//    十年.mp3
	//    富士山下.mp3
}
