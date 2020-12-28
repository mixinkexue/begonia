// Code generated by Begonia. DO NOT EDIT.
// versions:
// 	Begonia v1.0.2
// source: register.go
// begonia client file

package register

import (
	"github.com/MashiroC/begonia/core/register"
	"github.com/MashiroC/begonia/internal/coding"
)

var (
	_CoreRegisterServiceRegisterInSchema = `
{
			"namespace":"begonia.func.Register",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":{
				"type": "record",
				"name": "Service",
				"fields":[{"name":"Name","type":"string"}
,{"name":"Mode","type":"string"}
,{"name":"Funs","type":{
				"type": "array",
				"items": {
				"type": "record",
				"name": "FunInfo",
				"fields":[{"name":"Name","type":"string"}
,{"name":"InSchema","type":"string"}
,{"name":"OutSchema","type":"string"}

				]
			}
			}}

				]
			},"alias":"si"}

			]
		}`
	_CoreRegisterServiceRegisterOutSchema = `
{
			"namespace":"begonia.func.Register",
			"type":"record",
			"name":"Out",
			"fields":[
				
			]
		}`
	_CoreRegisterServiceRegisterInCoder  coding.Coder
	_CoreRegisterServiceRegisterOutCoder coding.Coder

	_CoreRegisterServiceServiceInfoInSchema = `
{
			"namespace":"begonia.func.ServiceInfo",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":"string","alias":"serviceName"}

			]
		}`
	_CoreRegisterServiceServiceInfoOutSchema = `
{
			"namespace":"begonia.func.ServiceInfo",
			"type":"record",
			"name":"Out",
			"fields":[
				{"name":"F1","type":{
				"type": "record",
				"name": "Service",
				"fields":[{"name":"Name","type":"string"}
,{"name":"Mode","type":"string"}
,{"name":"Funs","type":{
				"type": "array",
				"items": {
				"type": "record",
				"name": "FunInfo",
				"fields":[{"name":"Name","type":"string"}
,{"name":"InSchema","type":"string"}
,{"name":"OutSchema","type":"string"}

				]
			}
			}}

				]
			},"alias":"si"}

			]
		}`
	_CoreRegisterServiceServiceInfoInCoder  coding.Coder
	_CoreRegisterServiceServiceInfoOutCoder coding.Coder
)

type _CoreRegisterServiceRegisterIn struct {
	F1 register.Service
}

type _CoreRegisterServiceRegisterOut struct {
}

type _CoreRegisterServiceServiceInfoIn struct {
	F1 string
}

type _CoreRegisterServiceServiceInfoOut struct {
	F1 register.Service
}

func init() {
	var err error
	_CoreRegisterServiceRegisterInCoder, err = coding.NewAvro(_CoreRegisterServiceRegisterInSchema)
	if err != nil {
		panic(err)
	}
	_CoreRegisterServiceRegisterOutCoder, err = coding.NewAvro(_CoreRegisterServiceRegisterOutSchema)
	if err != nil {
		panic(err)
	}

	_CoreRegisterServiceServiceInfoInCoder, err = coding.NewAvro(_CoreRegisterServiceServiceInfoInSchema)
	if err != nil {
		panic(err)
	}
	_CoreRegisterServiceServiceInfoOutCoder, err = coding.NewAvro(_CoreRegisterServiceServiceInfoOutSchema)
	if err != nil {
		panic(err)
	}
}
