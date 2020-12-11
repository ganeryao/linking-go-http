package module

import (
	"github.com/ganeryao/linking-go-agile/common"
	"github.com/ganeryao/linking-go-agile/protos"
)

type SelfComponent interface {
	TestFunc(req *protos.LRequest) *protos.LResult
	DoFunc(req *protos.LRequest) *protos.LResult
}

// Base implements a default module for Component.
type SelfBase struct {
}

func (b *SelfBase) TestFunc(req *protos.LRequest) *protos.LResult {
	return common.ResultOk
}

func (b *SelfBase) DoFunc(req *protos.LRequest) *protos.LResult {
	return common.ResultOk
}
