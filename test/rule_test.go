package test

import (
	"PassGen/rules"
	"fmt"
	"testing"
)

func TestAbbrPinyinRule_Output(t *testing.T) {
	// 初始化规则
	abbrRule := rules.PinyinRule{}

	// 定义测试用例
	testCases := []string{
		"平某某大好人",
	}

	// 遍历每个测试用例并输出转换结果
	for _, input := range testCases {
		result := abbrRule.Apply([]string{input})
		fmt.Printf("Input: %s => AbbrPinyin: %v\n", input, result)
	}
}
