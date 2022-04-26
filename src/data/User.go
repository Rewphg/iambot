package data

type LineUserObject struct {
	DisplayName string `json:"displayName" validate:"required"`
	UserID      string `json:"userId"`
	Language    string `json:"language" validate:"required"`
	UserStatus  string `json:"statusMessage" validate:"required"`
	PictureURL  string `json:"pictureUrl" validate:"required"`
}
