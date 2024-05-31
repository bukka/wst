package actions

import (
	"github.com/bukka/wst/app"
	"github.com/bukka/wst/conf/types"
	appMocks "github.com/bukka/wst/mocks/generated/app"
	actionMocks "github.com/bukka/wst/mocks/generated/run/actions/action"
	benchMocks "github.com/bukka/wst/mocks/generated/run/actions/action/bench"
	expectMocks "github.com/bukka/wst/mocks/generated/run/actions/action/expect"
	notMocks "github.com/bukka/wst/mocks/generated/run/actions/action/not"
	parallelMocks "github.com/bukka/wst/mocks/generated/run/actions/action/parallel"
	reloadMocks "github.com/bukka/wst/mocks/generated/run/actions/action/reload"
	requestMocks "github.com/bukka/wst/mocks/generated/run/actions/action/request"
	restartMocks "github.com/bukka/wst/mocks/generated/run/actions/action/restart"
	startMocks "github.com/bukka/wst/mocks/generated/run/actions/action/start"
	stopMocks "github.com/bukka/wst/mocks/generated/run/actions/action/stop"
	expectationsMocks "github.com/bukka/wst/mocks/generated/run/expectations"
	parametersMocks "github.com/bukka/wst/mocks/generated/run/parameters"
	servicesMocks "github.com/bukka/wst/mocks/generated/run/services"
	"github.com/bukka/wst/run/actions/action"
	"github.com/bukka/wst/run/expectations"
	"github.com/bukka/wst/run/parameters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateActionMaker(t *testing.T) {
	fndMock := appMocks.NewMockFoundation(t)
	expectationsMakerMock := expectationsMocks.NewMockMaker(t)
	parametersMakerMock := parametersMocks.NewMockMaker(t)
	tests := []struct {
		name              string
		fnd               app.Foundation
		expectationsMaker expectations.Maker
		parametersMaker   parameters.Maker
	}{
		{
			name:              "create maker",
			fnd:               fndMock,
			expectationsMaker: expectationsMakerMock,
			parametersMaker:   parametersMakerMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateActionMaker(tt.fnd, tt.expectationsMaker, tt.parametersMaker)
			m, ok := got.(*nativeActionMaker)
			assert.True(t, ok)
			assert.Equal(t, tt.fnd, m.fnd)
			assert.NotNil(t, m.benchMaker)
			assert.NotNil(t, m.expectMaker)
			assert.NotNil(t, m.notMaker)
			assert.NotNil(t, m.parallelMaker)
			assert.NotNil(t, m.requestMaker)
			assert.NotNil(t, m.reloadMaker)
			assert.NotNil(t, m.restartMaker)
			assert.NotNil(t, m.startMaker)
			assert.NotNil(t, m.stopMaker)
		})
	}
}

func Test_nativeActionMaker_MakeAction(t *testing.T) {
	tests := []struct {
		name           string
		config         types.Action
		defaultTimeout int
		setupMocks     func(
			*testing.T,
			*nativeActionMaker,
			action.Action,
			*servicesMocks.MockServiceLocator,
			*benchMocks.MockMaker,
			*expectMocks.MockMaker,
			*notMocks.MockMaker,
			*parallelMocks.MockMaker,
			*requestMocks.MockMaker,
			*reloadMocks.MockMaker,
			*restartMocks.MockMaker,
			*startMocks.MockMaker,
			*stopMocks.MockMaker,
		)
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:           "successful bench action creation",
			config:         &types.BenchAction{Service: "svc"},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				benchMaker.On("Make", &types.BenchAction{Service: "svc"}, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful custom expectation action creation",
			config:         &types.CustomExpectationAction{Service: "svc"},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.CustomExpectationAction{Service: "svc"}
				expectMaker.On("MakeCustomAction", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful metrics expectation action creation",
			config:         &types.MetricsExpectationAction{Service: "svc"},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.MetricsExpectationAction{Service: "svc"}
				expectMaker.On("MakeMetricsAction", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful output expectation action creation",
			config:         &types.OutputExpectationAction{Service: "svc"},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.OutputExpectationAction{Service: "svc"}
				expectMaker.On("MakeOutputAction", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful response expectation action creation",
			config:         &types.ResponseExpectationAction{Service: "svc"},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.ResponseExpectationAction{Service: "svc"}
				expectMaker.On("MakeResponseAction", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful not action creation",
			config:         &types.NotAction{Timeout: 2000},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.NotAction{Timeout: 2000}
				notMaker.On("Make", cfg, sl, 5000, m).Return(a, nil)
			},
		},
		{
			name:           "successful parallel action creation",
			config:         &types.ParallelAction{Timeout: 2000},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.ParallelAction{Timeout: 2000}
				parallelMaker.On("Make", cfg, sl, 5000, m).Return(a, nil)
			},
		},
		{
			name:           "successful request action creation",
			config:         &types.RequestAction{Timeout: 2000},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.RequestAction{Timeout: 2000}
				requestMaker.On("Make", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful reload action creation",
			config:         &types.ReloadAction{Timeout: 2000},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.ReloadAction{Timeout: 2000}
				reloadMaker.On("Make", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful restart action creation",
			config:         &types.RestartAction{Timeout: 2000},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.RestartAction{Timeout: 2000}
				restartMaker.On("Make", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful start action creation",
			config:         &types.StartAction{Timeout: 2000},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.StartAction{Timeout: 2000}
				startMaker.On("Make", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "successful stop action creation",
			config:         &types.StopAction{Timeout: 2000},
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
				cfg := &types.StopAction{Timeout: 2000}
				stopMaker.On("Make", cfg, sl, 5000).Return(a, nil)
			},
		},
		{
			name:           "failed action creation due to invalid config type",
			config:         "test",
			defaultTimeout: 5000,
			setupMocks: func(
				t *testing.T,
				m *nativeActionMaker,
				a action.Action,
				sl *servicesMocks.MockServiceLocator,
				benchMaker *benchMocks.MockMaker,
				expectMaker *expectMocks.MockMaker,
				notMaker *notMocks.MockMaker,
				parallelMaker *parallelMocks.MockMaker,
				requestMaker *requestMocks.MockMaker,
				reloadMaker *reloadMocks.MockMaker,
				restartMaker *restartMocks.MockMaker,
				startMaker *startMocks.MockMaker,
				stopMaker *stopMocks.MockMaker,
			) {
			},
			expectError:      true,
			expectedErrorMsg: "unsupported action type: string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fndMock := appMocks.NewMockFoundation(t)
			slMock := servicesMocks.NewMockServiceLocator(t)
			benchMakerMock := benchMocks.NewMockMaker(t)
			expectMakerMock := expectMocks.NewMockMaker(t)
			notMakerMock := notMocks.NewMockMaker(t)
			parallelMakerMock := parallelMocks.NewMockMaker(t)
			requestMakerMock := requestMocks.NewMockMaker(t)
			reloadMakerMock := reloadMocks.NewMockMaker(t)
			restartMakerMock := restartMocks.NewMockMaker(t)
			startMakerMock := startMocks.NewMockMaker(t)
			stopMakerMock := stopMocks.NewMockMaker(t)
			actionMock := actionMocks.NewMockAction(t)

			m := &nativeActionMaker{
				fnd:           fndMock,
				benchMaker:    benchMakerMock,
				expectMaker:   expectMakerMock,
				notMaker:      notMakerMock,
				parallelMaker: parallelMakerMock,
				requestMaker:  requestMakerMock,
				reloadMaker:   reloadMakerMock,
				restartMaker:  restartMakerMock,
				startMaker:    startMakerMock,
				stopMaker:     stopMakerMock,
			}

			tt.setupMocks(
				t,
				m,
				actionMock,
				slMock,
				benchMakerMock,
				expectMakerMock,
				notMakerMock,
				parallelMakerMock,
				requestMakerMock,
				reloadMakerMock,
				restartMakerMock,
				startMakerMock,
				stopMakerMock,
			)

			got, err := m.MakeAction(tt.config, slMock, tt.defaultTimeout)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, got)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actionMock, got)
			}
		})
	}
}
