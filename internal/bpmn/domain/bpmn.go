package domain

type BPMNWorkflow struct {
	ID     string
	Name   string
	Steps  []BPMNStep
	Actors []BPMNActor
}

type BPMNStep struct {
	ID   string
	Name string
}

type BPMNActor struct {
	ID   string
	Name string
}
