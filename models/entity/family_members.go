package entity

type FamilyMembers struct {
	Id       string `gorm:"primaryKey;column:id;"`
	IdFamily string `gorm:"column:id_family;"`
	Family   Family `gorm:"foreignKey:IdFamily"`
	FullName string `gorm:"column:full_name;"`
	Address  string `gorm:"column:address;"`
	Phone    string `gorm:"column:phone;"`
	Email    string `gorm:"column:email;"`
}

func (FamilyMembers) TableName() string {
	return "family_members"
}
