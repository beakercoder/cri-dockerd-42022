/*
Copyright 2021 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Generated code, generated via: `mockgen k8s.io/kubernetes/cmd/kubelet/network NetworkPlugin > $GOPATH/src/k8s.io/kubernetes/cmd/kubelet/network/testing/mock_network_plugin.go`
// Edited by hand for boilerplate and gofmt.

package testing

import (
	"github.com/golang/mock/gomock"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/Mirantis/cri-dockerd/config"
	"github.com/Mirantis/cri-dockerd/network"
)

type MockNetworkPlugin struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkPluginRecorder
}

// Recorder for MockNetworkPlugin (not exported)
type _MockNetworkPluginRecorder struct {
	mock *MockNetworkPlugin
}

func NewMockNetworkPlugin(ctrl *gomock.Controller) *MockNetworkPlugin {
	mock := &MockNetworkPlugin{ctrl: ctrl}
	mock.recorder = &_MockNetworkPluginRecorder{mock}
	return mock
}

func (_m *MockNetworkPlugin) EXPECT() *_MockNetworkPluginRecorder {
	return _m.recorder
}

func (_m *MockNetworkPlugin) Capabilities() sets.Int {
	ret := _m.ctrl.Call(_m, "Capabilities")
	ret0, _ := ret[0].(sets.Int)
	return ret0
}

func (_m *MockNetworkPlugin) Finish() {
	_m.ctrl.Finish()
}

func (_mr *_MockNetworkPluginRecorder) Capabilities() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Capabilities")
}

func (_m *MockNetworkPlugin) Event(_param0 string, _param1 map[string]interface{}) {
	_m.ctrl.Call(_m, "Event", _param0, _param1)
}

func (_mr *_MockNetworkPluginRecorder) Event(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Event", arg0, arg1)
}

func (_m *MockNetworkPlugin) GetPodNetworkStatus(
	_param0 string,
	_param1 string,
	_param2 config.ContainerID,
) (*network.PodNetworkStatus, error) {
	ret := _m.ctrl.Call(_m, "GetPodNetworkStatus", _param0, _param1, _param2)
	ret0, _ := ret[0].(*network.PodNetworkStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkPluginRecorder) GetPodNetworkStatus(
	arg0, arg1, arg2 interface{},
) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPodNetworkStatus", arg0, arg1, arg2)
}

func (_m *MockNetworkPlugin) Init(
	_param0 network.Host,
	_param1 config.HairpinMode,
	nonMasqueradeCIDR string,
	mtu int,
) error {
	ret := _m.ctrl.Call(_m, "Init", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkPluginRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0, arg1)
}

func (_m *MockNetworkPlugin) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockNetworkPluginRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockNetworkPlugin) SetUpPod(
	_param0 string,
	_param1 string,
	_param2 config.ContainerID,
	annotations, options map[string]string,
) error {
	ret := _m.ctrl.Call(_m, "SetUpPod", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkPluginRecorder) SetUpPod(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetUpPod", arg0, arg1, arg2)
}

func (_m *MockNetworkPlugin) Status() error {
	ret := _m.ctrl.Call(_m, "Status")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkPluginRecorder) Status() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Status")
}

func (_m *MockNetworkPlugin) TearDownPod(
	_param0 string,
	_param1 string,
	_param2 config.ContainerID,
) error {
	ret := _m.ctrl.Call(_m, "TearDownPod", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkPluginRecorder) TearDownPod(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TearDownPod", arg0, arg1, arg2)
}
