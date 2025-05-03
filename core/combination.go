package core

import (
	"PassGen/utils"
	"fmt"
	"sort"
	"strings"
)

// 组合规则的映射
var combinationRules = map[string][]string{
	"B": {"TextResult"},      // "B" 代表 TextResult
	"C": {"ConcatStrResult"}, // "C" 代表 ConcatStrResult
	"A": {"ChineseResult"},   // "A" 代表 ChineseResult
	"D": {"OtherResult"},     // "D" 代表 OtherResult（文件内容）
}

// CombinationsByRules 用于根据输入的多个组合规则字符串生成组合
func CombinationsByRules(Rules string) (map[string][]string, error) {
	// 结果存储
	results := make(map[string][]string)

	// 判断是否为 "all" 规则，若是则直接返回 A、B、C、D 的所有部分组合
	if Rules == "all" {
		// 获取 A、B、C、D 的所有部分组合
		allRules := []string{"A", "B", "C", "D"}
		permutations, err := generateAllPermutations(allRules)
		if err != nil {
			return nil, fmt.Errorf("error generating all permutations: %v", err)
		}
		for _, permutation := range permutations {
			result, err := GenerateCombinations(permutation)
			if err != nil {
				return nil, fmt.Errorf("Error generating combinations for rule %s: %v", permutation, err)
			}
			results[permutation] = result
		}
		return results, nil
	}

	// 否则处理传入的规则
	rules := strings.Split(Rules, ",")
	for _, rule := range rules {
		result, err := GenerateCombinations(rule)
		if err != nil {
			return nil, fmt.Errorf("Error generating combinations for rule %s: %v", rule, err)
		}
		results[rule] = result
	}

	return results, nil
}

// generateAllPermutations 生成部分组合
func generateAllPermutations(rules []string) ([]string, error) {
	var results []string
	// 递归生成所有可能的组合（包括部分组合）
	// 需要生成不同长度的排列
	for i := 1; i <= len(rules); i++ {
		permuteSubsetHelper(rules, i, []string{}, &results)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no permutations generated")
	}
	return results, nil
}

// permuteSubsetHelper 递归生成部分组合的辅助函数
func permuteSubsetHelper(input []string, length int, current []string, results *[]string) {
	// 如果当前组合达到所需长度，加入结果
	if len(current) == length {
		*results = append(*results, strings.Join(current, ""))
		return
	}

	// 递归遍历生成子集
	for i := 0; i < len(input); i++ {
		// 将当前元素加入组合
		permuteSubsetHelper(input[i+1:], length, append(current, input[i]), results)
	}
}

// GenerateCombinations 根据规则生成组合
func GenerateCombinations(rule string) ([]string, error) {
	var results []string
	var temp []string

	// 遍历组合规则中的每个字符
	for _, r := range rule {
		key := string(r)
		// 根据规则映射获取对应的切片名称
		listName, exists := combinationRules[key]
		if !exists {
			return nil, fmt.Errorf("unsupported rule: %s", key)
		}

		// 获取当前切片
		var currentList []string
		switch listName[0] {
		case "TextResult":
			currentList = utils.TextResult // 直接引用 utils 包中的 TextResult
		case "ConcatStrResult":
			currentList = utils.ConcatStrResult // 直接引用 utils 包中的 ConcatStrResult
		case "ChineseResult":
			currentList = utils.ChineseResult // 直接引用 utils 包中的 ChineseResult
		case "OtherResult":
			currentList = utils.OtherResult // 直接引用 utils 包中的 OtherResult
		}

		// 如果列表为空，则跳过该部分
		if len(currentList) == 0 {
			continue
		}

		// 组合当前列表的元素
		if len(temp) == 0 {
			// 如果没有临时结果，直接赋值
			results = append(results, currentList...)
		} else {
			// 否则，生成新的组合
			for _, item := range currentList {
				for _, existing := range temp {
					results = append(results, existing+item)
				}
			}
		}
		temp = results
		results = nil // 清空结果以便下次使用
	}

	return temp, nil
}

// Combinations 根据给定的规则生成组合，去重并排序
func Combinations(rules string) ([]string, error) {
	fmt.Println("生成组合密码......")
	// 获取基于规则生成的组合
	results, err := CombinationsByRules(rules)
	if err != nil {
		return nil, fmt.Errorf("error generating combinations: %v", err)
	}

	// 创建一个用于存放所有组合结果的切片
	var allCombinations []string

	// 遍历每个规则的组合，将它们添加到 allCombinations 中
	for _, result := range results {
		allCombinations = append(allCombinations, result...)
	}

	// 使用 map 去重
	uniqueCombinations := make(map[string]struct{})
	for _, combination := range allCombinations {
		uniqueCombinations[combination] = struct{}{}
	}

	// 将去重后的组合转换为切片
	var finalCombinations []string
	for combination := range uniqueCombinations {
		finalCombinations = append(finalCombinations, combination)
	}

	// 对去重后的组合进行排序
	sort.Strings(finalCombinations)

	// 返回去重排序后的结果
	return finalCombinations, nil
}