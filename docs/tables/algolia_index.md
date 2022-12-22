# Table: algolia_index

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | Index name. | 
| created_at | timestamp | X | √ | Index creation date. If empty then the index has no records. | 
| last_build_time_secs | int | X | √ | Last build time in seconds. | 
| updated_at | timestamp | X | √ | Date of last update. An empty string means that the index has no records. | 
| primary | string | X | √ | Only present if the index is a replica. Contains the name of the related primary index. | 
| entries | int | X | √ | Number of records contained in the index. | 
| data_size | int | X | √ | Number of bytes of the index in minified format. | 
| file_size | int | X | √ | Number of bytes of the index binary file. | 
| replicas | json | X | √ | Only present if the index is a primary index with replicas. Contains the names of all linked replicas. | 
| settings | json | X | √ | Index settings. | 


