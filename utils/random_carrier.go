package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/carriers/domain"

func CreateRandomCarrier() domain.Carrier {
	carrier := domain.Carrier{
		ID:          0,
		Cid:         RandomString(4),
		CompanyName: RandomString(10),
		Address:     RandomString(10),
		Telephone:   RandomString(6),
		LocalityId:  1,
	}
	return carrier
}

func CreateRandomCarrierReport() domain.CarrierReport {
	carrierReport := domain.CarrierReport{
		LocalityId:    0,
		LocalityName:  RandomString(6),
		CarriersCount: RandomInt(1, 100),
	}
	return carrierReport
}

func CreateRandomListCarriers() []domain.Carrier {
	var listOfCarriers []domain.Carrier
	for i := 1; i <= 5; i++ {
		carrier := CreateRandomCarrier()
		carrier.ID = int64(i)
		listOfCarriers = append(listOfCarriers, carrier)
	}
	return listOfCarriers
}

func CreateRandomListCarriersReport() []domain.CarrierReport {
	var listOfCarriersReport []domain.CarrierReport
	for i := 1; i <= 5; i++ {
		report := CreateRandomCarrierReport()
		report.LocalityId = int64(i)
		listOfCarriersReport = append(listOfCarriersReport, report)
	}
	return listOfCarriersReport
}
