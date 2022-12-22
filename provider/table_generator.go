package provider

import (
	"github.com/selefra/selefra-provider-algolia/table_schema_generator"
	"github.com/selefra/selefra-provider-algolia/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableAlgoliaApiKeyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlgoliaIndexGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableAlgoliaLogGenerator{}),
	}
}
