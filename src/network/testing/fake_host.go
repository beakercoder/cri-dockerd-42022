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

package testing

// helper for testing plugins
// a fake host is created here that can be used by plugins for testing

import (
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"

	"github.com/Mirantis/cri-dockerd/network/hostport"
)

type fakeNetworkHost struct {
	fakeNamespaceGetter
	FakePortMappingGetter
	kubeClient clientset.Interface
	Legacy     bool
}

func NewFakeHost(kubeClient clientset.Interface) *fakeNetworkHost {
	host := &fakeNetworkHost{kubeClient: kubeClient, Legacy: true}
	return host
}

func (fnh *fakeNetworkHost) GetPodByName(name, namespace string) (*v1.Pod, bool) {
	return nil, false
}

func (fnh *fakeNetworkHost) GetKubeClient() clientset.Interface {
	return nil
}

func (nh *fakeNetworkHost) SupportsLegacyFeatures() bool {
	return nh.Legacy
}

type fakeNamespaceGetter struct {
	ns string
}

func (nh *fakeNamespaceGetter) GetNetNS(containerID string) (string, error) {
	return nh.ns, nil
}

type FakePortMappingGetter struct {
	PortMaps map[string][]*hostport.PortMapping
}

func (pm *FakePortMappingGetter) GetPodPortMappings(
	containerID string,
) ([]*hostport.PortMapping, error) {
	return pm.PortMaps[containerID], nil
}
