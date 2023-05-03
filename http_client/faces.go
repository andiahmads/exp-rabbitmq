package facest

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type FacesRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    FacesList
}

type FacesList struct {
	Count     int    `json:"count"`
	PageCount int    `json:"page_count"`
	Faces     []Face `json:"faces"`
}

type Face struct {
	FacesToken string       `json:"faces_token"`
	FacesID    string       `json:"faces_id"`
	FacesImage []FacesImage `json:"faces_image"`

	CreatedAt time.Time `json:"created_at"`
}

type FacesImage struct {
	ImagesToken string    `json:"images_token"`
	ImagesURL   string    `json:"images_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type FacesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func GetFaces(ctx context.Context, options *FacesListOptions) (*FacesList, error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf())
}
