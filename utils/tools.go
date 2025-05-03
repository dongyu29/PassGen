package utils

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// 注意，为了效率，这个函数只会检查第一个字符是否是中文
func containsChinese(input string) bool {
	if len(input) == 0 {
		return false // 空字符串没有中文字符
	}
	return unicode.Is(unicode.Han, rune(input[0])) // 只检查第一个字符
}

// 提取并返回中文字符，同时去掉原字符串中的中文字符
func extractChinese(input *string) {
	var (
		result          strings.Builder // 存储去除中文后的字符串
		chineseFragment strings.Builder // 用于临时存储连续的中文字符
	)

	for _, r := range *input {
		if unicode.Is(unicode.Han, r) {
			// 累积连续的中文字符
			chineseFragment.WriteRune(r)
		} else {
			// 如果存在已收集的中文字符，加入切片
			if chineseFragment.Len() > 0 {
				Chineselist = append(Chineselist, chineseFragment.String())
				chineseFragment.Reset()
			}
			// 将非中文字符写回，忽略多余的逗号
			if !(r == ',' && (result.Len() == 0 || result.String()[result.Len()-1] == ',')) {
				result.WriteRune(r)
			}
		}
	}

	// 处理结尾的中文字符
	if chineseFragment.Len() > 0 {
		Chineselist = append(Chineselist, chineseFragment.String())
	}

	// 更新去除中文后的字符串
	// 移除首尾多余的逗号
	updatedResult := strings.Trim(result.String(), ",")
	*input = updatedResult
}

// 从文件读取内容并追加到目标列表
func parseFileList(filename string, list *[]string) error {
	if filename == "" {
		return nil
	}

	// 确保文件名包含 .txt 后缀
	if !strings.HasSuffix(filename, ".txt") {
		filename = filename + ".txt"
	}

	// 构建文件路径，位于 ../resources 目录下
	filePath := "resources/" + filename

	parsedList, err := parseFileToList(filePath)
	if err != nil {
		return err
	}
	*list = append(*list, parsedList...)

	// 提取每一行的中文字符
	for _, line := range parsedList {
		extractChinese(&line) // 提取中文字符
	}

	return nil
}

// WriteToFile 去重后将结果写入到文件
func WriteToFile(combinations []string, filename string) error {
	// 使用 map 去重
	uniqueCombinations := make(map[string]struct{})
	for _, combination := range combinations {
		uniqueCombinations[combination] = struct{}{}
	}

	// 打开 result.txt 文件，如果不存在则创建
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// 将去重后的组合写入到文件
	for combination := range uniqueCombinations {
		_, err := file.WriteString(combination + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}

// FilterCombinations 根据 A 和 B 是否为空来过滤组合规则
func FilterCombinations(rules string) string {
	// 判断 A 和 B 是否为空
	AExists := len(ChineseResult) > 0 // 判断 ChineseResult 是否为空
	BExists := len(TextResult) > 0    // 判断 TextResult 是否为空

	// 将传入的字符串拆分成切片
	combinations := strings.Split(rules, ",")
	var filtered []string

	// 情况 1: A 和 B 都存在, 只要 A 或 B 中的一个为 TRUE, 就保留规则
	if AExists && BExists {
		filtered = combinations // 直接保留所有组合
	} else if AExists && !BExists {
		// 情况 2: 只存在 A，B 不存在，A 必须为 TRUE
		for _, combination := range combinations {
			if strings.Contains(combination, "A") {
				filtered = append(filtered, combination)
			}
		}
	} else if !AExists && BExists {
		// 情况 3: 只存在 B，A 不存在，B 必须为 TRUE
		for _, combination := range combinations {
			if strings.Contains(combination, "B") {
				filtered = append(filtered, combination)
			}
		}
	} else {
		// 情况 4: A 和 B 都不存在，直接丢弃所有规则
		// 返回空字符串表示丢弃所有规则
		return ""
	}

	// 将过滤后的切片重新组合成逗号分隔的字符串
	return strings.Join(filtered, ",")
}

func RemoveEmptyStrings(slice *[]string) {
	var result []string
	for _, str := range *slice {
		if str != "" {
			result = append(result, str)
		}
	}
	*slice = result
}

// 清理四个切片
func CleanResultSlices() {
	// 直接修改原始切片内容
	RemoveEmptyStrings(&ChineseResult)
	RemoveEmptyStrings(&TextResult)
	RemoveEmptyStrings(&ConcatStrResult)
	RemoveEmptyStrings(&OtherResult)
}
