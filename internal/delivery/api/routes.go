package api

func (s *server) Routes()  {
	s.Router.POST("/create_user", s.createUser)
}