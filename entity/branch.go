package entity

func NewBranch() *Branch {
	return &Branch{
		SubBranches: map[string]*Branch{},
	}
}

type Branch struct {
	SubBranches map[string]*Branch
}

func (branch *Branch) Find(parts []string) *Branch {

	if len(parts) == 0 {

		return branch
	}
	sub, ok := branch.SubBranches[parts[0]]
	if !ok {
		newBranch := NewBranch()
		branch.SubBranches[parts[0]] = newBranch
		sub = newBranch
	}
	return sub.Find(parts[1:])
}
