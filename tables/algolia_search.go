package tables

import (
	"context"
	"github.com/selefra/selefra-provider-algolia/algolia_client"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/selefra/selefra-provider-algolia/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableAlgoliaSearchGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAlgoliaSearchGenerator{}

func (x *TableAlgoliaSearchGenerator) GetTableName() string {
	return "algolia_search"
}

func (x *TableAlgoliaSearchGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAlgoliaSearchGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAlgoliaSearchGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAlgoliaSearchGenerator) GetDataSource() *schema.DataSource {
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
			for i, hit := range result.Hits {
				hitData := map[string]interface{}{}
				for k, v := range hit {
					switch k {

					case "objectID", "_rankingInfo", "_snippetResult", "_highlightResult":
						continue
					default:
						hitData[k] = v
					}
				}
				row := hitRow{
					Index:           parentIndex.Name,
					Rank:            i + 1,
					Hit:             hitData,
					ObjectID:        hit["objectID"].(string),
					RankingInfo:     hit["_rankingInfo"],
					SnippetResult:   hit["_snippetResult"],
					HighlightResult: hit["_highlightResult"],
				}
				resultChannel <- row
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

type hitRow struct {
	Index           string      `json:"index"`
	Rank            int         `json:"rank"`
	ObjectID        string      `json:"objectID"`
	Hit             interface{} `json:"hit"`
	RankingInfo     interface{} `json:"_rankingInfo"`
	SnippetResult   interface{} `json:"_snippetResult"`
	HighlightResult interface{} `json:"_highlightResult"`
}

func (x *TableAlgoliaSearchGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAlgoliaSearchGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("snippet_result").ColumnType(schema.ColumnTypeJSON).Description("Snippet information.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("index").ColumnType(schema.ColumnTypeString).Description("Name of the index for the search result.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rank").ColumnType(schema.ColumnTypeInt).Description("Rank (position) of the search result. The top result is number 1.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hit").ColumnType(schema.ColumnTypeJSON).Description("Hit data of the search result.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("highlight_result").ColumnType(schema.ColumnTypeJSON).Description("Highlight information.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("object_id").ColumnType(schema.ColumnTypeString).Description("Object ID for this search result.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ranking_info").ColumnType(schema.ColumnTypeJSON).Description("Ranking information for the search result.").Build(),
	}
}

func (x *TableAlgoliaSearchGenerator) GetSubTables() []*schema.Table {
	return nil
}
