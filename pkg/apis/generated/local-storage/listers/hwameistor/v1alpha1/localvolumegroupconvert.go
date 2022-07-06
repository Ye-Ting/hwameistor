/*
Copyright 2022 Contributors to The HwameiStor project.

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
	v1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/local-storage/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LocalVolumeGroupConvertLister helps list LocalVolumeGroupConverts.
type LocalVolumeGroupConvertLister interface {
	// List lists all LocalVolumeGroupConverts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LocalVolumeGroupConvert, err error)
	// Get retrieves the LocalVolumeGroupConvert from the index for a given name.
	Get(name string) (*v1alpha1.LocalVolumeGroupConvert, error)
	LocalVolumeGroupConvertListerExpansion
}

// localVolumeGroupConvertLister implements the LocalVolumeGroupConvertLister interface.
type localVolumeGroupConvertLister struct {
	indexer cache.Indexer
}

// NewLocalVolumeGroupConvertLister returns a new LocalVolumeGroupConvertLister.
func NewLocalVolumeGroupConvertLister(indexer cache.Indexer) LocalVolumeGroupConvertLister {
	return &localVolumeGroupConvertLister{indexer: indexer}
}

// List lists all LocalVolumeGroupConverts in the indexer.
func (s *localVolumeGroupConvertLister) List(selector labels.Selector) (ret []*v1alpha1.LocalVolumeGroupConvert, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LocalVolumeGroupConvert))
	})
	return ret, err
}

// Get retrieves the LocalVolumeGroupConvert from the index for a given name.
func (s *localVolumeGroupConvertLister) Get(name string) (*v1alpha1.LocalVolumeGroupConvert, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("localvolumegroupconvert"), name)
	}
	return obj.(*v1alpha1.LocalVolumeGroupConvert), nil
}
