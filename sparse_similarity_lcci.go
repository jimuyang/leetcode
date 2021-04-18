package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func computeSimilarities(docs [][]int) []string {
	for _, doc := range docs {
		sort.Slice(doc, func(i, j int) bool {
			return doc[i] < doc[j]
		})
		//sort.IntSlice(doc).Sort()
	}
	res := make([]string, 0)
	for i := 0; i < len(docs); i++ {

		for j := i + 1; j < len(docs); j++ {
			interNum := getInter(docs[i], docs[j])
			unionNum := len(docs[i]) + len(docs[j]) - interNum
			if interNum == 0 {
				continue
			}
			val := math.Floor(float64(interNum)/float64(unionNum)*10000+0.5) / float64(10000)
			res = append(res, fmt.Sprintf("%d,%d: %.4f", i, j, val))
		}
	}

	return res
}

func computeSimilarities2(docs [][]int) []string {
	occur := make(map[int][]int)

	for id, doc := range docs {
		for _, w := range doc {
			occur[w] = append(occur[w], id)
		}
	}

	res := make(map[string]int)
	for _, ids := range occur {
		for i := 0; i < len(ids); i++ {
			for j := i + 1; j < len(ids); j++ {
				id1, id2 := ids[i], ids[j]
				if id1 > id2 {
					id1, id2 = id2, id1
				}
				res[fmt.Sprintf("%d,%d", id1, id2)]++
			}
		}
	}

	result := make([]string, 0, len(res))
	for idPair, interNum := range res {
		s := strings.Split(idPair, ",")
		id1, id2 := s[0], s[1]
		i1, _ := strconv.Atoi(id1)
		i2, _ := strconv.Atoi(id2)

		unionNum := len(docs[i1]) + len(docs[i2]) - interNum
		val := math.Floor(float64(interNum)/float64(unionNum)*10000+0.5) / float64(10000)
		result = append(result, fmt.Sprintf("%d,%d: %.4f", i1, i2, val))
	}

	return result
}

func getInter(doc1 []int, doc2 []int) int {
	res := 0
	for i, j := 0, 0; i < len(doc1) && j < len(doc2); {
		if doc1[i] == doc2[j] {
			i++
			j++
			res++
		} else if doc1[i] < doc2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
