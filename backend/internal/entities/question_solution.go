package entities

import "gitlab.com/project-quiz/internal/entities/base"

type QuestionSolution struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	SolutionImgUrl string `json:"solution_img_url"`
	QuestionID     int    `json:"question_id"`
	SolutionText   string `json:"solution_text"`
	SolutionType   string `json:"solution_type"`
	PdfFileUrl     string `json:"pdf_file_url"`
	Link           string `json:"link"`
	base.Timestamp
}
