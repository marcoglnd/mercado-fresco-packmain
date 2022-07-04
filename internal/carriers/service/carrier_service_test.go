package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/domain"
	mock "github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := &domain.Carrier{
		Cid:         "CID#1",
		CompanyName: "Some Company",
		Address:     "Rua Sao Paulo 200",
		Telephone:   "30021025",
		LocalityId:  1,
	}
	repositoryMock.EXPECT().FindByCid(ctx, "CID#1").Return(nil, nil)
	repositoryMock.EXPECT().Create(ctx, carrierFake).Return(carrierFake, nil)

	carrier, err := service.Create(ctx, carrierFake)

	assert.Nil(t, err)
	assert.NotNil(t, carrier)
	assert.Equal(t, carrierFake, carrier)
}

func TestCreateFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := &domain.Carrier{
		Cid:         "CID#1",
		CompanyName: "Some Company",
		Address:     "Rua Sao Paulo 200",
		Telephone:   "30021025",
		LocalityId:  1,
	}
	repositoryMock.EXPECT().FindByCid(ctx, "CID#1").Return(nil, nil)
	repositoryMock.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("repo error"))

	carrier, err := service.Create(ctx, carrierFake)

	assert.Nil(t, carrier)
	assert.NotNil(t, err)
}

func TestCreateConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := &domain.Carrier{
		Cid:         "CID#1",
		CompanyName: "Some Company",
		Address:     "Rua Sao Paulo 200",
		Telephone:   "30021025",
		LocalityId:  1,
	}
	repositoryMock.EXPECT().FindByCid(ctx, "CID#1").Return(carrierFake, nil)

	carrier, err := service.Create(ctx, carrierFake)

	assert.Nil(t, carrier)
	assert.NotNil(t, err)
}
