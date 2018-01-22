package api

type Api interface {
	Run()
}

func New() Api {
	return PingApi{}
}
