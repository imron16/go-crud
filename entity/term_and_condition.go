package entity

// create entity
type TermAndCondition struct {
	Id        int64  `gorm:"primary_key;auto_increment" json:"id"`
	Title     string `gorm:"column:title;" json:"title"`
	TitleEn   string `gorm:"column:title_en;" json:"title_en"`
	Content   string `gorm:"column:content;" json:"content"`
	ContentEn string `gorm:"column:content_en;" json:"content_en"`
}
