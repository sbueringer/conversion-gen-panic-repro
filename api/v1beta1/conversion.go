package v1beta1

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"conversion-gen-panic-repro/api/v1beta2"
)

func (src *ClusterClass) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta2.ClusterClass)

	if err := Convert_v1beta1_ClusterClass_To_v1beta2_ClusterClass(src, dst, nil); err != nil {
		return err
	}

	if dst.Spec.OpenAPIV3Schema.Items != nil && dst.Spec.OpenAPIV3Schema.Items.ExclusiveMaximum != nil {
		// Accessing ExclusiveMaximum to showcase that the generated conversion code is not safe
		// The issue with the generated code is that a cast + unsafe.Pointer is used to convert two
		// structs that are not identical:
		//   out.Items = (*v1beta2.JSONSchemaProps)(unsafe.Pointer(in.Items))
		fmt.Println(*dst.Spec.OpenAPIV3Schema.Items.ExclusiveMaximum)
	}

	return nil
}

func (dst *ClusterClass) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta2.ClusterClass)

	if err := Convert_v1beta2_ClusterClass_To_v1beta1_ClusterClass(src, dst, nil); err != nil {
		return err
	}

	return nil
}
