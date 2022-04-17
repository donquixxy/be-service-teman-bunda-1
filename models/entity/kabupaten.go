package entity

type Kabupaten struct {
	IdKabu   int    `gorm:"primaryKey;column:idkabu;"`
	IdProp   int    `gorm:"column:idprop;"`
	KdKabu   string `gorm:"column:kdkabu;"`
	NamaKabu string `gorm:"column:nama_kabu;"`
}

func (Kabupaten) TableName() string {
	return "master_kabupaten"
}
