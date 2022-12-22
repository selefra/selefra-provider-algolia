package tables

import (
	"context"
	"github.com/selefra/selefra-provider-algolia/algolia_client"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/selefra/selefra-provider-algolia/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlgoliaLogGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlgoliaLogGenerator{}

func (x *TableAlgoliaLogGenerator) GetTableName() string {
	return "algolia_log"
}

func (x *TableAlgoliaLogGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlgoliaLogGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlgoliaLogGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlgoliaLogGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := algolia_client.Connect(ctx, taskClient.(*algolia_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			result, err := conn.GetLogs(opt.Length(1000))
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range result.Logs {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableAlgoliaLogGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlgoliaLogGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("number_api_calls").ColumnType(schema.ColumnTypeInt).Description("Number of API calls.").
			Extractor(column_value_extractor.StructSelector("NbAPICalls")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_number_hits").ColumnType(schema.ColumnTypeInt).Description("Number of hits returned for the query.").
			Extractor(column_value_extractor.StructSelector("QueryNbHits")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("Time when the log entry was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("method").ColumnType(schema.ColumnTypeString).Description("HTTP method used for the query, e.g. GET, POST.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exhaustive").ColumnType(schema.ColumnTypeBool).Description("Exhaustive flags used during the query.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inner_queries").ColumnType(schema.ColumnTypeString).Description("Contains an object for each performed query with the indexName, queryID, offset, and userToken.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("processing_time_ms").ColumnType(schema.ColumnTypeInt).Description("Processing time for the request in milliseconds.").
			Extractor(column_value_extractor.StructSelector("ProcessingTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("index").ColumnType(schema.ColumnTypeString).Description("Index the query was executed against, or null for metadata queries.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sha1").ColumnType(schema.ColumnTypeString).Description("SHA1 ID of the log entry.").
			Extractor(column_value_extractor.StructSelector("SHA1")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip").ColumnType(schema.ColumnTypeIp).Description("IP address of the request client.").
			Extractor(column_value_extractor.StructSelector("IP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("answer_code").ColumnType(schema.ColumnTypeInt).Description("Code of the answer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("answer").ColumnType(schema.ColumnTypeString).Description("Answer body, truncated after 1000 characters. JSON format, but returned as a string due to the truncation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_body").ColumnType(schema.ColumnTypeString).Description("Request body, truncated after 1000 characters.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_headers").ColumnType(schema.ColumnTypeString).Description("HTTP headers for the query.").
			Extractor(column_value_extractor.StructSelector("QueryHeaders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("URL of the query request.").
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
	}
}

func (x *TableAlgoliaLogGenerator) GetSubTables() []*schema.Table {
	return nil
}
