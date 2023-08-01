package handler

import (
	"github.com/ahmetson/common-lib/data_type/key_value"
	"github.com/ahmetson/service-lib/communication/command"
	"github.com/ahmetson/service-lib/communication/message"
	"github.com/ahmetson/service-lib/log"
	"github.com/ahmetson/service-lib/remote"
)

var onAdd = func(request message.Request, _ *log.Logger, _ ...*remote.ClientSocket) message.Reply {
	x, err := request.Parameters.GetUint64("x")
	if err != nil {
		return request.Fail("request.Parameters: " + err.Error())
	}

	y, err := request.Parameters.GetUint64("y")
	if err != nil {
		return request.Fail("request.Parameters: " + err.Error())
	}

	sum := x + y

	parameters := key_value.Empty()
	parameters.Set("z", sum)

	return request.Ok(parameters)
}

var onMul = func(request message.Request, _ *log.Logger, _ ...*remote.ClientSocket) message.Reply {
	x, err := request.Parameters.GetUint64("x")
	if err != nil {
		return request.Fail("request.Parameters: " + err.Error())
	}

	y, err := request.Parameters.GetUint64("y")
	if err != nil {
		return request.Fail("request.Parameters: " + err.Error())
	}

	sum := x * y

	parameters := key_value.Empty()
	parameters.Set("z", sum)

	return request.Ok(parameters)
}

func Add() *command.Route {
	return command.NewRoute("add", onAdd)
}
func Mul() *command.Route {
	return command.NewRoute("mul", onMul)
}
