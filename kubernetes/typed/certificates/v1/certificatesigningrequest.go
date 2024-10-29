/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"

	v1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	certificatesv1 "k8s.io/client-go/applyconfigurations/certificates/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CertificateSigningRequestsGetter has a method to return a CertificateSigningRequestInterface.
// A group's client should implement this interface.
type CertificateSigningRequestsGetter interface {
	CertificateSigningRequests() CertificateSigningRequestInterface
}

// CertificateSigningRequestInterface has methods to work with CertificateSigningRequest resources.
type CertificateSigningRequestInterface interface {
	Create(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.CreateOptions) (*v1.CertificateSigningRequest, error)
	Update(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (*v1.CertificateSigningRequest, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (*v1.CertificateSigningRequest, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CertificateSigningRequest, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CertificateSigningRequestList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CertificateSigningRequest, err error)
	Apply(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CertificateSigningRequest, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CertificateSigningRequest, err error)
	UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (*v1.CertificateSigningRequest, error)

	CertificateSigningRequestExpansion
}

// certificateSigningRequests implements CertificateSigningRequestInterface
type certificateSigningRequests struct {
	*gentype.ClientWithListAndApply[*v1.CertificateSigningRequest, *v1.CertificateSigningRequestList, *certificatesv1.CertificateSigningRequestApplyConfiguration]
}

// newCertificateSigningRequests returns a CertificateSigningRequests
func newCertificateSigningRequests(c *CertificatesV1Client) *certificateSigningRequests {
	return &certificateSigningRequests{
		gentype.NewClientWithListAndApply[*v1.CertificateSigningRequest, *v1.CertificateSigningRequestList, *certificatesv1.CertificateSigningRequestApplyConfiguration](
			"certificatesigningrequests",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.CertificateSigningRequest { return &v1.CertificateSigningRequest{} },
			func() *v1.CertificateSigningRequestList { return &v1.CertificateSigningRequestList{} },
			gentype.PrefersProtobuf[*v1.CertificateSigningRequest](),
		),
	}
}

// UpdateApproval takes the top resource name and the representation of a certificateSigningRequest and updates it. Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *certificateSigningRequests) UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (result *v1.CertificateSigningRequest, err error) {
	result = &v1.CertificateSigningRequest{}
	err = c.GetClient().Put().
		UseProtobufAsDefault().
		Resource("certificatesigningrequests").
		Name(certificateSigningRequestName).
		SubResource("approval").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}
