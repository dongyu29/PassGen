package rules

import (
	"PassGen/core"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// PinyinRule 将中文转换为拼音
type PinyinRule3 struct{}

// Apply 实现规则接口，将输入的中文切片转换为拼音切片
func (r PinyinRule3) Apply(inputs []string) []string {
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

		// 将拼音列表合并为一个字符串并添加到变体列表
		variants = append(variants, strings.Join(pinyinWord, ""))
	}

	// 返回拼音的变体
	return variants
}

// init 函数用于注册规则
func init() {
	core.RegisterRule("chinese", PinyinRule3{}, 3)
}
