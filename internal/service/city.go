package service

import (
	cities "github.com/Kit0b0y/SkillboxHomeWork/NewSkillbox/Interim_certification"
	"github.com/Kit0b0y/SkillboxHomeWork/NewSkillbox/Interim_certification/internal/repository"
)

type CityService struct {
	repo repository.CityList
}

func NewCityService(repo repository.CityList) *CityService {
	return &CityService{repo: repo}
}


func (s *CityService) Create(city cities.CityRequest) (string, error) {
	return s.repo.Create(city)
}


func (s *CityService) Delete(id int) error {
	return s.repo.Delete(id)
}


func (s *CityService) SetPopulation(id, population int) error {
	return s.repo.SetPopulation(id, population)
}


func (s *CityService) GetFromRegion(region string) ([]string, error) {
	return s.repo.GetFromRegion(region)
}


func (s *CityService) GetFromDistrict(distinct string) ([]string, error) {
	return s.repo.GetFromDistrict(distinct)
}


func (s *CityService) GetFromPopulation(start, end int) ([]string, error) {
	return s.repo.GetFromPopulation(start, end)
}


func (s *CityService) GetFromFoundation(start, end int) ([]string, error) {
	return s.repo.GetFromFoundation(start, end)
}


func (s *CityService) GetFull(id int) (*cities.City, error) {
	return s.repo.GetFull(id)
}