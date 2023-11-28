package params

type AttemptAnswerQuestionParam struct {
	QuestionID int `json:"question_id" validate:"required"`
	OptionID   int `json:"option_id" validate:"required"`
	UserID     int `json:"user_id" validate:"required"`
}

type AttemptClearAnswerQuestionParam struct {
	QuestionID int `json:"question_id" validate:"required"`
	UserID     int `json:"user_id" validate:"required"`
}

type AttemptGetLatestAnswersParam struct {
	QuestionIDs []int `json:"question_ids" validate:"required"`
	UserID      int   `json:"user_id" validate:"required"`
}

type AttemptSubmitAnswerQuestionParam struct {
	QuestionID int `json:"question_id" validate:"required"`
	UserID     int `json:"user_id" validate:"required"`
}

type AttemptSubmitAnswerResponse struct {
	TrueAnswerID     int  `json:"true_answer_id"`
	AnswerID         int  `json:"answer_id"`
	AttemptValue     bool `json:"attempt_value"`
	TrueAnswerStreak int  `json:"true_answer_streak"`
}
