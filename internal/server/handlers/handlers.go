package handlers

type Handlers struct {
	EmailHandler EmailHandler
	KafkaHandler KafkaHandler
	//add more
}

func NewHandler() *Handlers {
	return &Handlers{}
}

func (h *Handlers) SetEmailHandler(handler EmailHandler) Handlers {
	h.EmailHandler = handler
	return *h
}

func (h *Handlers) SetKafkaHandler(handler KafkaHandler) Handlers {
	h.KafkaHandler = handler
	return *h
}
