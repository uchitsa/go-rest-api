package models

import (
	"fmt"
	u "innohack-backend/utils"

	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Id          uuid.uuid `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Answers     uint      `json:"answers"`
}

func (c *Course) Create() map[string]interface{} {
	if res, ok := c.Validate(); !ok {
		return res
	}

	GetDB().Create(c)
	res := u.Message(true, "success")
	res["course"] = c
	return res

}

func (c *Course) Validate() (map[string]interface{}, bool) {
	if c.Name == "" {
		return u.Message(false, "Course name must not be empty"), false
	}

	return u.Message(true, "success"), true
}

func GetCourse(id uint) *Course {
	course := &Course{}
	err := GetDB().Table("courses").Where("id = ?", id).First(course).Error
	if err != nil {
		return nil
	}
	return course
}

func GetCourses() []*Course {
	courses := make([]*Course, 0)
	err := GetDB().Table("courses").Find(&courses).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return courses
}
