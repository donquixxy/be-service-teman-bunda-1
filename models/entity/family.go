package entity

type Family struct {
	Id          string `gorm:"primaryKey;column:id;"`
	IdProvinsi  int    `gorm:"column:id_provinsi;"`
	IdKabupaten int    `gorm:"column:id_kabupaten;"`
	IdKecamatan int    `gorm:"column:id_kecamatan;"`
	IdKelurahan int    `gorm:"column:id_kelurahan;"`
}

func (Family) TableName() string {
	return "family"
}
