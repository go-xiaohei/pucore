package setting

import (
	"github.com/go-xiaohei/pucore/app"
	"github.com/go-xiaohei/pucore/core"
)

type Module struct{}

func (sm *Module) Id() string {
	return "SETTING"
}

func (sm *Module) Prepare(ctx *pucore.ModuleContext) error {
	return nil
}

func (sm *Module) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (sm *Module) Disable(ctx *pucore.ModuleContext) error {
	return pucore.ErrModuleDisableIgnore
}

func (sm *Module) Install(ctx *pucore.ModuleContext) error {
	database := new(app.Db)
	ctx.Injector.Get(database)
	var err error
	if err = database.Sync2(new(Setting)); err != nil {
		return err
	}

	// insert settings
	setting := &Setting{
		Name:   "general",
		UserId: 9,
	}
	setting.Encode(&SettingGeneral{
		Title:          "PUGO",
		SubTitle:       "Simple Blog Engine",
		Keyword:        "pugo,blog,go,golang",
		Description:    "PUGO is a simple blog engine by golang",
		HostName:       "http://localhost",
		HeroImage:      "/img/bg.png",
		TopAvatarImage: "/img/logo.png",
	})

	/*
		setting2 := &Setting{
			Name:   "media",
			UserId: 0,
			Type:   "media",
		}
		setting2.Encode(&SettingMedia{
			MaxFileSize: 10 * 1024,
			ImageFile:   []string{"jpg", "jpeg", "png", "gif", "bmp", "vbmp"},
			DocFile:     []string{"txt", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf"},
			CommonFile:  []string{"zip", "rar"},
			DynamicLink: false,
		})

		setting3 := &Setting{
			Name:   "content",
			UserId: 0,
			Type:   "content",
		}
		setting3.Encode(&SettingContent{
			PageSize:         5,
			RSSFullText:      true,
			RSSNumberLimit:   0,
			TopPage:          0,
			PageDisallowLink: []string{"article", "archive", "feed", "comment", "admin", "sitemap"},
		})

		setting4 := &Setting{
			Name:   "comment",
			UserId: 0,
			Type:   "comment",
		}
		setting4.Encode(&SettingComment{
			IsPager:        false,
			PageSize:       10,
			Order:          "create_time DESC",
			CheckAll:       false,
			CheckNoPass:    true,
			CheckRefer:     true,
			AutoCloseDay:   30,
			SubmitDuration: 60,
			MaxLength:      512,
			MinLength:      2,
		})

		setting5 := &Setting{
			Name:   "menu",
			UserId: 0,
			Type:   "menu",
		}
		setting5.Encode([]*SettingMenu{
			{
				"Home", "/", "Home",
				false,
			},
			{
				"Archive", "/archive", "Archive",
				false,
			},
			{
				"About", "/about.html", "About",
				false,
			},
		})
	*/
	if _, err := database.Insert(setting); err != nil {
		return err
	}
	return nil
}
