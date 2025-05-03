package rules

import (
	"PassGen/utils"
	"fmt"
)

func Test() {
	fmt.Println("规则已加载")
	// 将 Textlist 的内容添加到 TextResult
	utils.TextResult = append(utils.TextResult, utils.Textlist...)

	// 将 ConcatStrlist 的内容添加到 ConcatStrResult
	utils.ConcatStrResult = append(utils.ConcatStrResult, utils.ConcatStrlist...)

	// 将 Chineselist 的内容添加到 ChineseResult
	utils.ChineseResult = append(utils.ChineseResult, "")
}
