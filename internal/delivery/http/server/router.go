package server

func (s *Server) createController() {
	s.serv.Use(s.mdlware.CORSMiddleware)

	commonGroup := s.serv.Group("/api")

	publicGroup := commonGroup.Group("")

	publicGroup.POST("/auth/login", s.GetUserToken)
	publicGroup.POST("/auth/registration", s.CreateUser)
	//publicGroup.POST("auth/refresh")

	// authGroup := commonGroup.Group("").Use(s.mdlware.JwtTokenCheck)
	// .Use(s.mdlware.RoleMiddleware) //todo middleware

	//authGroup.GET("/admin", s.Amin)
	//authGroup.GET("/users", s.GetAllUsers)
}
