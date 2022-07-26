/*
Copyright 2018 Openstorage.org

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
	"context"
	"time"

	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SchedulePoliciesGetter has a method to return a SchedulePolicyInterface.
// A group's client should implement this interface.
type SchedulePoliciesGetter interface {
	SchedulePolicies() SchedulePolicyInterface
}

// SchedulePolicyInterface has methods to work with SchedulePolicy resources.
type SchedulePolicyInterface interface {
	Create(ctx context.Context, schedulePolicy *v1alpha1.SchedulePolicy, opts v1.CreateOptions) (*v1alpha1.SchedulePolicy, error)
	Update(ctx context.Context, schedulePolicy *v1alpha1.SchedulePolicy, opts v1.UpdateOptions) (*v1alpha1.SchedulePolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SchedulePolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SchedulePolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SchedulePolicy, err error)
	SchedulePolicyExpansion
}

// schedulePolicies implements SchedulePolicyInterface
type schedulePolicies struct {
	client rest.Interface
}

// newSchedulePolicies returns a SchedulePolicies
func newSchedulePolicies(c *StorkV1alpha1Client) *schedulePolicies {
	return &schedulePolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the schedulePolicy, and returns the corresponding schedulePolicy object, and an error if there is any.
func (c *schedulePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Get().
		Resource("schedulepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SchedulePolicies that match those selectors.
func (c *schedulePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SchedulePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SchedulePolicyList{}
	err = c.client.Get().
		Resource("schedulepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulePolicies.
func (c *schedulePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("schedulepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a schedulePolicy and creates it.  Returns the server's representation of the schedulePolicy, and an error, if there is any.
func (c *schedulePolicies) Create(ctx context.Context, schedulePolicy *v1alpha1.SchedulePolicy, opts v1.CreateOptions) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Post().
		Resource("schedulepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schedulePolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a schedulePolicy and updates it. Returns the server's representation of the schedulePolicy, and an error, if there is any.
func (c *schedulePolicies) Update(ctx context.Context, schedulePolicy *v1alpha1.SchedulePolicy, opts v1.UpdateOptions) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Put().
		Resource("schedulepolicies").
		Name(schedulePolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schedulePolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the schedulePolicy and deletes it. Returns an error if one occurs.
func (c *schedulePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("schedulepolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("schedulepolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched schedulePolicy.
func (c *schedulePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SchedulePolicy, err error) {
	result = &v1alpha1.SchedulePolicy{}
	err = c.client.Patch(pt).
		Resource("schedulepolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
