package entity

type FamilyMembers struct {
	Id          string `gorm:"primaryKey;column:id;"`
	IdFamily    string `gorm:"column:id_family;"`
	Family      Family `gorm:"foreignKey:IdFamily"`
	FullName    string `gorm:"column:full_name;"`
	Address     string `gorm:"column:address;"`
	Phone       string `gorm:"column:phone;"`
	Email       string `gorm:"column:email;"`
	IdProvinsi  int    `gorm:"column:id_provinsi;"`
	IdKabupaten int    `gorm:"column:id_kabupaten;"`
	IdKecamatan int    `gorm:"column:id_kecamatan;"`
	IdKelurahan int    `gorm:"column:id_kelurahan;"`
}

func (FamilyMembers) TableName() string {
	return "family_members"
}
