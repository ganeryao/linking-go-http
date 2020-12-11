package module

import (
	"github.com/ganeryao/linking-go-agile/common"
	"github.com/ganeryao/linking-go-agile/protos"
)

type Component interface {
	TestFunc(req *protos.LRequest) *protos.LResult
	DoFunc(req *protos.LRequest) *protos.LResult
}

// Base implements a default module for Component.
type Base struct {
}

func (b *Base) TestFunc(req *protos.LRequest) *protos.LResult {
	return common.ResultOk
}

func (b *Base) DoFunc(req *protos.LRequest) *protos.LResult {
	return common.ResultOk
}
