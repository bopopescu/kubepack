/*
Copyright The Kubepack Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubepack.dev/kubepack/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Applications returns a ApplicationInformer.
	Applications() ApplicationInformer
	// Bundles returns a BundleInformer.
	Bundles() BundleInformer
	// Charts returns a ChartInformer.
	Charts() ChartInformer
	// Orders returns a OrderInformer.
	Orders() OrderInformer
	// Plans returns a PlanInformer.
	Plans() PlanInformer
	// Products returns a ProductInformer.
	Products() ProductInformer
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

// Applications returns a ApplicationInformer.
func (v *version) Applications() ApplicationInformer {
	return &applicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Bundles returns a BundleInformer.
func (v *version) Bundles() BundleInformer {
	return &bundleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Charts returns a ChartInformer.
func (v *version) Charts() ChartInformer {
	return &chartInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Orders returns a OrderInformer.
func (v *version) Orders() OrderInformer {
	return &orderInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Plans returns a PlanInformer.
func (v *version) Plans() PlanInformer {
	return &planInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Products returns a ProductInformer.
func (v *version) Products() ProductInformer {
	return &productInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
