// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/deploymentmanager/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeploymentManagerDeploymentsGetter has a method to return a DeploymentManagerDeploymentInterface.
// A group's client should implement this interface.
type DeploymentManagerDeploymentsGetter interface {
	DeploymentManagerDeployments(namespace string) DeploymentManagerDeploymentInterface
}

// DeploymentManagerDeploymentInterface has methods to work with DeploymentManagerDeployment resources.
type DeploymentManagerDeploymentInterface interface {
	Create(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.CreateOptions) (*v1alpha1.DeploymentManagerDeployment, error)
	Update(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.UpdateOptions) (*v1alpha1.DeploymentManagerDeployment, error)
	UpdateStatus(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.UpdateOptions) (*v1alpha1.DeploymentManagerDeployment, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DeploymentManagerDeployment, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DeploymentManagerDeploymentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeploymentManagerDeployment, err error)
	DeploymentManagerDeploymentExpansion
}

// deploymentManagerDeployments implements DeploymentManagerDeploymentInterface
type deploymentManagerDeployments struct {
	client rest.Interface
	ns     string
}

// newDeploymentManagerDeployments returns a DeploymentManagerDeployments
func newDeploymentManagerDeployments(c *DeploymentmanagerV1alpha1Client, namespace string) *deploymentManagerDeployments {
	return &deploymentManagerDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deploymentManagerDeployment, and returns the corresponding deploymentManagerDeployment object, and an error if there is any.
func (c *deploymentManagerDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	result = &v1alpha1.DeploymentManagerDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeploymentManagerDeployments that match those selectors.
func (c *deploymentManagerDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeploymentManagerDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeploymentManagerDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deploymentManagerDeployments.
func (c *deploymentManagerDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deploymentManagerDeployment and creates it.  Returns the server's representation of the deploymentManagerDeployment, and an error, if there is any.
func (c *deploymentManagerDeployments) Create(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.CreateOptions) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	result = &v1alpha1.DeploymentManagerDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentManagerDeployment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deploymentManagerDeployment and updates it. Returns the server's representation of the deploymentManagerDeployment, and an error, if there is any.
func (c *deploymentManagerDeployments) Update(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.UpdateOptions) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	result = &v1alpha1.DeploymentManagerDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		Name(deploymentManagerDeployment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentManagerDeployment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deploymentManagerDeployments) UpdateStatus(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.UpdateOptions) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	result = &v1alpha1.DeploymentManagerDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		Name(deploymentManagerDeployment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentManagerDeployment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deploymentManagerDeployment and deletes it. Returns an error if one occurs.
func (c *deploymentManagerDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deploymentManagerDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deploymentManagerDeployment.
func (c *deploymentManagerDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	result = &v1alpha1.DeploymentManagerDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deploymentmanagerdeployments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}