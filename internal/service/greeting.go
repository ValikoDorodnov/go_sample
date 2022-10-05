package service

type GreetingService struct {
}

func NewGreetingService() *GreetingService {
	return &GreetingService{}
}

func (gs *GreetingService) HelloProcess() (string, error) {
	return "Hello", nil
}
