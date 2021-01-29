package serialization

type UserInfo struct {
	NickName  string `json:"nick_name" gorm:"type:varchar(50);comment:'昵称'"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(1024);comment:'头像'"`
	OpenID    string `json:"open_id" gorm:"type:varchar(512);uniqueIndex:OpenIDUniq;comment:'OpenID'"`
	Phone     string `json:"phone" gorm:"type:varchar(20);comment:'手机号'"`
	Gender    int    `json:"gender" gorm:"type:int;comment:'性别'"`
	Province  string `json:"province" gorm:"type:varchar(50);comment:'省会'"`
	City      string `json:"city" gorm:"type:varchar(50);comment:'城市'"`
	Language  string `json:"language" gorm:"type:varchar(50);comment:'语言'"`
	Country   string `json:"country" gorm:"type:varchar(50);comment:'国家'"`
	Darktheme bool   `json:"dark_theme" gorm:"default:0;type:tinyint(1);comment:'是否暗黑主题'"`
}
