package utils

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// ConvertToPinyin 将一组字符串转换为拼音
func ConvertToPinyin(inputs []string) []string {
	var variants []string

	// 初始化拼音转换参数
	a := pinyin.NewArgs()
	a.Heteronym = false // 不使用多音字

	// 遍历输入的每个字符串，转换为拼音
	for _, input := range inputs {
		// 获取每个字符的拼音
		pinyinResult := pinyin.Pinyin(input, a)
		var pinyinWord []string

		// 将拼音结果拼接为一个完整的拼音词
		for _, syllables := range pinyinResult {
			if len(syllables) > 0 {
				pinyinWord = append(pinyinWord, syllables[0]) // 拼接每个字的拼音
			}
		}

		// 将拼音词转换为字符串并添加到结果列表中
		variants = append(variants, strings.Join(pinyinWord, ""))
	}

	return variants
}

// generateCombinations 生成字符串的各种大小写组合
func generateCombinations(s string) []string {
	n := len(s)
	totalCombinations := 1 << n
	var combinations []string

	for i := 0; i < totalCombinations; i++ {
		var sb strings.Builder
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 {
				sb.WriteByte(s[j]) // 小写
			} else {
				sb.WriteByte(s[j] - 'a' + 'A') // 大写
			}
		}
		combinations = append(combinations, sb.String())
	}

	return combinations
}
