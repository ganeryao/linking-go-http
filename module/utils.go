package module

import (
	"errors"
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/ganeryao/linking-go-agile/common"
	e "github.com/ganeryao/linking-go-agile/errors"
	"github.com/ganeryao/linking-go-agile/protos"
	"github.com/ganeryao/linking-go-http/constants"
	"reflect"
	"runtime/debug"
	"strconv"
)

type HandlerMsg struct {
	Uid     string
	Api     string
	Msg     interface{}
}

func DoHandleMsg(msg HandlerMsg) (*protos.LResult, error) {
	handler := handlers[msg.Api]
	args := []reflect.Value{handler.Receiver, reflect.ValueOf(msg)}
	ret, err := call(handler.Method, args)
	if err != nil {
		_ = log4go.Error("DoHandleMsg=============" + err.Error())
		return common.OfResultFail(msg.Api, e.ErrInternalCode, err.Error()), err
	}
	if ret != nil {
		return ret.(*protos.LResult), nil
	}
	return common.ResultOk, nil
}

// call calls a method that returns an interface and an error and recovers in case of panic
func call(method reflect.Method, args []reflect.Value) (rets interface{}, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			// Try to use logger from context here to help trace error cause
			stackTrace := debug.Stack()
			stackTraceAsRawStringLiteral := strconv.Quote(string(stackTrace))
			_ = log4go.Error("panic - lk http/dispatch: methodName=%s panicData=%v stackTrace=%s", method.Name, rec, stackTraceAsRawStringLiteral)

			if s, ok := rec.(string); ok {
				err = errors.New(s)
			} else {
				err = fmt.Errorf("rpc call internal error - %s: %v", method.Name, rec)
			}
		}
	}()

	r := method.Func.Call(args)
	// r can have 0 length in case of notify handlers
	// otherwise it will have 2 outputs: an interface and an error
	if len(r) == 2 {
		if v := r[1].Interface(); v != nil {
			err = v.(error)
		} else if !r[0].IsNil() {
			rets = r[0].Interface()
		} else {
			err = constants.ErrReplyShouldBeNotNull
		}
	}
	return
}