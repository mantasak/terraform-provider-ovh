// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"
)

func DomainZoneImportResourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"zone_file": schema.StringAttribute{
			CustomType: ovhtypes.TfStringType{},
			Required:   true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Description:         "Zone file that will be imported",
			MarkdownDescription: "Zone file that will be imported",
		},
		"exported_content": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Zone file exported from the API",
			MarkdownDescription: "Zone file exported from the API",
		},
		"zone_name": schema.StringAttribute{
			CustomType: ovhtypes.TfStringType{},
			Required:   true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Description:         "Zone name",
			MarkdownDescription: "Zone name",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type DomainZoneImportModel struct {
	ZoneFile        ovhtypes.TfStringValue `tfsdk:"zone_file" json:"zoneFile"`
	ZoneName        ovhtypes.TfStringValue `tfsdk:"zone_name" json:"zoneName"`
	ExportedContent ovhtypes.TfStringValue `tfsdk:"exported_content" json:"-"`
}

type DomainZoneImportWritableModel struct {
	ZoneFile *ovhtypes.TfStringValue `tfsdk:"zone_file" json:"zoneFile,omitempty"`
}

func (v DomainZoneImportModel) ToCreate() *DomainZoneImportWritableModel {
	res := &DomainZoneImportWritableModel{}

	if !v.ZoneFile.IsUnknown() {
		res.ZoneFile = &v.ZoneFile
	}

	return res
}
