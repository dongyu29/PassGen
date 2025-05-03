package rules

import (
	"PassGen/core"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type PartialAbbrPinyinRule struct{}

// Apply 将输入的中文字符转换为简拼和部分全拼的不同大小写组合，遵循大写规则
func (r PartialAbbrPinyinRule) Apply(inputs []string) []string {
	var variants []string

	for _, input := range inputs {
		a := pinyin.NewArgs()
		a.Heteronym = false
		pinyinResult := pinyin.Pinyin(input, a)

		// 拼音首字母简拼和部分全拼组合
		var abbrev []string
		for i, syllables := range pinyinResult {
			if i == 0 || i == 1 { // 保持前两个字的全拼
				abbrev = append(abbrev, syllables[0])
			} else { // 后面的字只保留首字母
				abbrev = append(abbrev, string(syllables[0][0]))
			}
		}
		abbrevStr := strings.Join(abbrev, "")

		// 生成符合大小写规则的所有可能组合
		variants = append(variants, generateCaseCombinations(abbrevStr)...)
	}

	return variants
}

// generateCaseCombinations 生成特定规则的大小写组合
func generateCaseCombinations(abbrev string) []string {
	var combinations []string

	// 全小写组合
	combinations = append(combinations, strings.ToLower(abbrev))

	// 逐步增加大写字母的组合
	for i := 1; i <= len(abbrev); i++ {
		comb := strings.ToUpper(abbrev[:i]) + strings.ToLower(abbrev[i:])
		combinations = append(combinations, comb)
	}

	return combinations
}

// init 函数用于注册规则
func init() {
	core.RegisterRule("chinese", PartialAbbrPinyinRule{}, 3)
}
