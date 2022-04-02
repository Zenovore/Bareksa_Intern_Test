package news

type RequestSearchNews struct {
	GUID             string `json:"guid" validate:"omitempty"`
	Title            string `json:"title" validate:"omitempty"`
	Tag              string `json:"tag" validate:"omitempty"`
	Topic            string `json:"topic"`
	IsPublishedValid int    `json:"is_published_valid" validate:"omitempty"`
	IsPublished      int    `json:"is_published" validate:"omitempty"`
	IsDeletedValid   int    `json:"is_deleted_valid" validate:"omitempty"`
	IsDeleted        int    `json:"is_deleted" validate:"omitempty"`
}

type RequestFilterNews struct {
	Tags []string `json:"tags" validate:"required"`
}

type RequestUpdateNews struct {
	GUID             string   `json:"guid" validate:"required"`
	Title            string   `json:"title" validate:"omitempty"`
	Content          string   `json:"content" validate:"omitempty"`
	AddedTags        []string `json:"added_tags" validate:"omitempty"`
	DeletedTags      []string `json:"deleted_tags" validate:"omitempty"`
	IsPublishedValid int      `json:"is_published_valid" validate:"omitempty"`
	IsPublished      int      `json:"is_published" validate:"omitempty"`
}

type RequestCreateNews struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags" validate:"omitempty"`
}

type RequestDeleteNews struct {
	GUID string `json:"guid" validate:"required"`
}

type RequestAddTag struct {
	GUIDNews string `json:"guid_news" validate:"required"`
	GUIDTag  string `json:"guid_tag" validate:"required"`
}
