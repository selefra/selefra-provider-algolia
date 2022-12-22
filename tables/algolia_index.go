package tables

import (
	"context"
	"github.com/selefra/selefra-provider-algolia/algolia_client"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/selefra/selefra-provider-algolia/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlgoliaIndexGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlgoliaIndexGenerator{}

func (x *TableAlgoliaIndexGenerator) GetTableName() string {
	return "algolia_index"
}

func (x *TableAlgoliaIndexGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlgoliaIndexGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlgoliaIndexGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlgoliaIndexGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			conn, err := algolia_client.Connect(ctx, taskClient.(*algolia_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			result, err := conn.ListIndices()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range result.Items {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func getIndexSettings(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	item := result.(search.IndexRes)
	conn, err := algolia_client.Connect(ctx, taskClient.(*algolia_client.Client).Config)
	if err != nil {
		return nil, err
	}
	index := conn.InitIndex(item.Name)
	settings, err := index.GetSettings()
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (x *TableAlgoliaIndexGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlgoliaIndexGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Index name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Index creation date. If empty then the index has no records.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_build_time_secs").ColumnType(schema.ColumnTypeInt).Description("Last build time in seconds.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("Date of last update. An empty string means that the index has no records.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary").ColumnType(schema.ColumnTypeString).Description("Only present if the index is a replica. Contains the name of the related primary index.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entries").ColumnType(schema.ColumnTypeInt).Description("Number of records contained in the index.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_size").ColumnType(schema.ColumnTypeInt).Description("Number of bytes of the index in minified format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_size").ColumnType(schema.ColumnTypeInt).Description("Number of bytes of the index binary file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replicas").ColumnType(schema.ColumnTypeJSON).Description("Only present if the index is a primary index with replicas. Contains the names of all linked replicas.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).Description("Index settings.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := getIndexSettings(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
	}
}

func (x *TableAlgoliaIndexGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAlgoliaSearchMetadataGenerator{}),
		table_schema_generator.GenTableSchema(&TableAlgoliaSearchGenerator{}),
	}
}
