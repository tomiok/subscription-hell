package web

import "github.com/tomiok/subscription-hell/internal/subscription"

type SubHandler struct {
	service *subscription.Service
}

func NewSubHandler(service *subscription.Service) *SubHandler {
	return &SubHandler{service: service}
}
