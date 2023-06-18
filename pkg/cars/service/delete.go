package services

func (s *carService) CarDelete(id string) error {
	//GET ID FROM THE URL
	//DELETE CAR
	err := s.store.CarDelete(id)
	if err!= nil {
        return err
    }

	//RESPOND WITH SUCCESS
	return nil
}