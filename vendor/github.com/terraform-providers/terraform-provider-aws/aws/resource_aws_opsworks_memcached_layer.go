package aws

import (
	"github.com/appilon/terraform-plugin-sdk/schema"
)

func resourceAwsOpsworksMemcachedLayer() *schema.Resource {
	layerType := &opsworksLayerType{
		TypeName:         "memcached",
		DefaultLayerName: "Memcached",

		Attributes: map[string]*opsworksLayerTypeAttribute{
			"allocated_memory": {
				AttrName: "MemcachedMemory",
				Type:     schema.TypeInt,
				Default:  512,
			},
		},
	}

	return layerType.SchemaResource()
}
