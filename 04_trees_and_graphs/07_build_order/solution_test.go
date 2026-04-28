package buildorder

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type testCase struct {
	projects     []Project
	dependencies []ProjectDependency
	cyclical     bool
}

func validOrdering(test testCase, got []Project) bool {
	if len(got) != len(test.projects) {
		return false
	}

	dependencies := map[ProjectDependency]struct{}{}
	for _, dep := range test.dependencies {
		dependencies[dep] = struct{}{}
	}

	for _, proj := range got {
		// if some dependency where dependent == proj return false
		for dep := range dependencies {
			if dep.dependent == proj {
				return false
			}
		}
		// remove depencies where depended == proj
		for dep := range dependencies {
			if dep.depended == proj {
				delete(dependencies, dep)
			}
		}
	}

	return true
}

func TestBuildOrder(t *testing.T) {
	var a, b, c, d, e, f Project
	a = "a"
	b = "b"
	c = "c"
	d = "d"
	e = "e"
	f = "f"

	testCases := []testCase{
		{
			projects:     []Project{},
			dependencies: []ProjectDependency{},
			cyclical:     false,
		},
		{
			projects:     []Project{a},
			dependencies: []ProjectDependency{},
			cyclical:     false,
		}, {
			projects:     []Project{a, b},
			dependencies: []ProjectDependency{{depended: a, dependent: b}},
			cyclical:     false,
		},
		{
			projects: []Project{a, b, c},
			dependencies: []ProjectDependency{
				{depended: a, dependent: b},
				{depended: b, dependent: c},
			},
			cyclical: false,
		},
		{
			projects: []Project{a, b, c, d, e, f},
			dependencies: []ProjectDependency{
				{depended: a, dependent: d},
				{depended: f, dependent: b},
				{depended: b, dependent: d},
				{depended: f, dependent: a},
				{depended: d, dependent: c},
			},
			cyclical: false,
		},
		{
			projects: []Project{a, b, c},
			dependencies: []ProjectDependency{
				{depended: a, dependent: b},
				{depended: b, dependent: c},
				{depended: c, dependent: a},
			},
			cyclical: true,
		},
	}

	for i, tc := range testCases {
		t.Run(lib.SignedToString(i), func(t *testing.T) {
			got, ok := BuildOrder(tc.projects, tc.dependencies)
			if ok == tc.cyclical {
				t.Fatalf("expected ok=%t, got %t", !tc.cyclical, ok)
			} else if ok && !validOrdering(tc, got) {
				t.Fatalf("%v is an invalid ordering", got)
			} else if ok {
				t.Logf("got %v", got)
			} else {
				t.Logf("cyclical")
			}
		})
	}
}
