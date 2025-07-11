package v1beta1

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/randfill"

	"conversion-gen-panic-repro/api/v1beta2"
	"conversion-gen-panic-repro/pkg/conversion"
)

func TestFuzzyConversion(t *testing.T) {
	t.Run("for ClusterClass", conversion.FuzzTestFunc(conversion.FuzzTestFuncInput{
		Hub:         &v1beta2.ClusterClass{},
		Spoke:       &ClusterClass{},
		FuzzerFuncs: []fuzzer.FuzzerFuncs{ClusterClassFuncs},
	}))
}

func ClusterClassFuncs(_ runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		hubJSONSchemaProps,
		spokeJSONSchemaProps,
	}
}

func hubJSONSchemaProps(in *v1beta2.JSONSchemaProps, c randfill.Continue) {
	_ = fillHubJSONSchemaProps(in, c)
	in.Items = fillHubJSONSchemaProps(&v1beta2.JSONSchemaProps{}, c)
}

func fillHubJSONSchemaProps(in *v1beta2.JSONSchemaProps, c randfill.Continue) *v1beta2.JSONSchemaProps {
	in.ExclusiveMaximum = ptr.To(c.Bool())
	return in
}

func spokeJSONSchemaProps(in *JSONSchemaProps, c randfill.Continue) {
	_ = fillSpokeJSONSchemaProps(in, c)
	in.Items = fillSpokeJSONSchemaProps(&JSONSchemaProps{}, c)
}

func fillSpokeJSONSchemaProps(in *JSONSchemaProps, c randfill.Continue) *JSONSchemaProps {
	in.ExclusiveMaximum = c.Bool()
	return in
}
