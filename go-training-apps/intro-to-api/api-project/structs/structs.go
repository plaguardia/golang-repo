package structs

type responce struct {
	code int         `json:"code"` //this end piece serializes as json
	body interface{} `json:"body"`
}
