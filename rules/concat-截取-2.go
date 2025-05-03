package rules

import (
	"PassGen/core"
)

// CliceRule 截取字符串的前2 个，3 个字符
type CliceRule struct{}

// Apply 实现规则接口，对每个输入的字符串进行截取
func (r CliceRule) Apply(inputs []string) []string {
	var result []string

	// 遍历每个输入字符串
	for _, input := range inputs {
		// 截取前 2 个字符
		if len(input) >= 2 {
			result = append(result, input[:2])
		}
		// 截取前 3 个字符
		if len(input) >= 3 {
			result = append(result, input[:3])
		}
	}

	// 返回所有截取的字符组合
	return result
}

// init 函数用于注册规则
func init() {
	core.RegisterRule("concat", CliceRule{}, 2)
}
