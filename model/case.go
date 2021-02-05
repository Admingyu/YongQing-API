package model

type Category struct {
	BasicModel
	Icon string `gorm:"type:varchar(2048);not null;comment:'图标'"`
	Text string `gorm:"type:varchar(10);not null;comment:'名称'"`
	Page string `gorm:"type:varchar(2048);not null;comment:'页面'"`
}

type Case struct {
	BasicModel
	Image    string `gorm:"type:varchar(2048);not null;comment:'封面'"`
	Video    string `gorm:"type:varchar(2048);comment:'视频'"`
	Desc     string `gorm:"type:varchar(20);not null;comment:'描述'"`
	Price    string `gorm:"type:varchar(20);not null;comment:'价格'"`
	Location string `gorm:"type:varchar(20);comment:'地点'"`
	Content  string `gorm:"type:text;comment:'内容'"`
}

type Slides struct {
	BasicModel
	Name  string `gorm:"type:varchar(50);comment:'名称'"`
	Image string `gorm:"type:varchar(2048);comment:'图片'"`
	Link  string `gorm:"type:varchar(2048);comment:'链接'"`
}
