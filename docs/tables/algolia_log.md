# Table: algolia_log

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| number_api_calls | int | X | √ | Number of API calls. | 
| query_number_hits | int | X | √ | Number of hits returned for the query. | 
| timestamp | timestamp | X | √ | Time when the log entry was created. | 
| method | string | X | √ | HTTP method used for the query, e.g. GET, POST. | 
| exhaustive | bool | X | √ | Exhaustive flags used during the query. | 
| inner_queries | string | X | √ | Contains an object for each performed query with the indexName, queryID, offset, and userToken. | 
| processing_time_ms | int | X | √ | Processing time for the request in milliseconds. | 
| index | string | X | √ | Index the query was executed against, or null for metadata queries. | 
| sha1 | string | X | √ | SHA1 ID of the log entry. | 
| ip | ip | X | √ | IP address of the request client. | 
| answer_code | int | X | √ | Code of the answer. | 
| answer | string | X | √ | Answer body, truncated after 1000 characters. JSON format, but returned as a string due to the truncation. | 
| query_body | string | X | √ | Request body, truncated after 1000 characters. | 
| query_headers | string | X | √ | HTTP headers for the query. | 
| url | string | X | √ | URL of the query request. | 


