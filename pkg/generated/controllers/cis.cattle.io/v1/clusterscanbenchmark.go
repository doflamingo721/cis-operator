/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/rancher/cis-operator/pkg/apis/cis.cattle.io/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type ClusterScanBenchmarkHandler func(string, *v1.ClusterScanBenchmark) (*v1.ClusterScanBenchmark, error)

type ClusterScanBenchmarkController interface {
	generic.ControllerMeta
	ClusterScanBenchmarkClient

	OnChange(ctx context.Context, name string, sync ClusterScanBenchmarkHandler)
	OnRemove(ctx context.Context, name string, sync ClusterScanBenchmarkHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() ClusterScanBenchmarkCache
}

type ClusterScanBenchmarkClient interface {
	Create(*v1.ClusterScanBenchmark) (*v1.ClusterScanBenchmark, error)
	Update(*v1.ClusterScanBenchmark) (*v1.ClusterScanBenchmark, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterScanBenchmark, error)
	List(opts metav1.ListOptions) (*v1.ClusterScanBenchmarkList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterScanBenchmark, err error)
}

type ClusterScanBenchmarkCache interface {
	Get(name string) (*v1.ClusterScanBenchmark, error)
	List(selector labels.Selector) ([]*v1.ClusterScanBenchmark, error)

	AddIndexer(indexName string, indexer ClusterScanBenchmarkIndexer)
	GetByIndex(indexName, key string) ([]*v1.ClusterScanBenchmark, error)
}

type ClusterScanBenchmarkIndexer func(obj *v1.ClusterScanBenchmark) ([]string, error)

type clusterScanBenchmarkController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewClusterScanBenchmarkController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ClusterScanBenchmarkController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &clusterScanBenchmarkController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromClusterScanBenchmarkHandlerToHandler(sync ClusterScanBenchmarkHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.ClusterScanBenchmark
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.ClusterScanBenchmark))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *clusterScanBenchmarkController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.ClusterScanBenchmark))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateClusterScanBenchmarkDeepCopyOnChange(client ClusterScanBenchmarkClient, obj *v1.ClusterScanBenchmark, handler func(obj *v1.ClusterScanBenchmark) (*v1.ClusterScanBenchmark, error)) (*v1.ClusterScanBenchmark, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *clusterScanBenchmarkController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *clusterScanBenchmarkController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *clusterScanBenchmarkController) OnChange(ctx context.Context, name string, sync ClusterScanBenchmarkHandler) {
	c.AddGenericHandler(ctx, name, FromClusterScanBenchmarkHandlerToHandler(sync))
}

func (c *clusterScanBenchmarkController) OnRemove(ctx context.Context, name string, sync ClusterScanBenchmarkHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromClusterScanBenchmarkHandlerToHandler(sync)))
}

func (c *clusterScanBenchmarkController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *clusterScanBenchmarkController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *clusterScanBenchmarkController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *clusterScanBenchmarkController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *clusterScanBenchmarkController) Cache() ClusterScanBenchmarkCache {
	return &clusterScanBenchmarkCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *clusterScanBenchmarkController) Create(obj *v1.ClusterScanBenchmark) (*v1.ClusterScanBenchmark, error) {
	result := &v1.ClusterScanBenchmark{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *clusterScanBenchmarkController) Update(obj *v1.ClusterScanBenchmark) (*v1.ClusterScanBenchmark, error) {
	result := &v1.ClusterScanBenchmark{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *clusterScanBenchmarkController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *clusterScanBenchmarkController) Get(name string, options metav1.GetOptions) (*v1.ClusterScanBenchmark, error) {
	result := &v1.ClusterScanBenchmark{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *clusterScanBenchmarkController) List(opts metav1.ListOptions) (*v1.ClusterScanBenchmarkList, error) {
	result := &v1.ClusterScanBenchmarkList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *clusterScanBenchmarkController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *clusterScanBenchmarkController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.ClusterScanBenchmark, error) {
	result := &v1.ClusterScanBenchmark{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type clusterScanBenchmarkCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *clusterScanBenchmarkCache) Get(name string) (*v1.ClusterScanBenchmark, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.ClusterScanBenchmark), nil
}

func (c *clusterScanBenchmarkCache) List(selector labels.Selector) (ret []*v1.ClusterScanBenchmark, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterScanBenchmark))
	})

	return ret, err
}

func (c *clusterScanBenchmarkCache) AddIndexer(indexName string, indexer ClusterScanBenchmarkIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.ClusterScanBenchmark))
		},
	}))
}

func (c *clusterScanBenchmarkCache) GetByIndex(indexName, key string) (result []*v1.ClusterScanBenchmark, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.ClusterScanBenchmark, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.ClusterScanBenchmark))
	}
	return result, nil
}
