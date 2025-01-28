//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializeSimpleService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeSimpleDatabase() *DatabaseRepository {
	wire.Build(
		NewDatabaseRepository,
		NewDatabaseMysql,
		NewDatabasePostgresql,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooService, NewFooRepository)
var barSet = wire.NewSet(NewBarService, NewBarRepository)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// Injector yang salah
//func InitializeHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

var helloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

func InitializeHelloService() *HelloService {
	wire.Build(NewHelloService, helloSet)
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBar() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

// InitializedReader returns an implementation of io.Reader initialized with the standard input (os.Stdin).
// Menginject interface io Reader dengan value Stdin
func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewFile, NewConnection)
	return nil, nil
}
