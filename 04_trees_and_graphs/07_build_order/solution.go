package buildorder

import "github.com/NoamBechhofer/CTCI-Go/lib"

type Project string

type ProjectDependency struct {
	depended  Project
	dependent Project
}

// returns false if there is a cycle
func BuildOrder(projects []Project, dependencies []ProjectDependency) ([]Project, bool) {
	dependedToDependents := map[Project][]Project{}

	for _, dep := range dependencies {
		dependedToDependents[dep.depended] = append(dependedToDependents[dep.depended], dep.dependent)
	}

	indegrees := map[Project]uint{}
	for _, proj := range projects {
		indegrees[proj] = 0
	}
	for _, dep := range dependencies {
		indegrees[dep.dependent]++
	}

	queue := lib.ArrayQueue[Project]{}

	for proj, indegree := range indegrees {
		if indegree == 0 {
			queue.Add(proj)
		}
	}

	ret := make([]Project, 0, len(projects))
	for !queue.IsEmpty() {
		proj, _ := queue.Remove()
		ret = append(ret, proj)
		for _, dependency := range dependedToDependents[proj] {
			indegrees[dependency]--
			if indegrees[dependency] == 0 {
				queue.Add(dependency)
			}
		}
	}

	return ret, len(ret) == len(projects)
}
