package main

import (
	"fmt"
	"sort"
)

type obj struct {
	key      int
	category string
	value    int
}

type cateSort struct {
	keys     []obj
	category string
	sumValue int
}

// Struct 구조로 전환
func bindObject(genres []string, plays []int) (mapper map[string]cateSort) {

	mapper = make(map[string]cateSort)

	for i := 0; i < len(genres); i++ {
		key := i
		category := genres[i]
		value := plays[i]

		person := obj{key, category, value}

		var catesortObj cateSort
		var ok bool

		catesortObj, ok = mapper[category]
		if ok {
			catesortObj.category = category
			catesortObj.sumValue = catesortObj.sumValue + value
			catesortObj.keys = append(catesortObj.keys, person)

		} else {
			catesortObj = cateSort{category: category, keys: []obj{person}, sumValue: value}
		}

		mapper[category] = catesortObj

	}

	return

}

//CateGori sturuct to slice Sorted Value
func structToSliceSotredValue(categoriMap map[string]cateSort) (sortCateArr []cateSort) {
	sortCateArr = []cateSort{}
	for _, values := range categoriMap {

		var valueKeyArray []obj = values.keys
		sort.Slice(valueKeyArray, func(i, j int) bool {
			return valueKeyArray[i].value > valueKeyArray[j].value
		})
		values.keys = valueKeyArray
		sortCateArr = append(sortCateArr, values)
	}

	return

}

func solution(genres []string, plays []int) []int {

	var categoriMap map[string]cateSort
	var sortCateArr = []cateSort{}
	var result = []int{}

	//1. map 형식으로 변환 해줌
	categoriMap = bindObject(genres, plays)

	//2. Slice 형식으로 오브젝트 맵핑과 동시에 하위 레벨 정렬 수행
	sortCateArr = structToSliceSotredValue(categoriMap)

	//대분류 소팅
	sort.Slice(sortCateArr, func(i, j int) bool {
		var sorted bool = false

		if sortCateArr[i].sumValue > sortCateArr[j].sumValue {
			sorted = true
		} else if sortCateArr[i].sumValue == sortCateArr[j].sumValue && sortCateArr[i].category < sortCateArr[j].category {
			sorted = true
		}

		return sorted
	})

	// 2개 씩 빼서 출력
	for _, valArr := range sortCateArr {
		for i := 0; i < len(valArr.keys); i++ {
			if i == 2 {
				break
			} else {
				result = append(result, valArr.keys[i].key)
			}

		}
	}

	return result
}

func main() {

	fmt.Println("soulution: ", solution([]string{"classic", "pop", "classic", "classic", "pop"}, []int{500, 600, 150, 800, 2500}))
}
