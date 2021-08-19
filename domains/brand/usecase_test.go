package brand_test

import (
	"context"
	"errors"
	"os"
	"shoe-store/domains"
	"shoe-store/domains/brand"
	"shoe-store/domains/brand/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	brandRepository mocks.Repository
	brandUsecase    brand.Usecase
)

func setup() {
	brandUsecase = brand.NewBrandUsecase(2, &brandRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := []brand.Domain{
			{
				ID:   1,
				Name: "Adidas",
			},
			{
				ID:   2,
				Name: "Nike",
			},
		}
		brandRepository.On("Find", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := brandUsecase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, len(result), 2)
	})
	t.Run("test case 2, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		brandRepository.On("Find", mock.Anything, mock.AnythingOfType("string")).Return([]brand.Domain{}, errNotFound).Once()
		result, err := brandUsecase.GetAll(context.Background())

		assert.Equal(t, result, []brand.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := brand.Domain{
			ID:   1,
			Name: "Adidas",
		}
		brandRepository.On("FindByID", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := brandUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := brandUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, brand.Domain{})
		assert.Equal(t, err, domains.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		brandRepository.On("FindByID", mock.AnythingOfType("int")).Return(brand.Domain{}, errNotFound).Once()
		result, err := brandUsecase.GetByID(context.Background(), 10)

		assert.Equal(t, result, brand.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := brand.Domain{
			Name: "Adidas",
		}
		brandRepository.On("GetByName", mock.Anything, mock.Anything).Return(brand.Domain{}, nil).Once()
		brandRepository.On("Store", mock.Anything, mock.Anything).Return(domain, nil).Once()

		result, err := brandUsecase.Store(context.Background(), &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, validation name", func(t *testing.T) {
		errNotFound := domains.ErrInternalServer
		domain := brand.Domain{
			Name: "Adidas",
		}
		brandRepository.On("GetByName", mock.Anything, mock.Anything).Return(brand.Domain{}, errNotFound).Once()
		brandRepository.On("Store", mock.Anything, mock.Anything).Return(domain, nil).Once()

		_, err := brandUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, errNotFound)
	})
	t.Run("test case 3, duplicate error", func(t *testing.T) {
		errDuplicate := domains.ErrDuplicateData
		domain := brand.Domain{
			Name: "Adidas",
		}
		brandRepository.On("GetByName", mock.Anything, mock.Anything).Return(domain, nil).Once()
		brandRepository.On("Store", mock.Anything, mock.Anything).Return(domain, nil).Once()

		_, err := brandUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, errDuplicate)
	})
	// t.Run("test case 4, repository error", func(t *testing.T) {
	// 	errServer := errors.New("(Repo) ID Not Found")
	// 	domain := brand.Domain{
	// 		Name: "Adidas",
	// 	}
	// 	brandRepository.On("GetByName", mock.Anything, mock.Anything).Return(brand.Domain{}, nil).Once()
	// 	brandRepository.On("Store", mock.Anything, mock.Anything).Return(brand.Domain{}, errServer).Once()

	// 	_, err := brandUsecase.Store(context.Background(), &domain)

	// 	assert.Equal(t, err, errServer)
	// })
}
