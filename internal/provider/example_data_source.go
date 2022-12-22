package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &MessageDataSource{}

func NewMessageDataSource() datasource.DataSource {
	return &MessageDataSource{}
}

// MessageDataSource defines the data source implementation.
type MessageDataSource struct{}

// MessageDataSourceModel describes the data source data model.
type MessageDataSourceModel struct {
	Target  types.String `tfsdk:"target"`
	Message types.String `tfsdk:"message"`
	Id      types.String `tfsdk:"id"`
}

func (d *MessageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_message"
}

func (d *MessageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"target": schema.StringAttribute{
				Optional: true,
			},
			"message": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *MessageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *MessageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MessageDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if data.Target.IsNull() {
		data.Target = types.StringValue("world")
	}

	if resp.Diagnostics.HasError() {
		return
	}

	data.Message = types.StringValue(fmt.Sprintf("hello %s", data.Target.ValueString()))
	data.Id = types.StringValue(fmt.Sprintf("hello %s", data.Target.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}
