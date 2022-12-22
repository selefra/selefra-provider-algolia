package tables

import (
	"context"
	"github.com/selefra/selefra-provider-algolia/algolia_client"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/selefra/selefra-provider-algolia/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAlgoliaSearchMetadataGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlgoliaSearchMetadataGenerator{}

func (x *TableAlgoliaSearchMetadataGenerator) GetTableName() string {
	return "algolia_search_metadata"
}

func (x *TableAlgoliaSearchMetadataGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlgoliaSearchMetadataGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlgoliaSearchMetadataGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlgoliaSearchMetadataGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			conn, err := algolia_client.Connect(ctx, taskClient.(*algolia_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			parentIndex := task.ParentRawResult.(search.IndexRes)
			index := conn.InitIndex(parentIndex.Name)
			result, err := index.Search(
				"",
				opt.AttributesToRetrieve("*"),
				opt.AttributesToSnippet("*:20"),
				opt.ClickAnalytics(true),
				opt.Facets("*"),
				opt.GetRankingInfo(true),
				opt.HitsPerPage(1000),
				opt.Page(0),
				opt.ResponseFields("*"),
				opt.SnippetEllipsisText("…"),
			)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- metadataRow{
				Index:  parentIndex.Name,
				Result: result,
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

type metadataRow struct {
	Index  string
	Result search.QueryRes
}

func (x *TableAlgoliaSearchMetadataGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlgoliaSearchMetadataGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("facets_stats").ColumnType(schema.ColumnTypeJSON).Description("Statistics for numerical facets.").
			Extractor(column_value_extractor.StructSelector("Result.FacetsStats")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hits_per_page").ColumnType(schema.ColumnTypeInt).Description("The maximum number of hits returned per page.").
			Extractor(column_value_extractor.StructSelector("Result.HitsPerPage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).Description("Used to return warnings about the query.").
			Extractor(column_value_extractor.StructSelector("Result.Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the search query, to be sent in Insights methods. This identifier links events back to the search query it represents.").
			Extractor(column_value_extractor.StructSelector("Result.QueryID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("server_used").ColumnType(schema.ColumnTypeString).Description("Actual host name of the server that processed the request.").
			Extractor(column_value_extractor.StructSelector("Result.ServerUsed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_data").ColumnType(schema.ColumnTypeJSON).Description("User data results from the search.").
			Extractor(column_value_extractor.StructSelector("Result.UserData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("index").ColumnType(schema.ColumnTypeString).Description("Name of the index for the search result.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exhaustive_facets_count").ColumnType(schema.ColumnTypeBool).Description("Whether the facet count is exhaustive (true) or approximate (false).").
			Extractor(column_value_extractor.StructSelector("Result.ExhaustiveFacetsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("index_used").ColumnType(schema.ColumnTypeString).Description("Index name used for the query. In the case of an A/B test, the targeted index isn’t always the index used by the query.").
			Extractor(column_value_extractor.StructSelector("Result.IndexUsed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parsed_query").ColumnType(schema.ColumnTypeString).Description("The query string that will be searched, after normalization. Normalization includes removing stop words (if removeStopWords is enabled), and transforming portions of the query string into phrase queries.").
			Extractor(column_value_extractor.StructSelector("Result.ParsedQuery")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_after_removal").ColumnType(schema.ColumnTypeString).Description("A markup text indicating which parts of the original query have been removed in order to retrieve a non-empty result set.").
			Extractor(column_value_extractor.StructSelector("Result.QueryAfterRemoval")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ab_test_variant_id").ColumnType(schema.ColumnTypeInt).Description("If a search encounters an index that is being A/B tested, this reports the variant ID of the index used.").
			Extractor(column_value_extractor.StructSelector("Result.ABTestVariantID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("around_lat_long").ColumnType(schema.ColumnTypeString).Description("The computed geo location. Format: ${lat},${lng}, where the latitude and longitude are expressed as decimal floating point numbers.").
			Extractor(column_value_extractor.StructSelector("Result.AroundLatLng")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("automatic_radius").ColumnType(schema.ColumnTypeString).Description("The automatically computed radius.").
			Extractor(column_value_extractor.StructSelector("Result.AutomaticRadius")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exhaustive_num_hits").ColumnType(schema.ColumnTypeBool).Description("Whether the nbHits is exhaustive (true) or approximate (false).").
			Extractor(column_value_extractor.StructSelector("Result.ExhaustiveNbHits")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("facets").ColumnType(schema.ColumnTypeJSON).Description("A mapping of each facet name to the corresponding facet counts.").
			Extractor(column_value_extractor.StructSelector("Result.Facets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_pages").ColumnType(schema.ColumnTypeInt).Description("The number of returned pages. Calculation is based on the total number of hits (nbHits) divided by the number of hits per page (hitsPerPage), rounded up to the nearest integer.").
			Extractor(column_value_extractor.StructSelector("Result.NbPages")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_hits").ColumnType(schema.ColumnTypeInt).Description("The number of hits matched by the query.").
			Extractor(column_value_extractor.StructSelector("Result.NbHits")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("page").ColumnType(schema.ColumnTypeInt).Description("Index of the current page (zero-based). See the page search parameter.").
			Extractor(column_value_extractor.StructSelector("Result.Page")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("params").ColumnType(schema.ColumnTypeJSON).Description("Parameters passed to the search.").
			Extractor(column_value_extractor.StructSelector("Result.Params")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("processing_time_ms").ColumnType(schema.ColumnTypeInt).Description("Time the server took to process the request, in milliseconds. This does not include network time.").
			Extractor(column_value_extractor.StructSelector("Result.ProcessingTimeMS")).Build(),
	}
}

func (x *TableAlgoliaSearchMetadataGenerator) GetSubTables() []*schema.Table {
	return nil
}
