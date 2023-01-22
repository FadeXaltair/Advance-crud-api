package config

type Request struct {
	Id        int    `json:"id" binding:"required"`
	Country   string `json:"country"`
	Continent string `json:"continent"`
	Currency  string `json:"currency"`
}
