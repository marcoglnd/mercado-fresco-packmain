package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/service"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := utils.CreateRandomCarrier()
	repositoryMock.EXPECT().FindByCid(ctx, carrierFake.Cid).Return(nil, nil)
	repositoryMock.EXPECT().Create(ctx, &carrierFake).Return(&carrierFake, nil)

	carrier, err := service.Create(ctx, &carrierFake)

	assert.Nil(t, err)
	assert.NotNil(t, carrier)
	assert.Equal(t, &carrierFake, carrier)
}

func TestCreateFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := utils.CreateRandomCarrier()
	repositoryMock.EXPECT().FindByCid(ctx, carrierFake.Cid).Return(nil, nil)
	repositoryMock.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("repo error"))

	carrier, err := service.Create(ctx, &carrierFake)

	assert.Nil(t, carrier)
	assert.NotNil(t, err)
}

func TestCreateConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := utils.CreateRandomCarrier()
	repositoryMock.EXPECT().FindByCid(ctx, carrierFake.Cid).Return(&carrierFake, nil)

	carrier, err := service.Create(ctx, &carrierFake)

	assert.Nil(t, carrier)
	assert.NotNil(t, err)
}

func TestCreateConflictFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := utils.CreateRandomCarrier()
	repositoryMock.EXPECT().FindByCid(ctx, carrierFake.Cid).Return(nil, errors.New("error"))

	carrier, err := service.Create(ctx, &carrierFake)

	assert.Nil(t, carrier)
	assert.NotNil(t, err)
}

func TestReportAllOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carriersReportFake := utils.CreateRandomListCarriersReport()
	repositoryMock.EXPECT().GetAllCarriersReport(ctx).Return(&carriersReportFake, nil)

	reports, err := service.GetAllCarriersReport(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, reports)
	assert.Equal(t, len(*reports), len(carriersReportFake))
}

func TestReportAllFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()

	repositoryMock.EXPECT().GetAllCarriersReport(ctx).Return(nil, errors.New("some error"))

	reports, err := service.GetAllCarriersReport(ctx)

	assert.Error(t, err)
	assert.NotNil(t, err)
	assert.Nil(t, reports)
}

func TestReportByIdOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carriersReportFake := utils.CreateRandomCarrierReport()
	repositoryMock.EXPECT().GetCarriersReportById(ctx, carriersReportFake.LocalityId).Return(&carriersReportFake, nil)

	reports, err := service.GetCarriersReportById(ctx, carriersReportFake.LocalityId)

	assert.Nil(t, err)
	assert.NotNil(t, reports)
}

func TestReportByIdFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	repositoryMock.EXPECT().GetCarriersReportById(ctx, gomock.Any()).Return(nil, errors.New("repo"))

	reports, err := service.GetCarriersReportById(ctx, int64(1))

	assert.NotNil(t, err)
	assert.Nil(t, reports)
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := utils.CreateRandomCarrier()
	repositoryMock.EXPECT().FindById(ctx, carrierFake.ID).Return(&carrierFake, nil)

	carrier, err := service.FindById(ctx, carrierFake.ID)

	assert.Nil(t, err)
	assert.NotNil(t, carrier)
}

func TestFindByIdFailDb(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	repositoryMock.EXPECT().FindById(ctx, gomock.Any()).Return(nil, errors.New("error"))

	carrier, err := service.FindById(ctx, int64(1))

	assert.NotNil(t, err)
	assert.Nil(t, carrier)
}

func TestFindByIdFailNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	repositoryMock.EXPECT().FindById(ctx, gomock.Any()).Return(nil, nil)

	carrier, err := service.FindById(ctx, int64(1))

	assert.NotNil(t, err)
	assert.Nil(t, carrier)
}

func TestFindByCid(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	carrierFake := utils.CreateRandomCarrier()
	repositoryMock.EXPECT().FindByCid(ctx, carrierFake.Cid).Return(&carrierFake, nil)

	carrier, err := service.FindByCid(ctx, carrierFake.Cid)

	assert.Nil(t, err)
	assert.NotNil(t, carrier)
}

func TestFindByCidFailDb(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockCarrierRepository(ctrl)
	service := service.NewCarrierService(repositoryMock)
	ctx := context.TODO()
	repositoryMock.EXPECT().FindByCid(ctx, gomock.Any()).Return(nil, errors.New("error"))

	carrier, err := service.FindByCid(ctx, "")

	assert.NotNil(t, err)
	assert.Nil(t, carrier)
}
