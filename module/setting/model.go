package setting

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	SETTING_TYPE_GENERAL = iota + 1
	SETTING_TYPE_MEDIA
	SETTING_TYPE_CONTENT
	SETTING_TYPE_COMMENT
	SETTING_TYPE_MENU
)

type Setting struct {
	Id         int64
	Name       string `xorm:"VARCHAR(50) notnull index(name)"`
	Value      string `xorm:"TEXT notnull"`
	UserId     int64
	Type       int8  `xorm:"INT(8) index(type)"`
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
	Title               string `json:"title"`
	SubTitle            string `json:"sub_title"`
	Keyword             string `json:"keyword_meta"`
	Description         string `json:"description_meta"`
	HostName            string `json:"host_name"`
	HeroImage           string `json:"hero_image"`
	TopAvatarImage      string `json:"top_avatar_image"`
	TopAvatarIsExternal bool   `json:"top_avatar_extern"`
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
	PageSize         int      `json:"content_page_size"`
	RSSFullText      bool     `json:"rss_full_text"`
	RSSNumberLimit   int      `json:"rss_number_limit"`
	TopPage          int64    `json:"top_page"`
	PageDisallowLink []string `json:"page_disallow_link"`
}

func (sc SettingContent) DisallowLink() string {
	return strings.Join(sc.PageDisallowLink, " ")
}

type SettingComment struct {
	IsPager   bool   `json:"comment_is_pager"`
	PageSize  int    `json:"comment_page_size"`
	Order     string `json:"comment_order"`
	MaxLength int    `json:"comment_max_length"`
	MinLength int    `json:"comment_min_length"`

	CheckAll    bool `json:"comment_check_all"`
	CheckNoPass bool `json:"comment_check_no_pass"`
	CheckRefer  bool `json:"comment_check_refer"`

	AutoCloseDay    int64 `json:"comment_auto_close_day"`
	SubmitDuration  int64 `json:"comment_submit_duration"`
	ShowWaitComment bool  `json:"comment_show_wait"`
}

type SettingMenu struct {
	Name      string `json:"menu_name"`
	Link      string `json:"menu_link"`
	Title     string `json:"menu_title"`
	IsNewPage bool   `json:"menu_new_page"`
}

type SettingMedia struct {
	MaxFileSize int64    `json:"max_file_size"`
	ImageFile   []string `json:"image_file"`
	DocFile     []string `json:"doc_file"`
	CommonFile  []string `json:"common_file"`
	DynamicLink bool     `json:"dync_link"`
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
