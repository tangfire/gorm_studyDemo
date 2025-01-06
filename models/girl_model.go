package models

type GirlModel struct {
	ID      int
	Name    string     `gorm:"size:16"`
	BoyList []BoyModel `gorm:"foreignkey:GirlID"`
}

type BoyModel struct {
	ID        int
	Name      string `gorm:"size:16"`
	GirlID    int
	GirlModel GirlModel `gorm:"foreignkey:GirlID"`
}
