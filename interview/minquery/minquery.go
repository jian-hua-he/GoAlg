package minquery

import (
	"sort"
)

type ByLen []string

func (s ByLen) Len() int           { return len(s) }
func (s ByLen) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByLen) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

func generateQueries(keywords []string, maxCharLen int) []string {
	result := []string{}
	sort.Sort(ByLen(keywords))

	return result
}
