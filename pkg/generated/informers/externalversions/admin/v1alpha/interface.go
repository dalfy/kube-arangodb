//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha

import (
	internalinterfaces "github.com/arangodb/kube-arangodb/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ArangoCollections returns a ArangoCollectionInformer.
	ArangoCollections() ArangoCollectionInformer
	// ArangoDatabases returns a ArangoDatabaseInformer.
	ArangoDatabases() ArangoDatabaseInformer
	// ArangoGraphs returns a ArangoGraphInformer.
	ArangoGraphs() ArangoGraphInformer
	// ArangoUsers returns a ArangoUserInformer.
	ArangoUsers() ArangoUserInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ArangoCollections returns a ArangoCollectionInformer.
func (v *version) ArangoCollections() ArangoCollectionInformer {
	return &arangoCollectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ArangoDatabases returns a ArangoDatabaseInformer.
func (v *version) ArangoDatabases() ArangoDatabaseInformer {
	return &arangoDatabaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ArangoGraphs returns a ArangoGraphInformer.
func (v *version) ArangoGraphs() ArangoGraphInformer {
	return &arangoGraphInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ArangoUsers returns a ArangoUserInformer.
func (v *version) ArangoUsers() ArangoUserInformer {
	return &arangoUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
