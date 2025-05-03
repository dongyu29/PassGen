package rules

import (
	"PassGen/core"
	"strings"
)

// CaseCombinationRule 生成字符串的首部连续大写组合
type CaseCombinationRule struct{}

// Apply 实现规则接口，将输入的字符串切片生成首部连续大写的组合
func (r CaseCombinationRule) Apply(inputs []string) []string {
	var variants []string
	for _, input := range inputs {
		variants = append(variants, generateCase(input)...)
	}
	return variants
}

// generateCaseCombinations 生成字符串的首部连续大写组合
func generateCase(s string) []string {
	var combinations []string
	n := len(s)

	// 遍历首部需要大写的字符数量，1 到 n 个
	for i := 1; i <= n; i++ {
		var sb strings.Builder
		// 将前 i 个字符转换为大写
		for j := 0; j < i; j++ {
			sb.WriteByte(s[j] - 'a' + 'A')
		}
		// 其余字符保持小写
		sb.WriteString(s[i:])
		combinations = append(combinations, sb.String())
	}

	return combinations
}

// init 函数用于注册大小写组合规则
func init() {
	core.RegisterRule("text", CaseCombinationRule{}, 3)
}
