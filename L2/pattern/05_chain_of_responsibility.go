package pattern

type Handler interface {
	SendRequest(message int) string
}

type SecuritySystemHandler struct {
	next Handler
}

func (h *SecuritySystemHandler) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Система безопасности активирована"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type NavigationSystemHandler struct {
	next Handler
}

func (h *NavigationSystemHandler) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Система навигации запущена"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type ParkingAssistHandler struct {
	next Handler
}

func (h *ParkingAssistHandler) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Система помощи при парковке активирована"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
