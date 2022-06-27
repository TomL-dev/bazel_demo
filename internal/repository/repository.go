package repository

type Instance struct {
}

func New() *Instance {
	i := new(Instance)
	return i
}
