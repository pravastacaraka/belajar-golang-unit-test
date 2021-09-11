package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/** Studi Kasus
 * Membuat aplikasi golang yang melakukan query ke database
 * 1. Buat kontrak untuk interface repository
 * 2. Buat layer repository untuk koneksi ke database
 * 3. Buat layer service untuk bussiness logic
 **/

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	
	result, err := categoryService.Get("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id: "2",
		Name: "Laptop",
	}

	// program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)
	
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
