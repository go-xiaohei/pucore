package setting

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Setting struct {
	Id         int64
	Name       string `xorm:"VARCHAR(20) not nullindex(name)"`
	Value      string `xorm:"TEXT notnull"`
	UserId     int64
	CreateTime int64 `xorm:"INT(12) created"`
}

func (s *Setting) Encode(v interface{}) {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	s.Value = string(bytes)
}

func (s *Setting) Decode(v interface{}) error {
	return json.Unmarshal([]byte(s.Value), v)
}

type SettingGeneral struct {
	Title               string
	SubTitle            string
	Keyword             string
	Description         string
	HostName            string
	HeroImage           string
	TopAvatarImage      string
	TopAvatarIsExternal bool
}

func (sg SettingGeneral) FullTitle() string {
	return fmt.Sprintf("%s - %s", sg.Title, sg.SubTitle)
}

func (sg SettingGeneral) TopAvatarUrl(themeLink string) string {
	if sg.TopAvatarIsExternal {
		return sg.TopAvatarImage
	}
	return themeLink + sg.TopAvatarImage
}

type SettingContent struct {
	PageSize         int
	RSSFullText      bool
	RSSNumberLimit   int
	TopPage          int64
	PageDisallowLink []string
}

func (sc SettingContent) DisallowLink() string {
	return strings.Join(sc.PageDisallowLink, " ")
}

type SettingComment struct {
	IsPager   bool
	PageSize  int
	Order     string
	MaxLength int
	MinLength int

	CheckAll    bool
	CheckNoPass bool
	CheckRefer  bool

	AutoCloseDay    int64
	SubmitDuration  int64
	ShowWaitComment bool
}

type SettingMenu struct {
	Name      string
	Link      string
	Title     string
	IsNewPage bool
}

type SettingMedia struct {
	MaxFileSize int64
	ImageFile   []string
	DocFile     []string
	CommonFile  []string
	DynamicLink bool
}

func (sm SettingMedia) Image() string {
	return strings.Join(sm.ImageFile, " ")
}

func (sm SettingMedia) Doc() string {
	return strings.Join(sm.DocFile, " ")
}

func (sm SettingMedia) Common() string {
	return strings.Join(sm.CommonFile, " ")
}
