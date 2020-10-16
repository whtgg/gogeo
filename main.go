package main

import (
	"flag"
	"fmt"
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	text = "在今天教育部举办的新闻发布会上，教育部体育卫生与艺术教育司司长王登峰表示，学校的体育中考要不断总结经验，逐年增加分值，要达到跟语数外同分值的水平。“目前全国有一家，云南省已经做到了从今年开始体育中考跟语文数学外语一样都是100分。而在此基础上，我们通过不断总结经验，立即启动体育在高考中计分的研究。”\n\n近日，中共中央办公厅、国务院办公厅印发了《关于全面加强和改进新时代学校体育工作的意见》和《关于全面加强和改进新时代学校美育工作的意见》（以下简称《意见》），就全面贯彻党的教育方针，加强和改进新时代学校体育、美育工作进行了系统设计和全面部署。\n\n王登峰表示，“文件进一步明确了体育中考的方向，而且要进一步扩大它的分值和影响力，同时开展研究在高考录取中体育素养如何评价、如何计分。这对于学校体育，特别是体育作为学生升学评价体系的一个重要组成部分带有革命性的提法或者观点，这也是我们下一步要重点推进的一项工作。而关于美育，跟体育一样，除了前面的综合性评价以外，美育中考要在试点基础上尽快推广。”\n\n此外，王登峰表示，目前全国已经有4个省开展美育中考计分，同时还有6个省、12个地市已经开始了中考美育的计分，分值在10分到40分之间，云南省从今年开始要增加到40分。“这个试点已经有了相当的基础，下一步我们要把美育中考的工作做的更加扎实，到2022年力争全覆盖，全面实行美育中考。”"
	seg gse.Segmenter
	posSeg pos.Segmenter
)

func cut() {
	seg.LoadDict()
	//分词 accurate mode
	hmm := seg.Cut(text, true) //true 表示HMM 算法
	fmt.Println("cut: ", hmm) //标点符号不会去掉
	trimHmm := seg.Trim(hmm)
	fmt.Println("trimHmm all: ", trimHmm)//去标点符号分词 返回list
	fmt.Println("trimHmm all: ", len(trimHmm))//去标点符号分词 返回list

	//分词查找 search engine mode
	trimSearch := seg.Trim(seg.CutSearch(text,true))
	fmt.Println("trimSearch",trimSearch)
	fmt.Println("trimSearch",len(trimSearch))

	//分词查找 full mode
	trimCutAll := seg.Trim(seg.CutAll(text))
	fmt.Println("trimCutAll",trimCutAll)
	fmt.Println("trimCutAll",len(trimCutAll))

	//分词返回字符串并且保持原文的特殊符号
	//segString := seg.String(text,true)
	//fmt.Println("segString",segString)
	//Find
	findFre,ok := seg.Find("教育部")//在词库里面查找
	fmt.Println(findFre) //出现频率 ？？？ 不知频率的意思
	fmt.Println(ok) //是否存在
	//自定义词库
	seg.LoadDict("./data/data1.txt") //自定义词库
	findFreC,_ := seg.Find("教育部门累")//在词库里面查找
	fmt.Println("findFreC",findFreC)
	//手动添加词库
	seg.AddToken("教育部门累红",87986,"n")
	findFreT,_ := seg.Find("教育部门累红")//在词库里面查找
	fmt.Println("findFreT",findFreT)

	//位置 pos
	posSeg.WithGse(seg)
	pp := posSeg.Cut(text,true)
	for t,p := range pp {
		fmt.Println("text",t)
		fmt.Println("pos",p)
	}
	fmt.Println(posSeg.TrimWithPos(pp,"x")) //去掉位置为x分词
}

func main() {
	flag.Parse()
	cut()
}
