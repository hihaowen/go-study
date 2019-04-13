package pipe_filter

import "fmt"

type WordsFilter interface {
	Analyse(words string)
}

// 反动Filter
type FanDongWordsFilter struct {
}

func (f FanDongWordsFilter) Analyse(words string) {
	fmt.Println("内容包含反动内容")
}

// 色情Filter
type SeQingWordsFilter struct {
}

func (s SeQingWordsFilter) Analyse(words string) {
	fmt.Println("内容包含色情内容")
}

// 暴力Filter
type BaoLiWordsFilter struct {
}

func (b BaoLiWordsFilter) Analyse(words string) {
	fmt.Println("内容包含暴力内容")
}

func WordsFilterReport(content string, filters ... WordsFilter) {
	for _, filter := range filters {
		filter.Analyse(content)
	}
}
