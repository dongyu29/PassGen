package rules

import (
	"PassGen/core"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// PinyinCliceRule 将中文转换为拼音，并提供部分截取和排列功能
type PinyinCliceRule struct{}

// Apply 实现规则接口，将输入的中文切片转换为拼音切片
func (r PinyinCliceRule) Apply(inputs []string) []string {
	var variants []string

	// 遍历每个输入字符串，转换为拼音
	for _, input := range inputs {
		a := pinyin.NewArgs()
		a.Heteronym = false // 是否使用多音字

		pinyinResult := pinyin.Pinyin(input, a)

		var pinyinWord []string

		// 将拼音结果拼接为一个完整的拼音
		for _, syllables := range pinyinResult {
			if len(syllables) > 0 {
				pinyinWord = append(pinyinWord, syllables[0]) // 拼接每个字的拼音
			}
		}

		// 获取拼音的全拼
		fullPinyin := strings.Join(pinyinWord, "")
		// 获取拼音的部分拼音组合
		partCombinations := getContinuousPartCombinations(pinyinWord)

		// 将结果合并到变体列表
		variants = append(variants, fullPinyin)
		variants = append(variants, partCombinations...)
	}

	// 返回拼音的变体
	return variants
}

// getContinuousPartCombinations 获取拼音的连续部分拼音组合
func getContinuousPartCombinations(pinyinWord []string) []string {
	var combinations []string
	n := len(pinyinWord)

	// 遍历所有的连续子序列
	for start := 0; start < n; start++ {
		for end := start + 1; end <= n; end++ {
			part := strings.Join(pinyinWord[start:end], "")
			combinations = append(combinations, part)
		}
	}

	return combinations
}

// init 函数用于注册拼音截取规则
func init() {
	core.RegisterRule("chinese", PinyinCliceRule{}, 3)
}
