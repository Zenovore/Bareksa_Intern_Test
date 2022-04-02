package tags

type RequestCreateTag struct {
	Name string `json:"name" validate:"required"`
}

type RequestSearchTag struct {
	GUID string `json:"guid" validate:"omitempty"`
	Name string `json:"name" validate:"omitempty"`
}

type RequestUpdateTag struct {
	GUID string `json:"guid" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type RequestDeleteTag struct {
	GUID string `json:"guid" validate:"required"`
}
