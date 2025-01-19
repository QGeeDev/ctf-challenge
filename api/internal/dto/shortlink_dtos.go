package dto

type CreateShortlinkRequest struct {
	TargetUrl string `json:"target_url"`
}

type ShortlinkDbObject struct {
	Id         int    `json:"id"`
	Slug       string `json:"slug"`
	FullLink   string `json:"full_link"`
	QRImagesId int    `json:"qr_images_id"`
}
