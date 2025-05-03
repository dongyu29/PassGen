package rules

import (
	"PassGen/core"
	"strings"
)

// UppercaseRule 将字符串转换为大写
type UppercaseRule struct{}

// Apply 实现规则接口，将输入的字符串切片转换为大写字母
func (r UppercaseRule) Apply(inputs []string) []string {
	var result []string

	// 遍历每个输入字符串，转换为大写
	for _, input := range inputs {
		result = append(result, strings.ToUpper(input))
	}

	// 返回转换为大写的字符串切片
	return result
}

// init 函数用于注册规则
func init() {
	core.RegisterRule("text", UppercaseRule{}, 1)
}
