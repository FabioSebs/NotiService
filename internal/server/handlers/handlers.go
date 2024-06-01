package handlers

type Handlers struct {
	EmailHandler EmailHandler
	//add more
}

func NewHandler() *Handlers {
	return &Handlers{}
}

func (h *Handlers) SetEmailHandler(handler EmailHandler) Handlers {
	h.EmailHandler = handler
	return *h
}
