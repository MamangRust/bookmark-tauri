package mapper

type Mapper struct {
	PostMapper     PostMapping
	CategoryMapper CategoryMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		PostMapper:     NewPostMapper(),
		CategoryMapper: NewCategoryMapper(),
	}
}
