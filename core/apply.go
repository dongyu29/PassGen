package core

import (
	"PassGen/utils"
	"fmt"
	"strings"
	"sync"
)

// ApplyRules 根据规则集对不同类别的列表进行解析，并将结果存储到对应的结果列表中
func ApplyRules(level int, noRules string) {
	fmt.Println("调用规则.....")
	allRuleSets := GenerateRuleSets()
	var wg sync.WaitGroup
	resultsChan := make(chan struct {
		category string
		results  []string
	}, 10)

	// 将 noRules 参数解析为一个 map，方便快速查找要禁用的类别
	disabledCategories := make(map[string]bool)
	for _, category := range strings.Split(noRules, ",") {
		disabledCategories[strings.TrimSpace(category)] = true
	}

	// 启动收集 Goroutine，合并处理后的结果
	go func() {
		for res := range resultsChan {
			switch res.category {
			case "text":
				utils.TextResult = append(utils.TextResult, res.results...)
			case "concat":
				utils.ConcatStrResult = append(utils.ConcatStrResult, res.results...)
			case "chinese":
				utils.ChineseResult = append(utils.ChineseResult, res.results...)
			}
		}
	}()

	for _, ruleSet := range allRuleSets {
		var sourceList []string
		switch ruleSet.Category {
		case "text":
			if disabledCategories["b"] { // 如果 "b" 禁用则跳过 text
				continue
			}
			sourceList = utils.Textlist
		case "concat":
			if disabledCategories["c"] { // 如果 "c" 禁用则跳过 concat
				continue
			}
			sourceList = utils.ConcatStrlist
		case "chinese":
			if disabledCategories["a"] { // 如果 "a" 禁用则跳过 chinese
				continue
			}
			sourceList = utils.Chineselist
		default:
			continue
		}

		// 将所有符合等级的规则组合处理
		var combinedRules []Rule
		for lvl := 1; lvl <= level; lvl++ {
			if rules, exists := ruleSet.Levels[lvl]; exists {
				combinedRules = append(combinedRules, rules...)
			}
		}

		if len(combinedRules) > 0 {
			wg.Add(1)
			go func(rules []Rule, source []string, category string) {
				defer wg.Done()
				var transformedResults []string
				for _, item := range source {
					for _, rule := range rules {
						transformedResults = append(transformedResults, rule.Apply([]string{item})...)
					}
				}
				resultsChan <- struct {
					category string
					results  []string
				}{category, transformedResults}
			}(combinedRules, sourceList, ruleSet.Category)
		}
	}

	// 等待所有 Goroutine 完成，并关闭通道
	wg.Wait()
	close(resultsChan)
}
