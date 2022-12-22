package tables

import (
	"context"
	"github.com/selefra/selefra-provider-algolia/algolia_client"
	"github.com/selefra/selefra-provider-algolia/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlgoliaApiKeyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlgoliaApiKeyGenerator{}

func (x *TableAlgoliaApiKeyGenerator) GetTableName() string {
	return "algolia_api_key"
}

func (x *TableAlgoliaApiKeyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlgoliaApiKeyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlgoliaApiKeyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlgoliaApiKeyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			conn, err := algolia_client.Connect(ctx, taskClient.(*algolia_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			result, err := conn.ListAPIKeys()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range result.Keys {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableAlgoliaApiKeyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlgoliaApiKeyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("key").ColumnType(schema.ColumnTypeString).Description("API key value.").
			Extractor(column_value_extractor.StructSelector("Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acl").ColumnType(schema.ColumnTypeJSON).Description("List of permissions the key contains.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date at which the key has been created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("indexes").ColumnType(schema.ColumnTypeJSON).Description("The list of targeted indices, if any.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("referers").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_hits_per_query").ColumnType(schema.ColumnTypeInt).Description("Maximum number of hits this API key can retrieve in one call.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_queries_per_ip_per_hour").ColumnType(schema.ColumnTypeInt).Description("Maximum number of API calls allowed from an IP address per hour.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_parameters").ColumnType(schema.ColumnTypeString).Description("Parameters added to all searches with this key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("validity").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp of the date at which the key expires. (0 means it will not expire automatically).").Build(),
	}
}

func (x *TableAlgoliaApiKeyGenerator) GetSubTables() []*schema.Table {
	return nil
}
