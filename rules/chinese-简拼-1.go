package rules

import (
	"PassGen/core"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// AbbrPinyinRule 将中文转换为各种格式的拼音简写
type AbbrPinyinRule struct{}

// Apply 实现规则接口，将输入的中文切片转换为拼音简写切片
func (r AbbrPinyinRule) Apply(inputs []string) []string {
	var variants []string

	// 遍历每个输入字符串，转换为拼音简写
	for _, input := range inputs {
		a := pinyin.NewArgs()
		a.Heteronym = false // 是否使用多音字

		pinyinResult := pinyin.Pinyin(input, a)

		// 提取每个字的拼音首字母
		var abbr []string
		for _, syllables := range pinyinResult {
			if len(syllables) > 0 {
				abbr = append(abbr, syllables[0][:1]) // 获取每个字的首字母
			}
		}

		abbrString := strings.Join(abbr, "")
		// 添加符合规则的大小写组合
		variants = append(variants, generateCombinations(abbrString)...)
	}

	// 返回所有组合的变体
	return variants
}

// generateCombinations 生成字符串的各种大小写组合，前面的字母不大写则后面的字母也不大写
func generateCombinations(s string) []string {
	var combinations []string
	n := len(s)

	// 使用递归生成符合规则的组合
	var dfs func(int, string)
	dfs = func(index int, current string) {
		if index == n {
			combinations = append(combinations, current)
			return
		}

		// 保持当前字符小写并继续递归
		dfs(index+1, current+string(s[index]))

		// 如果当前字符前面全是大写，才允许当前字符大写
		if index == 0 || (len(current) > 0 && current[index-1] >= 'A' && current[index-1] <= 'Z') {
			// 将当前字符大写并继续递归
			dfs(index+1, current+string(s[index]-'a'+'A'))
		}
	}

	// 从第一个字符开始生成组合
	dfs(0, "")

	return combinations
}

// init 函数用于注册规则
func init() {
	core.RegisterRule("chinese", AbbrPinyinRule{}, 1)
}
