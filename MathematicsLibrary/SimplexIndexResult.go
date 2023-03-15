package mathematicslibrary

type SimplexIndexResult struct {
	Index  []int
	Result SimplexResult
}

func NewSimplexIndexResult(Index []int, Result SimplexResult) SimplexIndexResult {
	return SimplexIndexResult{
		Index:  Index,
		Result: Result,
	}
}
