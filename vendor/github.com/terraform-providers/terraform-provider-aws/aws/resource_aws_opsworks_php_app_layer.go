package aws

import (
	"github.com/appilon/terraform-plugin-sdk/schema"
)

func resourceAwsOpsworksPhpAppLayer() *schema.Resource {
	layerType := &opsworksLayerType{
		TypeName:         "php-app",
		DefaultLayerName: "PHP App Server",

		Attributes: map[string]*opsworksLayerTypeAttribute{},
	}

	return layerType.SchemaResource()
}
