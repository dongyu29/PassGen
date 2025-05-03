package rules

import (
	"PassGen/core"
	"strings"
)

// LowercaseRule 将字符串转换为小写
type LowercaseRule struct{}

// Apply 实现规则接口，将输入的字符串切片转换为小写字母
func (r LowercaseRule) Apply(inputs []string) []string {
	var result []string

	// 遍历每个输入字符串，转换为小写
	for _, input := range inputs {
		result = append(result, strings.ToLower(input))
	}

	// 返回转换为小写的字符串切片
	return result
}

// init 函数用于注册规则
func init() {
	// 注册 LowercaseRule 规则到 "text" 类别，等级为 1
	core.RegisterRule("text", LowercaseRule{}, 1)
}
