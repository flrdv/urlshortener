package model

type CreateRedirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type GetRedirect struct {
	From string `json:"from"`
}
