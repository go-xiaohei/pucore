package boot

import "github.com/go-xiaohei/pucore/core"

type Upgrade struct {
}

func (u *Upgrade) Id() string {
	return "UPGRADE"
}

func (u *Upgrade) Prepare(ctx *pucore.ModuleContext) error {
	return nil

}

func (u *Upgrade) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (u *Upgrade) Disable(ctx *pucore.ModuleContext) error {
	return pucore.ErrModuleDisableIgnore // todo : maybe upgrade-process can be disabled when online update
}
