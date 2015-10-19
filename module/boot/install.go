package boot

import (
	"github.com/go-xiaohei/pucore/core"
	"gopkg.in/inconshreveable/log15.v2"
)

type InstallFlag bool

type Install struct {
}

func (ins *Install) Id() string {
	return "INSTALL"
}

func (ins *Install) Prepare(ctx *pucore.ModuleContext) error {
	// check install flag
	flag := InstallFlag(false)
	if err := ctx.Injector.Has(&flag); err != nil {
		log15.Debug("Boot.Install.Already")
		return nil
	}
	log15.Info("Boot.Install.Start")
	// add hook to Module-Enable-Before position,
	// try to call Install method of InstallModule implementation.
	ctx.Modular.Hook(pucore.MODULAR_HOOK_ENABLE_BEFORE, func(m pucore.IModule) error {
		if installM, ok := m.(InstallModule); ok {
			if err := installM.Install(ctx); err != nil {
				return err
			}
			log15.Info("Boot.Install." + m.Id())
		}
		return nil
	})
	return nil

}

func (ins *Install) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (ins *Install) Disable(ctx *pucore.ModuleContext) error {
	return pucore.ErrModuleDisableIgnore
}

// InstallModule defines modules support install-process.
type InstallModule interface {
	Install(*pucore.ModuleContext) error
}
