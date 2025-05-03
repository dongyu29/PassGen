package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 把输入的元素解析到list
func Parse(Elem *Element) error {
	fmt.Println("解析参数....")
	// 提取 Text 字段中的中文字符并去掉
	extractChinese(&Elem.Text)

	// 解析 Text 字段
	parseStringList(Elem.Text, &Textlist)

	// 提取 ConcatStr 字段中的中文字符并去掉
	extractChinese(&Elem.ConcatStr)

	// 解析 ConcatStr 字段
	parseStringList(Elem.ConcatStr, &ConcatStrlist)

	// 解析文件
	if err := parseFileList(Elem.TextFile, &Textlist); err != nil {
		return err
	}
	if err := parseFileList(Elem.ConcatStrFile, &ConcatStrlist); err != nil {
		return err
	}
	if err := parseFileList(Elem.Other, &OtherResult); err != nil {
		return err
	}

	return nil
}

// 将字符串按逗号分割并追加到目标列表
func parseStringList(input string, list *[]string) {
	if input != "" {
		parsed := strings.Split(input, ",")
		*list = append(*list, parsed...)
	}
}

// 解析函数，将文件内容解析到切片
func parseFileToList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
