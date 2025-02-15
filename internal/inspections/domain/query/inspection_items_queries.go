package query

type GetInspectionItemByIdQuery struct {
	ID string
}

type ListInspectionItemsQuery struct {
	Limit  int
	Offset int
}
