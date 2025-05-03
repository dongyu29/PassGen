package core

// Rule 接口
type Rule interface {
	Apply(input []string) []string
}

// 规则集
type RuleSet struct {
	Category string
	Levels   map[int][]Rule // 每个等级对应一个规则列表
}

// 自动注册规则列表
var ruleRegistry = make(map[string]*RuleSet)

// RegisterRule 注册规则
func RegisterRule(category string, rule Rule, level int) {
	// 获取或创建对应的 RuleSet
	ruleSet, exists := ruleRegistry[category]
	if !exists {
		ruleSet = &RuleSet{
			Category: category,
			Levels:   make(map[int][]Rule),
		}
		ruleRegistry[category] = ruleSet
	}

	// 将规则添加到对应等级的规则列表
	ruleSet.Levels[level] = append(ruleSet.Levels[level], rule)
}

// GenerateRuleSets 动态生成规则集，直接返回所有规则集的切片
func GenerateRuleSets() []*RuleSet {
	// 创建 RuleSet 的切片
	ruleSets := []*RuleSet{}

	// 遍历原始的 ruleRegistry，直接将指针添加到切片中
	for _, ruleSet := range ruleRegistry {
		ruleSets = append(ruleSets, ruleSet)
	}

	return ruleSets
}
