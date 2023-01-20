package server

import "go-crud/interface/container"

type handler struct {
	termConditionHandler *termConditionHandler
}

func setupHandler(container container.Container) *handler {
	termConditionHandler := newTermConditionHanlder(container.TermConditionService)
	return &handler{termConditionHandler: termConditionHandler}
}
