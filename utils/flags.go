package utils

import (
	"flag"
)

func Banner() {
	banner := `

8888888b.                              .d8888b.                    
888   Y88b                            d88P  Y88b                   
888    888                            888    888                   
888   d88P  8888b.  .d8888b  .d8888b  888         .d88b.  88888b.  
8888888P"      "88b 88K      88K      888  88888 d8P  Y8b 888 "88b 
888        .d888888 "Y8888b. "Y8888b. 888    888 88888888 888  888 
888        888  888      X88      X88 Y88b  d88P Y8b.     888  888 
888        "Y888888  88888P'  88888P'  "Y8888P88  "Y8888  888  888 


            Version: v1.0
            Author:东隅

`
	print(banner)
}

func Flag(Elem *Element) {
	Banner()
	flag.StringVar(&Elem.Text, "ab", "", "输入常规字符，进行中文转换、字符替换、大小写变换等")
	flag.StringVar(&Elem.TextFile, "abf", "", "常规字符文件")
	flag.StringVar(&Elem.ConcatStr, "c", "", "输入拼接字符，比如QQ号")
	flag.StringVar(&Elem.ConcatStrFile, "cf", "", "拼接字符文件")
	flag.StringVar(&ConcatRules, "r", DefaultRules, "拼接规则，常规字符是A，拼接字符是B，内置字符是C")
	flag.StringVar(&NoRules, "n", "", "不对某个类别使用规则，例如：-n a,b,c")
	flag.IntVar(&RuleLevel, "l", 2, "使用的规则等级")
	flag.StringVar(&Elem.Other, "d", "", "使用的拼接字典，直接参与密码生成")
	flag.Parse()
}
