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

package not

import (
	"context"
	"github.com/bukka/wst/app"
	"github.com/bukka/wst/conf/types"
	"github.com/bukka/wst/run/actions"
	"github.com/bukka/wst/run/instances/runtime"
	"github.com/bukka/wst/run/services"
	"time"
)

type ActionMaker struct {
	fnd app.Foundation
}

func CreateActionMaker(fnd app.Foundation) *ActionMaker {
	return &ActionMaker{
		fnd: fnd,
	}
}

func (m *ActionMaker) Make(
	config *types.NotAction,
	svcs services.Services,
	defaultTimeout int,
	actionMaker *actions.ActionMaker,
) (actions.Action, error) {
	if config.Timeout == 0 {
		config.Timeout = defaultTimeout
	}

	newAction, err := actionMaker.MakeAction(config.Action, svcs, config.Timeout)
	if err != nil {
		return nil, err
	}

	return &action{
		fnd:     m.fnd,
		action:  newAction,
		timeout: time.Duration(config.Timeout * 1e6),
	}, nil
}

type action struct {
	fnd     app.Foundation
	action  actions.Action
	timeout time.Duration
}

func (a *action) Timeout() time.Duration {
	return a.timeout
}

func (a *action) Execute(ctx context.Context, runData runtime.Data) (bool, error) {
	a.fnd.Logger().Infof("Executing not action")
	success, err := a.action.Execute(ctx, runData)
	if err != nil {
		return false, err
	}
	a.fnd.Logger().Infof("Executed action resulted to %t - inverting to %t", success, !success)
	if a.fnd.DryRun() {
		// always return success for dry run
		return true, nil
	}
	return !success, nil
}
