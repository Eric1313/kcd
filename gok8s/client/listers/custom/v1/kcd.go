/*
Copyright 2019 The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/eric1313/kcd/gok8s/apis/custom/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KCDLister helps list KCDs.
type KCDLister interface {
	// List lists all KCDs in the indexer.
	List(selector labels.Selector) (ret []*v1.KCD, err error)
	// KCDs returns an object that can list and get KCDs.
	KCDs(namespace string) KCDNamespaceLister
	KCDListerExpansion
}

// kCDLister implements the KCDLister interface.
type kCDLister struct {
	indexer cache.Indexer
}

// NewKCDLister returns a new KCDLister.
func NewKCDLister(indexer cache.Indexer) KCDLister {
	return &kCDLister{indexer: indexer}
}

// List lists all KCDs in the indexer.
func (s *kCDLister) List(selector labels.Selector) (ret []*v1.KCD, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KCD))
	})
	return ret, err
}

// KCDs returns an object that can list and get KCDs.
func (s *kCDLister) KCDs(namespace string) KCDNamespaceLister {
	return kCDNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KCDNamespaceLister helps list and get KCDs.
type KCDNamespaceLister interface {
	// List lists all KCDs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.KCD, err error)
	// Get retrieves the KCD from the indexer for a given namespace and name.
	Get(name string) (*v1.KCD, error)
	KCDNamespaceListerExpansion
}

// kCDNamespaceLister implements the KCDNamespaceLister
// interface.
type kCDNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KCDs in the indexer for a given namespace.
func (s kCDNamespaceLister) List(selector labels.Selector) (ret []*v1.KCD, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KCD))
	})
	return ret, err
}

// Get retrieves the KCD from the indexer for a given namespace and name.
func (s kCDNamespaceLister) Get(name string) (*v1.KCD, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kcd"), name)
	}
	return obj.(*v1.KCD), nil
}
