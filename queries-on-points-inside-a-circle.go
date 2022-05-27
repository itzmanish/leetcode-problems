package leetcodepractice

import (
	"log"
)

func QueriesOnPointsInsideACircle() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
	queries := [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}
	answer := countPoints(points, queries)
	log.Print(answer)
}

func countPoints(points [][]int, queries [][]int) []int {
	answer := make([]int, len(queries))
	for i, q := range queries {
		x2, y2, r := q[0], q[1], q[2]
		for _, p := range points {
			x1, y1 := p[0], p[1]
			// distance = sqrt((x1-x2)^2+(y1-y2)^2)
			x3 := x1 - x2
			y3 := (y1 - y2)
			squaredDistance := (x3 * x3) + (y3 * y3)
			if squaredDistance <= r*r {
				answer[i] += 1
			}
		}
	}
	return answer
}
