package stackorque

//TopSolution is so
func TopSolution() {

	arr := [][]int{

		{6, 9, 5, 7, 4},
		{3, 9, 9, 3, 5, 7, 2},
		{1, 5, 3, 6, 7, 6, 5},
		{5, 3, 1, 2, 3},
	}

	//}

	solution(arr[3])
}

type node struct {
	value int
}

type stack struct {
	nodes []node
}

func (s *stack) getStack() *[]node {
	return &s.nodes
}

func (s *stack) new(nodes []int) *stack {

	for val := range nodes {
		node := node{value: nodes[val]}
		s.nodes = append(s.nodes, node)
	}

	return s

}

func (s *stack) pop() int {
	len := len(s.nodes)
	val := s.nodes[len-1].value
	//fmt.Println("POP")
	if len != 0 {
		s.nodes = s.nodes[:len-1]
	}

	return val

}

func solution(heights []int) []int {

	solutionArr := make([]int, len(heights), len(heights))
	inputIndex := len(heights) - 1

	for i := len(heights) - 1; i > 0; i-- {
		var sampleStack stack
		sampleStack.new(heights[:i])

		remainNode := *sampleStack.getStack()

		//fmt.Println("remainNode", remainNode)

		for j := len(remainNode) - 1; j > -1; j-- {

			if sampleStack.pop() > heights[i] {
				//fmt.Println("this", i-j)
				solutionArr[inputIndex] = j + 1
				break
			}

		}
		inputIndex--

	}
	//fmt.Println(solutionArr)

	return solutionArr
}
