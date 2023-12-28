package queries

type FindAllQuery struct {
	Limit  int64
	Skip   int64
	Filter interface{}
}
