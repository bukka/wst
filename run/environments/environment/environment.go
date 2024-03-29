// Copyright 2024 Jakub Zelenka and The WST Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package environment

import (
	"context"
	"github.com/bukka/wst/app"
	"github.com/bukka/wst/conf/types"
	"github.com/bukka/wst/run/environments/environment/output"
	"github.com/bukka/wst/run/environments/task"
	"github.com/bukka/wst/run/services"
	"io"
	"os"
)

type Command struct {
	Name string
	Args []string
}

type Ports struct {
	Start int32
	Used  int32
	End   int32
}

type ContainerRegistryAuth struct {
	Username string
	Password string
}

type ContainerRegistry struct {
	Auth ContainerRegistryAuth
}

type Environment interface {
	Init(ctx context.Context) error
	Destroy(ctx context.Context) error
	RootPath(service services.Service) string
	RunTask(ctx context.Context, service services.Service, cmd *Command) (task.Task, error)
	ExecTaskCommand(ctx context.Context, service services.Service, target task.Task, cmd *Command) error
	ExecTaskSignal(ctx context.Context, service services.Service, target task.Task, signal os.Signal) error
	Output(ctx context.Context, target task.Task, outputType output.Type) (io.Reader, error)
	PortsStart() int32
	PortsEnd() int32
	ReservePort() int32
	ContainerRegistry() *ContainerRegistry
	MarkUsed()
	IsUsed() bool
}

type Maker struct {
	Fnd app.Foundation
}

func CreateMaker(fnd app.Foundation) *Maker {
	return &Maker{
		Fnd: fnd,
	}
}

type CommonEnvironment struct {
	Fnd   app.Foundation
	Used  bool
	Ports Ports
}

func (m *Maker) MakeCommonEnvironment(config *types.CommonEnvironment) *CommonEnvironment {
	return &CommonEnvironment{
		Fnd:  m.Fnd,
		Used: false,
		Ports: Ports{
			Start: config.Ports.Start,
			Used:  config.Ports.Start,
			End:   config.Ports.End,
		},
	}
}

func (e *CommonEnvironment) MarkUsed() {
	e.Used = true
}

func (e *CommonEnvironment) IsUsed() bool {
	return e.Used
}

func (e *CommonEnvironment) PortsStart() int32 {
	return e.Ports.Start
}

func (e *CommonEnvironment) PortsEnd() int32 {
	return e.Ports.End
}

func (e *CommonEnvironment) ReservePort() int32 {
	used := e.Ports.Used
	e.Ports.Used++
	return used
}

func (e *CommonEnvironment) ContainerRegistry() *ContainerRegistry {
	return nil
}

type ContainerEnvironment struct {
	CommonEnvironment
	Registry ContainerRegistry
}

func (m *Maker) MakeContainerEnvironment(config *types.ContainerEnvironment) *ContainerEnvironment {
	return &ContainerEnvironment{
		CommonEnvironment: *m.MakeCommonEnvironment(&config.CommonEnvironment),
		Registry: ContainerRegistry{
			Auth: ContainerRegistryAuth{
				Username: config.Registry.Auth.Username,
				Password: config.Registry.Auth.Password,
			},
		},
	}
}

func (e *ContainerEnvironment) ContainerRegistry() *ContainerRegistry {
	return &e.Registry
}
