package models

// TemplateData - holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Float32   map[string]float32
	Data      map[string]interface{}
	CSRFToken string //Crossed Sight Request Forgery Token
	Flash     string // post a (success) message to end user
	Warning   string
	Error     string
}