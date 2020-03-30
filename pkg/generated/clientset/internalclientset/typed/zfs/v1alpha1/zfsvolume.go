/*
Copyright 2019 The OpenEBS Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1alpha1"
	scheme "github.com/openebs/zfs-localpv/pkg/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ZFSVolumesGetter has a method to return a ZFSVolumeInterface.
// A group's client should implement this interface.
type ZFSVolumesGetter interface {
	ZFSVolumes(namespace string) ZFSVolumeInterface
}

// ZFSVolumeInterface has methods to work with ZFSVolume resources.
type ZFSVolumeInterface interface {
	Create(*v1alpha1.ZFSVolume) (*v1alpha1.ZFSVolume, error)
	Update(*v1alpha1.ZFSVolume) (*v1alpha1.ZFSVolume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ZFSVolume, error)
	List(opts v1.ListOptions) (*v1alpha1.ZFSVolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ZFSVolume, err error)
	ZFSVolumeExpansion
}

// zFSVolumes implements ZFSVolumeInterface
type zFSVolumes struct {
	client rest.Interface
	ns     string
}

// newZFSVolumes returns a ZFSVolumes
func newZFSVolumes(c *OpenebsV1alpha1Client, namespace string) *zFSVolumes {
	return &zFSVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the zFSVolume, and returns the corresponding zFSVolume object, and an error if there is any.
func (c *zFSVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.ZFSVolume, err error) {
	result = &v1alpha1.ZFSVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zfsvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ZFSVolumes that match those selectors.
func (c *zFSVolumes) List(opts v1.ListOptions) (result *v1alpha1.ZFSVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ZFSVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zfsvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested zFSVolumes.
func (c *zFSVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("zfsvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a zFSVolume and creates it.  Returns the server's representation of the zFSVolume, and an error, if there is any.
func (c *zFSVolumes) Create(zFSVolume *v1alpha1.ZFSVolume) (result *v1alpha1.ZFSVolume, err error) {
	result = &v1alpha1.ZFSVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("zfsvolumes").
		Body(zFSVolume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a zFSVolume and updates it. Returns the server's representation of the zFSVolume, and an error, if there is any.
func (c *zFSVolumes) Update(zFSVolume *v1alpha1.ZFSVolume) (result *v1alpha1.ZFSVolume, err error) {
	result = &v1alpha1.ZFSVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("zfsvolumes").
		Name(zFSVolume.Name).
		Body(zFSVolume).
		Do().
		Into(result)
	return
}

// Delete takes name of the zFSVolume and deletes it. Returns an error if one occurs.
func (c *zFSVolumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zfsvolumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *zFSVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zfsvolumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched zFSVolume.
func (c *zFSVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ZFSVolume, err error) {
	result = &v1alpha1.ZFSVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("zfsvolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}