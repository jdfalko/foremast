/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "foremast.ai/foremast/foremast-barrelman/pkg/apis/deployment/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeploymentMonitorLister helps list DeploymentMonitors.
type DeploymentMonitorLister interface {
	// List lists all DeploymentMonitors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DeploymentMonitor, err error)
	// DeploymentMonitors returns an object that can list and get DeploymentMonitors.
	DeploymentMonitors(namespace string) DeploymentMonitorNamespaceLister
	DeploymentMonitorListerExpansion
}

// deploymentMonitorLister implements the DeploymentMonitorLister interface.
type deploymentMonitorLister struct {
	indexer cache.Indexer
}

// NewDeploymentMonitorLister returns a new DeploymentMonitorLister.
func NewDeploymentMonitorLister(indexer cache.Indexer) DeploymentMonitorLister {
	return &deploymentMonitorLister{indexer: indexer}
}

// List lists all DeploymentMonitors in the indexer.
func (s *deploymentMonitorLister) List(selector labels.Selector) (ret []*v1alpha1.DeploymentMonitor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeploymentMonitor))
	})
	return ret, err
}

// DeploymentMonitors returns an object that can list and get DeploymentMonitors.
func (s *deploymentMonitorLister) DeploymentMonitors(namespace string) DeploymentMonitorNamespaceLister {
	return deploymentMonitorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeploymentMonitorNamespaceLister helps list and get DeploymentMonitors.
type DeploymentMonitorNamespaceLister interface {
	// List lists all DeploymentMonitors in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DeploymentMonitor, err error)
	// Get retrieves the DeploymentMonitor from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DeploymentMonitor, error)
	DeploymentMonitorNamespaceListerExpansion
}

// deploymentMonitorNamespaceLister implements the DeploymentMonitorNamespaceLister
// interface.
type deploymentMonitorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DeploymentMonitors in the indexer for a given namespace.
func (s deploymentMonitorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DeploymentMonitor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeploymentMonitor))
	})
	return ret, err
}

// Get retrieves the DeploymentMonitor from the indexer for a given namespace and name.
func (s deploymentMonitorNamespaceLister) Get(name string) (*v1alpha1.DeploymentMonitor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("deploymentmonitor"), name)
	}
	return obj.(*v1alpha1.DeploymentMonitor), nil
}