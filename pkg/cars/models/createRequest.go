package models

type CreateCarRequest struct {
	CAR_ID           *uint    `json:"CAR_ID" binding:"required"`
	COMPANY_NAME string  `json:"COMPANY_NAME" binding:"required"`
	CAR_NAME     string  `json:"CAR_NAME" binding:"required"`
	YEAR         int64   `json:"YEAR" binding:"required"`
	PRICE        float64 `json:"PRICE" binding:"required"`
}

func (r CreateCarRequest) ToCar() Car	{
	return Car{
		CAR_ID: 			r.CAR_ID,
        COMPANY_NAME: 	r.COMPANY_NAME,
        CAR_NAME: 		r.CAR_NAME,
        YEAR: 			r.YEAR,
        PRICE: 			r.PRICE,
	}
}