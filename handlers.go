package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetStudents :Get all students
func GetStudents(c *gin.Context) {
	var ss []Student
	if err := db.Find(&ss).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ss)
	}
}

// GetStudent :get one student by id
func GetStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var s Student
	if err := db.Where("id = ?", id).First(&s).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, s)
	}
}

// CreateStudent :create new student
func CreateStudent(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	major := c.PostForm("major")
	school := c.PostForm("school")
	uid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	s := Student{
		ID:     uid,
		Name:   name,
		Major:  major,
		School: school,
	}

	c.BindJSON(&s)
	db.Create(&s)
	c.JSON(200, s)
}

// UpdateStudent :update student info
func UpdateStudent(c *gin.Context) {
	name := c.PostForm("name")
	major := c.PostForm("major")
	school := c.PostForm("school")
	var s Student

	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&s).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	if name != "" {
		s.Name = name
	}
	if major != "" {
		s.Major = major
	}
	if school != "" {
		s.School = school
	}

	c.BindJSON(&s)
	db.Save(&s)
	c.JSON(200, s)
}

// DeleteStudent :delete student
func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var s Student
	d := db.Where("id = ?", id).Delete(&s)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
