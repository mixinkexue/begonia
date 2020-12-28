package register

import (
	"context"
	cRegister "github.com/MashiroC/begonia/core/register"
	"github.com/MashiroC/begonia/logic"
	"sync"
)

type Register interface {
	Register(name string, info []cRegister.FunInfo) (err error)
	Get(name string) (fs []cRegister.FunInfo, err error)
}

func NewLocalRegister(c *cRegister.CoreRegister) Register {
	return &localRegister{
		c: c,
	}
}

// localRegister 本地注册器
type localRegister struct {
	c *cRegister.CoreRegister
	l sync.Mutex
	m map[string][]cRegister.FunInfo
}

func (r *localRegister) Register(name string, info []cRegister.FunInfo) (err error) {
	ctx := context.WithValue(context.Background(), "info", map[string]string{"connID": "local", "reqID": "null"})
	return r.c.Register(ctx, cRegister.Service{
		Name: name,
		Mode: "avro",
		Funs: info,
	})
}

func (r *localRegister) Get(name string) (fs []cRegister.FunInfo, err error) {
	si, err := r.c.ServiceInfo(name)
	if err != nil {
		return
	}
	fs = si.Funs
	return
}

func NewRemoteRegister(lg *logic.Client) Register {
	return &remoteRegister{
		lg: lg,
	}
}

// remoteRegister 远程注册器
type remoteRegister struct {
	lg *logic.Client
}

func (r *remoteRegister) Register(name string, info []cRegister.FunInfo) (err error) {

	var in _CoreRegisterServiceRegisterIn

	in.F1 = cRegister.Service{
		Name: name,
		Mode: "avro",
		Funs: info,
	}

	b, err := _CoreRegisterServiceRegisterInCoder.Encode(in)
	if err != nil {
		panic(err)
	}

	res := r.lg.CallSync(&logic.Call{
		Service: "REGISTER",
		Fun:     "Register",
		Param:   b,
	})

	return res.Err
}

func (r *remoteRegister) Get(name string) (fs []cRegister.FunInfo, err error) {
	var in _CoreRegisterServiceServiceInfoIn
	in.F1 = name

	b, err := _CoreRegisterServiceServiceInfoInCoder.Encode(in)

	res := r.lg.CallSync(&logic.Call{
		Service: "REGISTER",
		Fun:     "ServiceInfo",
		Param:   b,
	})

	if res.Err != nil {
		err = res.Err
		return
	}

	var out _CoreRegisterServiceServiceInfoOut
	err = _CoreRegisterServiceServiceInfoOutCoder.DecodeIn(res.Result, &out)
	if err != nil {
		return
	}

	fs = out.F1.Funs
	return
}
