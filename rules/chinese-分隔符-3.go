package rules

import (
	"PassGen/core"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// PinyinRuleWithSeparators 将中文转换为拼音，并提供带分隔符的拼音形式
type PinyinRuleWithSeparators struct{}

// Apply 实现规则接口，将输入的中文切片转换为拼音切片
func (r PinyinRuleWithSeparators) Apply(inputs []string) []string {
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
		// 获取带下划线分隔符的拼音
		withUnderscore := strings.Join(pinyinWord, "_")
		// 获取带连字符分隔符的拼音
		withDash := strings.Join(pinyinWord, "-")
		// 获取带点分隔符的拼音
		withDot := strings.Join(pinyinWord, ".")
		// 获取拼音的首字母
		abbreviations := getAbbreviations(pinyinWord)

		// 将结果合并到变体列表
		variants = append(variants, fullPinyin, withUnderscore, withDash, withDot)
		variants = append(variants, abbreviations...)
	}

	// 返回拼音的变体
	return variants
}

// getAbbreviations 获取拼音的首字母组合（例如：l_x_b）
func getAbbreviations(pinyinWord []string) []string {
	var abbreviations []string

	// 获取拼音的首字母（如：l_x_b）
	for _, word := range pinyinWord {
		abbreviations = append(abbreviations, string(word[0]))
	}

	// 返回所有拼音首字母的组合，使用 _ 分隔
	abbreviationsWithUnderscore := strings.Join(abbreviations, "_")
	abbreviationsWithDash := strings.Join(abbreviations, "-")
	abbreviationsWithDot := strings.Join(abbreviations, ".")

	// 将不同的组合方式加入变体
	return []string{
		abbreviationsWithUnderscore,
		abbreviationsWithDash,
		abbreviationsWithDot,
	}
}

// init 函数用于注册规则
func init() {
	core.RegisterRule("chinese", PinyinRuleWithSeparators{}, 3)
}
