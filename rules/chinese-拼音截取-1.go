package rules

import (
	"PassGen/core"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// PinyinCliceRule3 将中文转换为拼音，并提供部分截取和排列功能
// 例如 平某某大好人，截取前三个字符平某某，[ping Ping PING pingmou Pingmou PINGMOU pingmoumou Pingmoumou PINGMOUMOU]
type PinyinCliceRule3 struct{}

// Apply 实现规则接口，将输入的中文切片转换为拼音切片
func (r PinyinCliceRule3) Apply(inputs []string) []string {
	var variants []string

	// 遍历每个输入字符串，转换为拼音
	for _, input := range inputs {
		// 获取前三个汉字
		trimmedInput := getFirstThreeChinese(input)

		a := pinyin.NewArgs()
		a.Heteronym = false // 不使用多音字

		pinyinResult := pinyin.Pinyin(trimmedInput, a)

		var pinyinWord []string

		// 将拼音结果拼接为一个完整的拼音
		for _, syllables := range pinyinResult {
			if len(syllables) > 0 {
				pinyinWord = append(pinyinWord, syllables[0]) // 拼接每个字的拼音
			}
		}

		// 获取拼音的分段组合，并生成不同格式的变体
		partCombinations := getSegmentedCombinationsWithCaseVariants(pinyinWord)

		// 将结果合并到变体列表
		variants = append(variants, partCombinations...)
	}

	// 返回拼音的变体
	return variants
}

// getFirstThreeChinese 获取输入字符串中的前三个汉字
func getFirstThreeChinese(input string) string {
	var result strings.Builder
	count := 0

	// 遍历字符串并找到前3个汉字
	for _, ch := range input {
		// 判断是否为汉字
		if isChinese(ch) {
			result.WriteRune(ch)
			count++
		}
		if count >= 3 {
			break
		}
	}

	return result.String()
}

// isChinese 判断字符是否为汉字
func isChinese(r rune) bool {
	return r >= 0x4e00 && r <= 0x9fff
}

// getSegmentedCombinationsWithCaseVariants 获取拼音的分段组合及其不同格式的变体
func getSegmentedCombinationsWithCaseVariants(pinyinWord []string) []string {
	var combinations []string
	n := len(pinyinWord)

	// 遍历不同的组合长度
	for i := 1; i <= n; i++ {
		// 每次组合前 i 个拼音
		segment := strings.Join(pinyinWord[:i], "")
		combinations = append(combinations, segment)                  // 原始格式
		combinations = append(combinations, capitalize(segment))      // 首字母大写
		combinations = append(combinations, strings.ToUpper(segment)) // 全部大写
	}

	return combinations
}

// capitalize 将字符串的首字母转换为大写，其他字母保持不变
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// init 函数用于注册拼音截取规则
func init() {
	core.RegisterRule("chinese", PinyinCliceRule3{}, 1)
}
