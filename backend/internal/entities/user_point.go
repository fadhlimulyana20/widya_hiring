package entities

type UserPoint struct {
	ID     int  `json:"id" gorm:"primaryKey"`
	UserID int  `json:"user_id"`
	Point  int  `json:"point"`
	User   User `json:"user"`
}
