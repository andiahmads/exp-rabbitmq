package service

import "design-pattern/cron_design/repository"

type Service interface {
	InserData() error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) InserData() error {
	data := repository.InputDataGame{}
	data.Mdate = "27-10-2022"
	data.Stadium = "kanjuruhan"
	data.Team1 = "persib"
	data.Team2 = "persib2"

	_, err := s.repository.InsertData(data)
	if err != nil {
		return err
	}

	return nil

}
