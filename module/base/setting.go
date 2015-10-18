package base

import "github.com/go-xiaohei/pucore/core"

type Setting struct {
}

func (s *Setting) Id() string {
	return "SETTING"
}

func (s *Setting) Prepare(ctx *pucore.ModuleContext) error {
	return nil
}

func (s *Setting) Enable(ctx *pucore.ModuleContext) error {
	return nil
}

func (s *Setting) Disable(ctx *pucore.ModuleContext) error {
	return nil
}

func (s *Setting) Install(ctx *pucore.ModuleContext) error {
	println("setting install")
	return nil
}
