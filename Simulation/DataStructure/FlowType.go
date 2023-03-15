package datastructure

type FlowType int

const (
	SlightlyCompressible FlowType = iota + 1
	InCompressible
	Compressible
	MultiphaseFlow
)
