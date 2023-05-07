package graph

import "github.com/Ibuki-Y/go-gql-git/internal"

func ComplexityConfig() internal.ComplexityRoot {
	var c internal.ComplexityRoot

	c.Repository.Issues = func(childComplexity int, after *string, before *string, first *int, last *int) int {
		var cnt int
		switch {
		case first != nil && last != nil:
			if *first < *last {
				cnt = *last
			} else {
				cnt = *first
			}
		case first != nil && last == nil:
			cnt = *first
		case first == nil && last != nil:
			cnt = *last
		default:
			cnt = 1
		}
		return cnt * childComplexity
	}
	return c
}
