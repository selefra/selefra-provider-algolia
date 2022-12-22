# Table: algolia_search_metadata

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| facets_stats | json | X | √ | Statistics for numerical facets. | 
| hits_per_page | int | X | √ | The maximum number of hits returned per page. | 
| message | string | X | √ | Used to return warnings about the query. | 
| query_id | string | X | √ | Unique identifier of the search query, to be sent in Insights methods. This identifier links events back to the search query it represents. | 
| server_used | string | X | √ | Actual host name of the server that processed the request. | 
| user_data | json | X | √ | User data results from the search. | 
| index | string | X | √ | Name of the index for the search result. | 
| exhaustive_facets_count | bool | X | √ | Whether the facet count is exhaustive (true) or approximate (false). | 
| index_used | string | X | √ | Index name used for the query. In the case of an A/B test, the targeted index isn’t always the index used by the query. | 
| parsed_query | string | X | √ | The query string that will be searched, after normalization. Normalization includes removing stop words (if removeStopWords is enabled), and transforming portions of the query string into phrase queries. | 
| query_after_removal | string | X | √ | A markup text indicating which parts of the original query have been removed in order to retrieve a non-empty result set. | 
| ab_test_variant_id | int | X | √ | If a search encounters an index that is being A/B tested, this reports the variant ID of the index used. | 
| around_lat_long | string | X | √ | The computed geo location. Format: ${lat},${lng}, where the latitude and longitude are expressed as decimal floating point numbers. | 
| automatic_radius | string | X | √ | The automatically computed radius. | 
| exhaustive_num_hits | bool | X | √ | Whether the nbHits is exhaustive (true) or approximate (false). | 
| facets | json | X | √ | A mapping of each facet name to the corresponding facet counts. | 
| num_pages | int | X | √ | The number of returned pages. Calculation is based on the total number of hits (nbHits) divided by the number of hits per page (hitsPerPage), rounded up to the nearest integer. | 
| num_hits | int | X | √ | The number of hits matched by the query. | 
| page | int | X | √ | Index of the current page (zero-based). See the page search parameter. | 
| params | json | X | √ | Parameters passed to the search. | 
| processing_time_ms | int | X | √ | Time the server took to process the request, in milliseconds. This does not include network time. | 


