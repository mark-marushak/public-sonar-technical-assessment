package repository

type ServiceCompare interface {
	Compare(string) bool
}

type service struct {
	repository InterfaceCompareFunc
}

func NewServiceCompare(repo InterfaceCompareFunc) ServiceCompare {
	return &service{
		repository: repo,
	}
}

func (s service) Compare(pattern string) bool {
	return s.repository.Compare(pattern)
}
