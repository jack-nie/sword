package algo

type minStack struct {
	minStack  []int
	dataStack []int
}

func (m *minStack) min() int {
	return m.minStack[len(m.minStack)-1]

}

func (m *minStack) push(i int) {
	m.dataStack = append(m.dataStack, i)
	if len(m.minStack) == 0 || i < m.min() {
		m.minStack = append(m.minStack, i)
	} else {
		m.minStack = append(m.minStack, m.min())
	}
}

func (m *minStack) pop() {
	m.minStack = m.minStack[:len(m.minStack)-1]
	m.dataStack = m.dataStack[:len(m.dataStack)-1]
}
