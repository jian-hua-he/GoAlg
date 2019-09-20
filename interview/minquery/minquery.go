package minquery

import (
	"fmt"
	"sort"
)

const (
	JOIN_STRING = " OR "
)

type ByLen []string

func (s ByLen) Len() int           { return len(s) }
func (s ByLen) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByLen) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

func generateQueries(keywords []string, maxCharLen int) []string {
	result := []string{}

	// Sort it by string length
	sort.Sort(ByLen(keywords))

	for len(keywords) != 0 {
		// Find the greatest one and remove it from array
		greatest := keywords[len(keywords)-1]
		keywords = keywords[:len(keywords)-1]

		// Compare to maxCharLen
		// If it is equal to maxCharLen then add it to result
		if len(greatest) == maxCharLen {
			result = append(result, greatest)
			continue
		}

		// If the length of greatest string add length of join string
		// and it equal to maxCharLen then add it to result
		if (len(greatest) + len(JOIN_STRING)) == maxCharLen {
			result = append(result, greatest)
			continue
		}

		// If it is too big then do nothing
		if (len(greatest) + len(JOIN_STRING)) > maxCharLen {
			continue
		}

		// If it still has space to join
		// Find the string as possible as we can
		joint := ""
		for i := len(keywords) - 1; i >= 0; i -= 1 {
			// Calculate the max every time
			max := maxCharLen - len(greatest) - len(joint)

			// If it equal to max then remove it from array
			// and add it to joint. break the loop
			if len(keywords[i])+len(JOIN_STRING) == max {
				joint = fmt.Sprintf("%s%s", JOIN_STRING, keywords[i])
				keywords = append(keywords[:i], keywords[i+1:]...)
				break
			}

			// If it bigger than max. Just continue
			if len(keywords[i])+len(JOIN_STRING) > max {
				continue
			}

			// If it less than max then remove it from array
			// and add it to joint. Continue to find another query
			joint = fmt.Sprintf("%s%s%s", joint, JOIN_STRING, keywords[i])
			keywords = append(keywords[:i], keywords[i+1:]...)
		}

		// If we find more query then compose query to result
		if joint != "" {
			r := fmt.Sprintf("%s%s", greatest, joint)
			result = append(result, r)
			continue
		}

		// If not. Just put it to result
		result = append(result, greatest)
	}

	return result
}
