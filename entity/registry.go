package entity

func NewRegistry(rootName string) *Registry {
	return &Registry{
		Root:   rootName,
		Titles: map[string]string{},
		Hashes: map[string]*Branch{},
	}
}

type Registry struct {
	Root   string             `json:"Root" bson:"Root"`     //the root branch that registry binding to
	Titles map[string]string  `json:"Titles" bson:"Titles"` //store the title of all sub branches
	Hashes map[string]*Branch `json:"-"`                    //fast search
}

func (registry *Registry) GetID() string {

	return registry.Root
}
