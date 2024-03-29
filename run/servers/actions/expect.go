package actions

import (
	"fmt"
	"github.com/bukka/wst/conf/types"
	"github.com/bukka/wst/run/actions/expect"
	"github.com/bukka/wst/run/parameters"
)

type ExpectAction interface {
	Parameters() parameters.Parameters
	OutputExpectation() *expect.OutputExpectation
	ResponseExpectation() *expect.ResponseExpectation
}

func (m *Maker) makeExpectAction(configAction types.ServerExpectationAction) (ExpectAction, error) {
	switch action := configAction.(type) {
	case *types.ServerOutputExpectation:
		params, err := m.parametersMaker.Make(action.Parameters)
		if err != nil {
			return nil, err
		}
		outputExpectation, err := m.expectMaker.MakeOutputExpectation(&action.Output)
		if err != nil {
			return nil, err
		}
		return &expectOutputAction{
			parameters:        params,
			outputExpectation: outputExpectation,
		}, nil
	case *types.ServerResponseExpectation:
		params, err := m.parametersMaker.Make(action.Parameters)
		if err != nil {
			return nil, err
		}
		responseExpectation, err := m.expectMaker.MakeResponseExpectation(&action.Response)
		if err != nil {
			return nil, err
		}
		return &expectResponseAction{
			parameters:          params,
			responseExpectation: responseExpectation,
		}, nil
	default:
		return nil, fmt.Errorf("invalid server expectation type %t", configAction)
	}
}

func (m *Maker) makeExpectActions(configActions map[string]types.ServerExpectationAction) (map[string]ExpectAction, error) {
	expectActions := make(map[string]ExpectAction, len(configActions))
	var err error
	for key, configAction := range configActions {
		expectActions[key], err = m.makeExpectAction(configAction)
		if err != nil {
			return nil, err
		}
	}
	return expectActions, nil
}

type expectOutputAction struct {
	parameters        parameters.Parameters
	outputExpectation *expect.OutputExpectation
}

func (a *expectOutputAction) OutputExpectation() *expect.OutputExpectation {
	return a.outputExpectation
}

func (a *expectOutputAction) ResponseExpectation() *expect.ResponseExpectation {
	return nil
}

func (a *expectOutputAction) Parameters() parameters.Parameters {
	return a.parameters
}

type expectResponseAction struct {
	parameters          parameters.Parameters
	responseExpectation *expect.ResponseExpectation
}

func (a *expectResponseAction) OutputExpectation() *expect.OutputExpectation {
	return nil
}

func (a *expectResponseAction) ResponseExpectation() *expect.ResponseExpectation {
	return a.responseExpectation
}

func (a *expectResponseAction) Parameters() parameters.Parameters {
	return a.parameters
}
