package domain

type SearchRequest struct {
    Name        string `json:"name"`
    ID          uint   `json:"id"`
    Title       string `json:"title"`
    IsCompleted *bool  `json:"iscompleted"`
    Style       string `json:"style"`
    Language    string `json:"language"`
}
