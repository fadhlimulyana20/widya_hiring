package params

type QuestionOptionCreate struct {
	Body        string `json:"body"`
	OptionValue bool   `json:"option_value"`
	IsImage     bool   `json:"is_image"`
	ImgPath     string `json:"img_path"`
}

type QuestionOptionAdd struct {
	Body        string `json:"body"`
	OptionValue bool   `json:"option_value"`
	IsImage     bool   `json:"is_image"`
	ImgPath     string `json:"img_path"`
	QuestionID  int    `json:"question_id"`
}

type QuestionOptionUpdate struct {
	ID          int    `json:"id"`
	Body        string `json:"body"`
	OptionValue bool   `json:"option_value"`
	IsImage     bool   `json:"is_image"`
	ImgPath     string `json:"img_path"`
	QuestionID  int    `json:"question_id"`
}
