package services

import (
	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/model"
	repositoriesInterfaces "github.com/nuttchai/go-rest/internal/repositories/interfaces"
	"github.com/nuttchai/go-rest/internal/services/interfaces"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type TSampleService struct {
	repository repositoriesInterfaces.ISampleRepository
}

var (
	SampleService interfaces.ISampleService
)

func (s *TSampleService) Test() string {
	return s.repository.Test()
}

func (s *TSampleService) GetSample(id string) (*model.Sample, error) {
	return s.repository.RetrieveOne(id)
}

func (s *TSampleService) CreateSample(sample *dto.CreateSampleDTO) (*model.Sample, error) {
	return s.repository.CreateOne(sample)
}

func (s *TSampleService) UpdateSample(sample *dto.UpdateSampleDTO) (*model.Sample, error) {
	return s.repository.UpdateOne(sample)
}

func (s *TSampleService) DeleteSample(id string) error {
	result, err := s.repository.DeleteOne(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
