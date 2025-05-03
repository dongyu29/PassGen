package utils

// 基础元素输入
type Element struct {
	Text          string
	ConcatStr     string
	TextFile      string
	ConcatStrFile string
	Other         string
}

// 元素解析List
var (
	Textlist      []string
	ConcatStrlist []string
	Chineselist   []string
)

// 规则处理完的结果
var (
	TextResult      []string
	ConcatStrResult []string
	ChineseResult   []string
	OtherResult     []string
)

// 拓展参数
var (
	RuleLevel   int
	IsYear      bool
	ConcatRules string
	NoRules     string
)

// 默认参数
var (
	DefaultRules = "ADCB,ACDB,ACD,ADC,BCD,BDC,ACB,ADB,BCA,BDA,BC,BD,AC,AD,A,B"
)
