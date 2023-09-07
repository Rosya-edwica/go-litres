package models

type Book struct {
	Id          int
	Name        string
	Description string
	Image       string
	IsAudio     bool
	Url         string
	OldPrice    float64
	NewPrice    float64
	Currency    string
	MinAge      int
	Language    string
	Rating      float32
	Pages       int
	Year        int
}

type JsonResponse struct {
	Payload struct {
		Data struct {
			Id          int     `json:"id"`
			Name        string  `json:"title"`
			Description string  `json:"html_annotation"`
			Image       string  `json:"cover_url"`
			IsAudio     int     `json:"art_type"`
			Url         string  `json:"url"`
			Price struct {
				OldPrice    float64 `json:"full_price"`
				NewPrice    float64 `json:"final_price"`
				Currency    string  `json:"currency"`
			} `json:"prices"`
			
			MinAge      int     `json:"min_age"`
			Language    string  `json:"language_code"`
			Rating struct {
				Code float32 `json:"rated_avg"`
			} `json:"rating"`
			Page struct {
				Count int `json:"current_pages_or_seconds"`
			} `json:"additional_info"`
			Year        string  `json:"date_written_at"`
		} `json:"data"`
	} `json:"payload"`
}
