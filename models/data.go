package models

type ApiData struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       []UserData  `json:"data"`
	Support    SupportInfo `json:"support"`
}

type UserData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type SupportInfo struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}
