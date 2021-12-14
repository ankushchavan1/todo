package todos

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"todo/internal/auth/user"
)

type Status string
const (
	Completed Status = "completed"
	Incomplete Status = "incomplete"
	Canceled Status = "canceled"
)

/*
	_Todo model
 */
type Todo struct {
	gorm.Model
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	IsBookmarked bool      `json:"isBookmarked"`
	Owner        user.User    `json:"owner" gorm:"foreignKey:ID"`
	Editors      []user.User  `json:"editors" gorm:"many2many:todo_editors;"`
	Status       Status    `json:"status"`
}