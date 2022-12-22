# Table: algolia_api_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| key | string | X | √ | API key value. | 
| acl | json | X | √ | List of permissions the key contains. | 
| created_at | timestamp | X | √ | The date at which the key has been created. | 
| indexes | json | X | √ | The list of targeted indices, if any. | 
| referers | json | X | √ |  | 
| description | string | X | √ | Description of the key. | 
| max_hits_per_query | int | X | √ | Maximum number of hits this API key can retrieve in one call. | 
| max_queries_per_ip_per_hour | int | X | √ | Maximum number of API calls allowed from an IP address per hour. | 
| query_parameters | string | X | √ | Parameters added to all searches with this key. | 
| validity | timestamp | X | √ | Timestamp of the date at which the key expires. (0 means it will not expire automatically). | 


