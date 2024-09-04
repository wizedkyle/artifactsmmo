package models

import "time"

type EventResponse struct {
	Data  []ActiveEvent `json:"data"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Pages int           `json:"pages"`
}

type ActiveEvent struct {
	Name         string    `json:"name"`
	Map          MapSchema `json:"map"`
	PreviousSkin string    `json:"previous_skin"`
	Duration     int       `json:"duration"`
	Expiration   time.Time `json:"expiration"`
	CreatedAt    time.Time `json:"created_at"`
}

type MapSchema struct {
	Name    string     `json:"name"`
	Skin    string     `json:"skin"`
	X       int        `json:"x"`
	Y       int        `json:"y"`
	Content MapContent `json:"content"`
}

type MapContent struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type GetAllEventsQueryParameters struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
