package rules

import (
	"PassGen/core"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// CaseCombinationRule 生成字符串的首部连续大写组合，支持汉字转换为拼音后处理
type CaseCombinationRule1 struct{}

// Apply 实现规则接口，将输入的字符串切片转换为拼音并生成首部连续大写的组合
func (r CaseCombinationRule1) Apply(inputs []string) []string {
	var variants []string
	for _, input := range inputs {
		// 转换汉字为拼音
		pinyinString := convertToPinyin(input)
		// 生成大写组合
		variants = append(variants, generateCase1(pinyinString)...)
	}
	return variants
}

// convertToPinyin 将汉字转换为拼音
func convertToPinyin(s string) string {
	a := pinyin.NewArgs()
	a.Heteronym = false // 不使用多音字

	pinyinResult := pinyin.Pinyin(s, a)

	var pinyinWords []string
	for _, syllables := range pinyinResult {
		if len(syllables) > 0 {
			pinyinWords = append(pinyinWords, syllables[0]) // 获取每个字的拼音
		}
	}
	// 拼接为单个字符串
	return strings.Join(pinyinWords, "")
}

// generateCase 生成字符串的首部连续大写组合
func generateCase1(s string) []string {
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
	core.RegisterRule("chinese", CaseCombinationRule1{}, 3)
}
