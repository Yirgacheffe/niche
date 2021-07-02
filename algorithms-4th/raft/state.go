package main

type State int

const (
	Follower State = iota
	Candidate
	Leader
	Terminate
)
