package entity

type Provinsi struct {
	IdProp   int    `gorm:"primaryKey;column:idprop;"`
	KdProp   string `gorm:"column:kdprop;"`
	NamaProp string `gorm:"column:nama_prop;"`
	KodeArea string `gorm:"column:kode_area;"`
}

func (Provinsi) TableName() string {
	return "master_provinsi"
}
