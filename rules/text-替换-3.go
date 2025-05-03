package rules

import (
	"PassGen/core"
)

// ReplaceRule 将指定字符替换为特定的字符
type ReplaceRule struct{}

// 替换映射，指定每个字符的替换选项
var replacementMap = map[string][]string{
	"a": {"a", "4", "@"}, // 'a' 可以替换成 'a', '4' 或 '@'
	"b": {"b", "8"},
	"g": {"g", "9"},
	"i": {"i", "1"},
	"l": {"l", "1"},
	"o": {"o", "0"},
	"z": {"z", "2"},
}

// GenerateReplacements 生成所有替换后的变体
func GenerateReplacements(input string) []string {
	var variants []string
	var helper func(prefix string, idx int)

	// helper 函数生成替换组合
	helper = func(prefix string, idx int) {
		// 如果已处理完所有字符，添加当前组合到变体列表
		if idx == len(input) {
			variants = append(variants, prefix)
			return
		}

		// 当前字符
		currentChar := string(input[idx])

		// 如果当前字符有替换选项，则尝试所有替换
		if options, found := replacementMap[currentChar]; found {
			for _, option := range options {
				// 递归生成所有可能的组合
				helper(prefix+option, idx+1)
			}
		} else {
			// 没有替换规则的字符直接加入
			helper(prefix+currentChar, idx+1)
		}
	}

	// 从空前缀和输入字符串的第一个字符开始生成变体
	helper("", 0)
	return variants
}

// Apply 实现规则接口，将输入的字符串进行替换并生成所有变体
func (r ReplaceRule) Apply(inputs []string) []string {
	var result []string

	// 遍历每个输入字符串，生成所有变体
	for _, input := range inputs {
		variants := GenerateReplacements(input)
		result = append(result, variants...)
	}

	// 返回所有替换后的变体
	return result
}

// init 函数用于注册规则
func init() {
	core.RegisterRule("text", ReplaceRule{}, 3)
}
