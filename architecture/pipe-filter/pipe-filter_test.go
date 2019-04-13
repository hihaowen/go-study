package pipe_filter

import "testing"

func TestPipeFilter(t *testing.T) {
	content := "这是一段内容"
	WordsFilterReport(content, FanDongWordsFilter{}, SeQingWordsFilter{}, BaoLiWordsFilter{})
}
