/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedNamespacePlacementLister helps list FederatedNamespacePlacements.
type FederatedNamespacePlacementLister interface {
	// List lists all FederatedNamespacePlacements in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedNamespacePlacement, err error)
	// Get retrieves the FederatedNamespacePlacement from the index for a given name.
	Get(name string) (*v1alpha1.FederatedNamespacePlacement, error)
	FederatedNamespacePlacementListerExpansion
}

// federatedNamespacePlacementLister implements the FederatedNamespacePlacementLister interface.
type federatedNamespacePlacementLister struct {
	indexer cache.Indexer
}

// NewFederatedNamespacePlacementLister returns a new FederatedNamespacePlacementLister.
func NewFederatedNamespacePlacementLister(indexer cache.Indexer) FederatedNamespacePlacementLister {
	return &federatedNamespacePlacementLister{indexer: indexer}
}

// List lists all FederatedNamespacePlacements in the indexer.
func (s *federatedNamespacePlacementLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedNamespacePlacement, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedNamespacePlacement))
	})
	return ret, err
}

// Get retrieves the FederatedNamespacePlacement from the index for a given name.
func (s *federatedNamespacePlacementLister) Get(name string) (*v1alpha1.FederatedNamespacePlacement, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federatednamespaceplacement"), name)
	}
	return obj.(*v1alpha1.FederatedNamespacePlacement), nil
}
