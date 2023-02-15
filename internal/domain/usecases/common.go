package usecases

func (s *St) SetMessageByStatusCode(statusCode int) string {
	output := s.cr.SetMessageByStatusCode(statusCode)

	return output
}
