package routebetweennodes

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	graph    *lib.Graph[int]
	s        *lib.GraphNode[int]
	e        *lib.GraphNode[int]
	expected bool
}

func TestRouteBetweenNodes(t *testing.T) {
	testCases := []TestCase{}

	// linked list
	{
		m1 := lib.GraphNode[int]{Val: 1}
		m2 := lib.GraphNode[int]{Val: 2}
		m3 := lib.GraphNode[int]{Val: 3}
		m1.Children = []*lib.GraphNode[int]{&m2}
		m2.Children = []*lib.GraphNode[int]{&m3}
		m3.Children = []*lib.GraphNode[int]{}
		linkedListGraph := []*lib.GraphNode[int]{&m1, &m2, &m3}
		for i := range 3 {
			for j := 0; j < i; j++ {
				testCases = append(testCases, TestCase{
					graph:    &lib.Graph[int]{Nodes: linkedListGraph},
					s:        linkedListGraph[i],
					e:        linkedListGraph[j],
					expected: false,
				})
			}
			for j := i; j < 3; j++ {
				testCases = append(testCases, TestCase{
					graph:    &lib.Graph[int]{Nodes: linkedListGraph},
					s:        linkedListGraph[i],
					e:        linkedListGraph[j],
					expected: true,
				})
			}
		}
	}

	// circular
	{
		n1 := lib.GraphNode[int]{Val: 1}
		n2 := lib.GraphNode[int]{Val: 2}
		n3 := lib.GraphNode[int]{Val: 3}
		n1.Children = []*lib.GraphNode[int]{&n2}
		n2.Children = []*lib.GraphNode[int]{&n3}
		n3.Children = []*lib.GraphNode[int]{&n1}
		circularGraph := []*lib.GraphNode[int]{&n1, &n2, &n3}
		for i := range 3 {
			for j := range 3 {
				testCases = append(testCases, TestCase{
					graph:    &lib.Graph[int]{Nodes: circularGraph},
					s:        circularGraph[i],
					e:        circularGraph[j],
					expected: true,
				})
			}
		}
	}

	for i, tc := range testCases {
		t.Run(lib.SignedToString(i), func(t *testing.T) {
			t.Logf("RouteBetweenNodes(\n%v, \n%v, \n%v)", tc.graph, tc.s, tc.e)
			got := RouteBetweenNodes(tc.graph, tc.s, tc.e)
			want := tc.expected
			if got != want {
				t.Fatalf("want %t, got %t", want, got)
			} else {
				t.Logf("got %t", got)
			}
		})
	}
}
