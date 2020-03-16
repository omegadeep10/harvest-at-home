package models

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`               // http response status code
	ErrorText      string `json:"error,omitempty"` // application-level error message, for debugging`
}

type SuccessResponse struct {
	HTTPStatusCode int    `json:"-"`      // http response status code
	StatusText     string `json:"status"` // user-level status message
}

type Item struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	PricePerUnit float64 `json:"price_per_unit"`
	UnitSize     string  `json:"unit_size"`
	Image        string  `json:"image"`
	FarmName     string  `json:"farm_name"`
	ProfilePic   string  `json:"profile_pic"`
}
