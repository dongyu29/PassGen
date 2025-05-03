package main

import (
	"PassGen/core"
	"PassGen/rules"
	"PassGen/utils"
	"fmt"
)

func main() {
	var Elem utils.Element
	utils.Flag(&Elem)
	utils.Parse(&Elem)
	rules.Test()
	core.ApplyRules(utils.RuleLevel, utils.NoRules)

	//去除结果切片为空的字符串
	utils.CleanResultSlices()
	fmt.Println("ChineseResult:", utils.ChineseResult)
	fmt.Println("TextResult:", utils.TextResult)
	fmt.Println("ConcatStrResult:", utils.ConcatStrResult)
	fmt.Println("OtherResult:", utils.OtherResult)

	//拼接之前先过滤一下有效的规则
	rule := utils.FilterCombinations(utils.ConcatRules)
	finalCombinations, err := core.Combinations(rule)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("结果如下")
	for _, combination := range finalCombinations {
		fmt.Println(combination)
	}
	// 调用 WriteToFile 函数将去重后的组合写入到文件
	err = utils.WriteToFile(finalCombinations, "result.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
