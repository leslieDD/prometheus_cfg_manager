package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"runtime"
)

type PersonCommonInfo struct {
	ProgramerSay []string `json:"programer_say"`
	Uname        string   `json:"uname"`
}

func GetProgramerSay() (*PersonCommonInfo, *BriefMessage) {
	pci := PersonCommonInfo{
		ProgramerSay: ProgramerSay,
	}
	if runtime.GOOS == "windows" {
		hostInfo := utils.HostInfo()
		pci.Uname = fmt.Sprintf("%s %s", hostInfo.Platform, hostInfo.PlatformFamily)
	} else {
		_, output, err := utils.ExecCmd("uname", "-a")
		if err != nil {
			config.Log.Error(err)
			hostInfo := utils.HostInfo()
			pci.Uname = fmt.Sprintf("%s %s", hostInfo.Platform, hostInfo.PlatformFamily)
		} else {
			pci.Uname = output
		}
	}
	return &pci, Success
}

var ProgramerSay = []string{
	"大部分情况下，构建程序的过程本质上是对规范调试的过程。——Fred Brooks，《人月神话》作者",
	"软件开发往往是这样：最开始的 90% 代码占用了开始的 90% 的开发时间；剩下 10% 代码同样需要 90% 的开发时间。——Tom Cargill",
	"当你试图解决一个你不理解的问题时，复杂化就产成了。——Andy Boothe",
	"生命太短暂，不要去做一些根本没有人想要的东西。——Ash Maurya",
	"任何一个傻瓜都会写能够让机器理解的代码，只有好的程序员才能写出人类可以理解的代码。——Martin Fowler",
	"靠代码行数来衡量开发进度，就像是凭重量来衡量飞机制造的进度。——比尔·盖茨", "这不是一个 bug，这只是一个未列出来的特性。——匿名",
	"作为一个程序员，郁闷的事情是，面对一个代码块，却不敢去修改。更糟糕的是，这个代码块还是自己写的。—— Peyton Jones",
	"它在我的机器上可以很好运行！——大部分程序员",
	"能说算不上什么，有本事就把你的代码给我看看。——Linus Torvalds，Linux 之父",
	"如果你交给某人一个程序，你将折磨他一整天；如果你教某人如何编写程序，你将折磨他一辈子。——David Leinweber", "软件设计有两种方式：一种方式是，使软件过于简单，明显没有缺陷；另一种方式是，使软件过于复杂，没有明显的缺陷。——C.A.R. Hoare", "其实，我尝试着使 Ruby 更自然，而不是简单。Ruby 看起来很简单，但内部是非常复杂的，就像我们的身体一样。——松本行弘，Ruby 之父", "用几个小时来制定计划，可以节省几周的编程时间。—— 匿名",
	"控制复杂性是计算机编程的本质。—— Brian Kernighan",
	"计算机科学领域的所有问题都可以通过其他方式间接解决。——David Wheeler",
	"编程是两队人马在竞争：软件工程师努力设计出最大最好的连白痴都会使用的程序；而宇宙在拼命制造最大最好的白痴。到目前为止，宇宙是胜利者。—— Rick Cook",
	"调试一个初次见到的代码比重写代码要困难两倍。因此，按照定义，如果你写代码非常巧妙，那么没有人足够聪明来调试它。—— Brian W. Kernighan",
	"我不是一个伟大的程序员，我只是一个具有良好习惯的优秀程序员。― Kent Beck",
	"你们中大多数人都熟悉程序员的美德，有三种：那就是懒惰、急躁和傲慢。– Larry Wall，Perl 語言发明人",
	"我认为对象就像是生物学里的细胞，或者网络中的一台计算机，只能够通过消息来通信——Alan Kay，Smalltalk 的发明人，面向对象之父",
	"当你选择了一种语言，意味着你还选择了一组技术、一个社区。——Joshua Bloch",
	"质量、速度、廉价，选择其中两个。——匿名", "过早的优化是罪恶之源。——Donald Knuth",
	"没有什么代码的执行速度比空代码更快。——Merb 核心原则",
	"如果你是房间里最聪明的人，那么你走错房间了。——匿名",
	"如果只需更改一个单一的代码行，你的部门需要花费多长时间？——Mary Poppendieck", "九个人不能让一个孩子在一个月内出生。——Fred Brooks，《人月神话》作者",
	"好代码本身就是最好的文档。当你需要添加一个注释时，你应该考虑如何修改代码才能不需要注释。——Steve McConnell，Code Complete 作者",
	"一个人在教会电脑之前，别说他真正理解这个东西了。——Donald Knuth",
	"无论在排练中演示是如何的顺利(高效)，当面对真正的现场观众时，出现错误的可能性跟在场观看的人数成正比。——佚名",
	"罗马帝国崩溃的一个主要原因是，没有他们没有有效的方法表示他们的C程序成功的终止。——Robert Firth", "C程序员永远不会灭亡。他们只是cast成了void。——佚名",
	"如果debugging是一种消灭bug的过程，那编程就一定是把bug放进去的过程。——Edsger Dijkstra",
	"你要么要软件质量，要么要指针算法；两者不可兼得。——(Bertrand Meyer)", "有两种方法能写出没有错误的程序；但只有第三种好用。——Alan J. Perlis",
	"最初的90%的代码用去了最初90%的开发时间。余下的10%的代码用掉另外90%的开发时间。——Tom Cargill",
	"程序员和上帝打赌要开发出更大更好——傻瓜都会用的软件。而上帝却总能创造出更大更傻的傻瓜。所以，上帝总能赢。——Anon",
	"你是否也有自己喜欢的名言，欢迎分享。", "UNIX很简单。但需要有一定天赋的人才能理解这种简单。——Dennis Ritchie",
	"有两种软件设计的方式：一种是使它足够简单以致于明显没有缺陷，另一种则是使它足够复杂以致于没有明显的缺陷。 ——C.A.R. Hoare",
	"软件工程的目标是控制复杂度，而不是增加复杂性。——Dr. Pamela Zave",
	"软件在能够复用前必须先能用。——Ralph Johnson",
	"优秀的判断力来自经验，但经验来自于错误的判断。——Fred Brooks",
	"‘理论’是你知道是这样，但它却不好用。‘实践’是它很好用，但你不知道是为什么。程序员将理论和实践结合到一起：既不好用，也不知道是为什么。——佚名",
	"当你想在你的代码中找到一个错误时，这很难；当你认为你的代码是不会有错误时，这就更难了。——Steve McConnell 《代码大全》",
	"如果建筑工人盖房子的方式跟程序员写程序一样，那第一只飞来的啄木鸟就将毁掉人类文明。——Gerald Weinberg",
	"项目开发的六个阶段：1. 充满热情 2. 醒悟 3. 痛苦 4. 找出罪魁祸首 5. 惩罚无辜 6. 褒奖闲人——佚名",
	"优秀的代码是它自己最好的文档。当你考虑要添加一个注释时，问问自己，“如何能改进这段代码，以让它不需要注释？”——Steve McConnell 《代码大全》", "我们这个世界的一个问题是，蠢人信誓旦旦，智人满腹狐疑。——Bertrand Russell",
}
