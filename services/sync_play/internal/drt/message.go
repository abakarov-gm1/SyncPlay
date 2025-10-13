package drt

type Message struct {
	Type    string  `json:"type"` // "chat" или "action"
	Message string  `json:"message,omitempty"`
	State   string  `json:"state,omitempty"` // "paused"/"playing"
	Time    float64 `json:"time,omitempty"`  // позиция видео
	LocalId string  `json:"localId"`         // уникальный ID клиента
}
