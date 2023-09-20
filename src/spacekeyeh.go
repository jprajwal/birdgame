package main

type SpaceKeyEventHandler struct {
	listeners []SpaceKeyListenable
}

func NewSpaceKeyEventHandler() *SpaceKeyEventHandler {
	return &SpaceKeyEventHandler{listeners: make([]SpaceKeyListenable, 0)}
}

func (eh *SpaceKeyEventHandler) AddListener(listener SpaceKeyListenable) {
	eh.listeners = append(eh.listeners, listener)
}

func (eh *SpaceKeyEventHandler) Notify() {
	for _, listener := range eh.listeners {
		listener.OnSpaceKeyPressed()
	}
}

