package repo

type Repository interface {
	TableName() string
	Columns() []string
	Columns2Query() string
}
