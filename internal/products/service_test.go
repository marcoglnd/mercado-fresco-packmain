package products_test

import (
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/products"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomProduct() (product Product) {
	product = Product{
		Id:                             1,
		Description:                    utils.RandomCategory(),
		ExpirationRate:                 utils.RandomCode(),
		FreezingRate:                   utils.RandomCode(),
		Height:                         float64(utils.RandomInt(1, 100)),
		Length:                         float64(utils.RandomInt(1, 100)),
		NetWeight:                      float64(utils.RandomInt(1, 100)),
		ProductCode:                    utils.RandomCategory(),
		RecommendedFreezingTemperature: float64(utils.RandomInt(1, 100)),
		Width:                          float64(utils.RandomInt(1, 100)),
		ProductTypeId:                  utils.RandomCode(),
		SellerId:                       utils.RandomCode(),
	}
	return
}

func createRandomListProduct() (listOfProducts []Product) {
	for i := 1; i <= 5; i++ {
		product := createRandomProduct()
		product.Id = i
		listOfProducts = append(listOfProducts, product)
	}
	return
}

func TestGetAll(t *testing.T) {
	mock := new(mocks.Repository)

	productsArg := createRandomListProduct()

	t.Run("GetAll in case of success", func(t *testing.T) {
		mock.On("GetAll").Return(productsArg, nil).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, list)

		for i := 0; i < len(productsArg); i++ {
			assert.Equal(t, productsArg[i].Id, list[i].Id)
			assert.Equal(t, productsArg[i].Description, list[i].Description)
			assert.Equal(t, productsArg[i].ExpirationRate, list[i].ExpirationRate)
			assert.Equal(t, productsArg[i].FreezingRate, list[i].FreezingRate)
			assert.Equal(t, productsArg[i].Height, list[i].Height)
			assert.Equal(t, productsArg[i].Length, list[i].Length)
			assert.Equal(t, productsArg[i].NetWeight, list[i].NetWeight)
			assert.Equal(t, productsArg[i].ProductCode, list[i].ProductCode)
			assert.Equal(t, productsArg[i].RecommendedFreezingTemperature, list[i].RecommendedFreezingTemperature)
			assert.Equal(t, productsArg[i].Width, list[i].Width)
			assert.Equal(t, productsArg[i].ProductTypeId, list[i].ProductTypeId)
			assert.Equal(t, productsArg[i].SellerId, list[i].SellerId)
		}

		mock.AssertExpectations(t)
	})

	t.Run("GetAll in case of error", func(t *testing.T) {
		mock.On("GetAll").Return(nil, errors.New("failed to retrieve products")).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.Error(t, err)
		assert.Empty(t, list)

		mock.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mock := new(mocks.Repository)

	productArg := createRandomProduct()

	t.Run("GetById in case of success", func(t *testing.T) {
		mock.On("GetById", productArg.Id).Return(productArg, nil).Once()

		service := NewService(mock)

		product, err := service.GetById(productArg.Id)

		assert.NoError(t, err)
		assert.NotEmpty(t, product)

		assert.Equal(t, productArg.Id, product.Id)
		assert.Equal(t, productArg.Description, product.Description)
		assert.Equal(t, productArg.ExpirationRate, product.ExpirationRate)
		assert.Equal(t, productArg.FreezingRate, product.FreezingRate)
		assert.Equal(t, productArg.Height, product.Height)
		assert.Equal(t, productArg.Length, product.Length)
		assert.Equal(t, productArg.NetWeight, product.NetWeight)
		assert.Equal(t, productArg.ProductCode, product.ProductCode)
		assert.Equal(t, productArg.RecommendedFreezingTemperature, product.RecommendedFreezingTemperature)
		assert.Equal(t, productArg.Width, product.Width)
		assert.Equal(t, productArg.ProductTypeId, product.ProductTypeId)
		assert.Equal(t, productArg.SellerId, product.SellerId)

		mock.AssertExpectations(t)

	})

	t.Run("GetById in case of error", func(t *testing.T) {
		mock.On("GetById", 185).Return(Product{}, errors.New("failed to retrieve product")).Once()

		service := NewService(mock)

		product, err := service.GetById(185)

		assert.Error(t, err)
		assert.Empty(t, product)

		mock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Verify product's ID increases when a new product is created", func(t *testing.T) {

		productsArg := createRandomListProduct()

		for _, product := range productsArg {
			mock.On("CreateNewProduct",
				product.Description,
				product.ExpirationRate,
				product.FreezingRate,
				product.Height,
				product.Length,
				product.NetWeight,
				product.ProductCode,
				product.RecommendedFreezingTemperature,
				product.Width,
				product.ProductTypeId,
				product.SellerId).Return(product, nil).Once()
		}

		service := NewService(mock)

		var list []Product

		for _, productArg := range productsArg {
			newProduct, err := service.CreateNewProduct(
				productArg.Description,
				productArg.ExpirationRate,
				productArg.FreezingRate,
				productArg.Height,
				productArg.Length,
				productArg.NetWeight,
				productArg.ProductCode,
				productArg.RecommendedFreezingTemperature,
				productArg.Width,
				productArg.ProductTypeId,
				productArg.SellerId)

			assert.NoError(t, err)
			assert.NotEmpty(t, newProduct)

			assert.Equal(t, productArg.Id, newProduct.Id)
			assert.Equal(t, productArg.Description, newProduct.Description)
			assert.Equal(t, productArg.ExpirationRate, newProduct.ExpirationRate)
			assert.Equal(t, productArg.FreezingRate, newProduct.FreezingRate)
			assert.Equal(t, productArg.Height, newProduct.Height)
			assert.Equal(t, productArg.Length, newProduct.Length)
			assert.Equal(t, productArg.NetWeight, newProduct.NetWeight)
			assert.Equal(t, productArg.ProductCode, newProduct.ProductCode)
			assert.Equal(t, productArg.RecommendedFreezingTemperature, newProduct.RecommendedFreezingTemperature)
			assert.Equal(t, productArg.Width, newProduct.Width)
			assert.Equal(t, productArg.ProductTypeId, newProduct.ProductTypeId)
			assert.Equal(t, productArg.SellerId, newProduct.SellerId)
			list = append(list, newProduct)
		}
		assert.True(t, list[0].Id == list[1].Id-1)

		mock.AssertExpectations(t)
	})

	t.Run("Verify when a ProductCode`s product already exists thrown an error", func(t *testing.T) {
		product1 := createRandomProduct()
		product2 := createRandomProduct()

		product2.ProductCode = product1.ProductCode

		expectedError := errors.New("cid already used")

		mock.On("CreateNewProduct",
			product1.Description,
			product1.ExpirationRate,
			product1.FreezingRate,
			product1.Height,
			product1.Length,
			product1.NetWeight,
			product1.ProductCode,
			product1.RecommendedFreezingTemperature,
			product1.Width,
			product1.ProductTypeId,
			product1.SellerId,
		).Return(product1, nil).Once()
		mock.On("CreateNewProduct",
			product2.Description,
			product2.ExpirationRate,
			product2.FreezingRate,
			product2.Height,
			product2.Length,
			product2.NetWeight,
			product2.ProductCode,
			product2.RecommendedFreezingTemperature,
			product2.Width,
			product2.ProductTypeId,
			product2.SellerId,
		).Return(Product{}, expectedError).Once()

		s := NewService(mock)
		newProduct1, err := s.CreateNewProduct(product1.Description,
			product1.ExpirationRate,
			product1.FreezingRate,
			product1.Height,
			product1.Length,
			product1.NetWeight,
			product1.ProductCode,
			product1.RecommendedFreezingTemperature,
			product1.Width,
			product1.ProductTypeId,
			product1.SellerId)
		assert.NoError(t, err)
		assert.NotEmpty(t, newProduct1)

		assert.Equal(t, product1, newProduct1)

		newProduct2, err := s.CreateNewProduct(
			product2.Description,
			product2.ExpirationRate,
			product2.FreezingRate,
			product2.Height,
			product2.Length,
			product2.NetWeight,
			product2.ProductCode,
			product2.RecommendedFreezingTemperature,
			product2.Width,
			product2.ProductTypeId,
			product2.SellerId)
		assert.Error(t, expectedError, err)
		assert.Empty(t, newProduct2)

		assert.NotEqual(t, product2, newProduct2)
		mock.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Update data in case of success", func(t *testing.T) {
		product1 := createRandomProduct()
		product2 := createRandomProduct()

		product2.Id = product1.Id

		mock.On("CreateNewProduct",
			product1.Description,
			product1.ExpirationRate,
			product1.FreezingRate,
			product1.Height,
			product1.Length,
			product1.NetWeight,
			product1.ProductCode,
			product1.RecommendedFreezingTemperature,
			product1.Width,
			product1.ProductTypeId,
			product1.SellerId,
		).Return(product1, nil).Once()
		mock.On("Update",
			product1.Id,
			product2.Description,
			product2.ExpirationRate,
			product2.FreezingRate,
			product2.Height,
			product2.Length,
			product2.NetWeight,
			product2.ProductCode,
			product2.RecommendedFreezingTemperature,
			product2.Width,
			product2.ProductTypeId,
			product2.SellerId,
		).Return(product2, nil).Once()

		s := NewService(mock)
		newProduct1, err := s.CreateNewProduct(
			product1.Description,
			product1.ExpirationRate,
			product1.FreezingRate,
			product1.Height,
			product1.Length,
			product1.NetWeight,
			product1.ProductCode,
			product1.RecommendedFreezingTemperature,
			product1.Width,
			product1.ProductTypeId,
			product1.SellerId,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, newProduct1)

		assert.Equal(t, product1, newProduct1)

		newproduct2, err := s.Update(
			product1.Id,
			product2.Description,
			product2.ExpirationRate,
			product2.FreezingRate,
			product2.Height,
			product2.Length,
			product2.NetWeight,
			product2.ProductCode,
			product2.RecommendedFreezingTemperature,
			product2.Width,
			product2.ProductTypeId,
			product2.SellerId,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, newproduct2)

		assert.Equal(t, product1.Id, newproduct2.Id)
		assert.NotEqual(t, product1.Description, newproduct2.Description)
		assert.NotEqual(t, product1.ExpirationRate, newproduct2.ExpirationRate)
		assert.NotEqual(t, product1.FreezingRate, newproduct2.FreezingRate)
		assert.NotEqual(t, product1.Height, newproduct2.Height)
		assert.NotEqual(t, product1.Length, newproduct2.Length)
		assert.NotEqual(t, product1.NetWeight, newproduct2.NetWeight)
		assert.NotEqual(t, product1.ProductCode, newproduct2.ProductCode)
		assert.NotEqual(t, product1.RecommendedFreezingTemperature, newproduct2.RecommendedFreezingTemperature)
		assert.NotEqual(t, product1.Width, newproduct2.Width)
		assert.NotEqual(t, product1.ProductTypeId, newproduct2.ProductTypeId)
		assert.NotEqual(t, product1.SellerId, newproduct2.SellerId)

		mock.AssertExpectations(t)
	})

	t.Run("Update throw an error in case of an nonexistent ID", func(t *testing.T) {
		product := createRandomProduct()

		mock.On("Update",
			product.Id,
			product.Description,
			product.ExpirationRate,
			product.FreezingRate,
			product.Height,
			product.Length,
			product.NetWeight,
			product.ProductCode,
			product.RecommendedFreezingTemperature,
			product.Width,
			product.ProductTypeId,
			product.SellerId,
		).Return(Product{}, errors.New("failed to retrieve product")).Once()

		service := NewService(mock)

		product, err := service.Update(
			product.Id,
			product.Description,
			product.ExpirationRate,
			product.FreezingRate,
			product.Height,
			product.Length,
			product.NetWeight,
			product.ProductCode,
			product.RecommendedFreezingTemperature,
			product.Width,
			product.ProductTypeId,
			product.SellerId,
		)

		assert.Error(t, err)
		assert.Empty(t, product)

		mock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mock := new(mocks.Repository)

	productArg := createRandomProduct()

	t.Run("Delete in case of success", func(t *testing.T) {
		mock.On("CreateNewProduct",
			productArg.Description,
			productArg.ExpirationRate,
			productArg.FreezingRate,
			productArg.Height,
			productArg.Length,
			productArg.NetWeight,
			productArg.ProductCode,
			productArg.RecommendedFreezingTemperature,
			productArg.Width,
			productArg.ProductTypeId,
			productArg.SellerId,
		).Return(productArg, nil).Once()
		mock.On("GetAll").Return([]Product{productArg}, nil).Once()
		mock.On("Delete", productArg.Id).Return(nil).Once()
		mock.On("GetAll").Return([]Product{}, nil).Once()

		service := NewService(mock)

		newproduct, err := service.CreateNewProduct(
			productArg.Description,
			productArg.ExpirationRate,
			productArg.FreezingRate,
			productArg.Height,
			productArg.Length,
			productArg.NetWeight,
			productArg.ProductCode,
			productArg.RecommendedFreezingTemperature,
			productArg.Width,
			productArg.ProductTypeId,
			productArg.SellerId,
		)
		assert.NoError(t, err)
		list1, err := service.GetAll()
		assert.NoError(t, err)
		err = service.Delete(newproduct.Id)
		assert.NoError(t, err)
		list2, err := service.GetAll()
		assert.NoError(t, err)

		assert.NotEmpty(t, list1)
		assert.NotEqual(t, list1, list2)
		assert.Empty(t, list2)

		mock.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mock.On("Delete", 185).Return(errors.New("product's ID not founded")).Once()

		service := NewService(mock)

		err := service.Delete(185)

		assert.Error(t, err)

		mock.AssertExpectations(t)
	})
}
