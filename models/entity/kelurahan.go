package entity

type Kelurahan struct {
	IdKelu            int             `gorm:"primaryKey;column:idkelu;"`
	IdKeca            int             `gorm:"column:idkeca;"`
	IdKabu            int             `gorm:"column:idkabu;"`
	IdProp            int             `gorm:"column:idprop;"`
	KdKelu            string          `gorm:"column:kdkelu;"`
	NamaKelu          string          `gorm:"column:nama_kelu;"`
	IdServiceZonaArea int             `gorm:"column:id_service_zona_area"`
	ServiceZonaArea   ServiceZonaArea `gorm:"foreignKey:IdServiceZonaArea"`
}

func (Kelurahan) TableName() string {
	return "master_kelurahan"
}
