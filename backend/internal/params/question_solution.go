package params

import "mime/multipart"

type QuestionSolutionCreate struct {
	SolutionImgUrl   string `json:"solution_img_url"`
	QuestionID       int    `json:"question_id" validate:"required"`
	SolutionText     string `json:"solution_text"`
	SolutionType     string `json:"solution_type" validate:"required"`
	PdfFileUrl       string `json:"pdf_file_url"`
	SolutionImageUrl string `json:"solution_image_url"`
	Link             string `json:"link"`
}

type QuestionSolutionWithFileUploadCreate struct {
	QuestionID   int                   `json:"question_id" validate:"required"`
	SolutionText string                `json:"solution_text"`
	SolutionType string                `json:"solution_type" validate:"required"`
	PdfFile      *multipart.FileHeader `json:"pdf_file"`
	SolutionImg  *multipart.FileHeader `json:"solution_img"`
}

type QuestionSolutionUpdate struct {
	ID               int    `json:"id" validate:"required"`
	SolutionImgUrl   string `json:"solution_img_url"`
	QuestionID       int    `json:"question_id" validate:"required"`
	SolutionText     string `json:"solution_text"`
	SolutionType     string `json:"solution_type" validate:"required"`
	PdfFileUrl       string `json:"pdf_file_url"`
	SolutionImageUrl string `json:"solution_image_url"`
	Link             string `json:"link"`
}
