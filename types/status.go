package types

type Status int

const (
	StatusSuccess Status = iota
	StatusFail
	StatusError
)
