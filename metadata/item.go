package metadata

type Transformer func(string) (string, error)

type Item struct {
	Path        string
	Transformer Transformer
}

func Identity(x string) (string, error) {
	return x, nil
}

func NewItem(path string) *Item {
	return &Item{
		Path:        path,
		Transformer: Identity,
	}
}

func NewItemWithTransformer(path string, transformer Transformer) *Item {
	return &Item{
		Path:        path,
		Transformer: transformer,
	}
}
