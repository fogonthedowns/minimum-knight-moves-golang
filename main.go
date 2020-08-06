package main

type Move struct {
	x     int
	y     int
	count int
	next  *Move
}

// implemented so that map works
type Coordinate struct {
	x int
	y int
}

func minKnightMoves(x int, y int) int {
	root := &Move{0, 0, 0, nil}
	ll := New(root)

	seen := make(map[Coordinate]bool, 0)

	// BFS
	for ll.length != 0 {
		v := ll.Dequeue()
		c := Coordinate{x: v.x, y: v.y}
		if _, ok := seen[c]; ok {
			continue
		}
		seen[c] = true

		// fmt.Println("got: (%v %v) expected(%v, %v)",v.x, v.y, x, y)
		if v.x == x && v.y == y {
			return v.count
		}

		moveKnight(v, ll)

	}
	return -1
}

func moveKnight(knight *Move, ll *LinkedList) {
	possibleMoves := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
	for _, value := range possibleMoves {
		var x, y, c int
		x = knight.x + value[0]
		y = knight.y + value[1]
		c = knight.count + 1
		ll.Enqueue(&Move{x, y, c, nil})
	}
}

type LinkedList struct {
	first  *Move
	last   *Move
	length int
}

func New(m *Move) (ll *LinkedList) {
	return &LinkedList{first: m, last: m, length: 1}
}

func (ll *LinkedList) Enqueue(m *Move) {
	if ll.first == nil {
		ll.first = m
		ll.last = m
		ll.length = 1
		return
	}

	ll.last.next = m
	ll.last = m
	ll.length++
}

func (ll *LinkedList) Dequeue() (m *Move) {
	if ll.length == 0 {
		return nil
	}

	m = ll.first
	if ll.length == 1 {
		ll.first = nil
		ll.last = nil
	} else {
		ll.first = ll.first.next
	}

	ll.length--
	return m
}
