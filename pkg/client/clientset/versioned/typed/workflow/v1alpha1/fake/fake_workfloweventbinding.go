// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkflowEventBindings implements WorkflowEventBindingInterface
type FakeWorkflowEventBindings struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var workfloweventbindingsResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "workfloweventbindings"}

var workfloweventbindingsKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "WorkflowEventBinding"}

// Get takes name of the workflowEventBinding, and returns the corresponding workflowEventBinding object, and an error if there is any.
func (c *FakeWorkflowEventBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.WorkflowEventBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workfloweventbindingsResource, c.ns, name), &v1alpha1.WorkflowEventBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowEventBinding), err
}

// List takes label and field selectors, and returns the list of WorkflowEventBindings that match those selectors.
func (c *FakeWorkflowEventBindings) List(opts v1.ListOptions) (result *v1alpha1.WorkflowEventBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workfloweventbindingsResource, workfloweventbindingsKind, c.ns, opts), &v1alpha1.WorkflowEventBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkflowEventBindingList{ListMeta: obj.(*v1alpha1.WorkflowEventBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkflowEventBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workflowEventBindings.
func (c *FakeWorkflowEventBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workfloweventbindingsResource, c.ns, opts))

}

// Create takes the representation of a workflowEventBinding and creates it.  Returns the server's representation of the workflowEventBinding, and an error, if there is any.
func (c *FakeWorkflowEventBindings) Create(workflowEventBinding *v1alpha1.WorkflowEventBinding) (result *v1alpha1.WorkflowEventBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workfloweventbindingsResource, c.ns, workflowEventBinding), &v1alpha1.WorkflowEventBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowEventBinding), err
}

// Update takes the representation of a workflowEventBinding and updates it. Returns the server's representation of the workflowEventBinding, and an error, if there is any.
func (c *FakeWorkflowEventBindings) Update(workflowEventBinding *v1alpha1.WorkflowEventBinding) (result *v1alpha1.WorkflowEventBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workfloweventbindingsResource, c.ns, workflowEventBinding), &v1alpha1.WorkflowEventBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowEventBinding), err
}

// Delete takes name of the workflowEventBinding and deletes it. Returns an error if one occurs.
func (c *FakeWorkflowEventBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workfloweventbindingsResource, c.ns, name), &v1alpha1.WorkflowEventBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkflowEventBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workfloweventbindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkflowEventBindingList{})
	return err
}

// Patch applies the patch and returns the patched workflowEventBinding.
func (c *FakeWorkflowEventBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WorkflowEventBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workfloweventbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WorkflowEventBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkflowEventBinding), err
}
