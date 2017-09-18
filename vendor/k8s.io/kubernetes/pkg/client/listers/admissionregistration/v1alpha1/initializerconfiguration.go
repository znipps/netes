/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/kubernetes/pkg/apis/admissionregistration/v1alpha1"
)

// InitializerConfigurationLister helps list InitializerConfigurations.
type InitializerConfigurationLister interface {
	// List lists all InitializerConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.InitializerConfiguration, err error)
	// Get retrieves the InitializerConfiguration from the index for a given name.
	Get(name string) (*v1alpha1.InitializerConfiguration, error)
	InitializerConfigurationListerExpansion
}

// initializerConfigurationLister implements the InitializerConfigurationLister interface.
type initializerConfigurationLister struct {
	indexer cache.Indexer
}

// NewInitializerConfigurationLister returns a new InitializerConfigurationLister.
func NewInitializerConfigurationLister(indexer cache.Indexer) InitializerConfigurationLister {
	return &initializerConfigurationLister{indexer: indexer}
}

// List lists all InitializerConfigurations in the indexer.
func (s *initializerConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.InitializerConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InitializerConfiguration))
	})
	return ret, err
}

// Get retrieves the InitializerConfiguration from the index for a given name.
func (s *initializerConfigurationLister) Get(name string) (*v1alpha1.InitializerConfiguration, error) {
	key := &v1alpha1.InitializerConfiguration{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("initializerconfiguration"), name)
	}
	return obj.(*v1alpha1.InitializerConfiguration), nil
}
