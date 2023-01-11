package provider

import (
	"context"
	"github.com/selefra/selefra-provider-algolia/algolia_client"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "algolia",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var algoliaConfig algolia_client.Configs

				err := config.Unmarshal(&algoliaConfig.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(algoliaConfig.Providers) == 0 {
					algoliaConfig.Providers = append(algoliaConfig.Providers, algolia_client.Config{})
				}

				if algoliaConfig.Providers[0].APIKey == "" {
					algoliaConfig.Providers[0].APIKey = os.Getenv("ALGOLIA_API_KEY")
				}

				if algoliaConfig.Providers[0].APIKey == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing domain in configuration")
				}

				if algoliaConfig.Providers[0].AppID == "" {
					algoliaConfig.Providers[0].AppID = os.Getenv("ALGOLIA_API_ID")
				}

				if algoliaConfig.Providers[0].AppID == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing token in configuration")
				}

				clients, err := algolia_client.NewClients(algoliaConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# app_id: "<YOUR_APP_ID>"
# api_key: "<YOUR_API_KEY>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var clientConfig algolia_client.Configs
				err := config.Unmarshal(&clientConfig.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
