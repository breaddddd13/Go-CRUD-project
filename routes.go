package main

func initializeRoutes() {
	router.GET("/students", GetStudents)
	router.GET("/students/:id", GetStudent)
	router.POST("/students", CreateStudent)
	router.PUT("/students/:id", UpdateStudent)
	router.DELETE("/students/:id", DeleteStudent)
}
