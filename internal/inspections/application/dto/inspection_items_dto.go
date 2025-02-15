package dto

type CreateInspectionItemsRequest struct {
	Question  string   `json:"question"`
	Answer    string   `json:"answer"`
	PhotoURLs []string `json:"photo_urls"`
	Score     int      `json:"score"`
	Comment   string   `json:"comment"`
}

type UpdateInspectionItemsRequest struct {
	Question  string   `json:"question"`
	Answer    string   `json:"answer"`
	PhotoURLs []string `json:"photo_urls"`
	Score     int      `json:"score"`
	Comment   string   `json:"comment"`
}
