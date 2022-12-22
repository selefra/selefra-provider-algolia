# Table: algolia_search

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| snippet_result | json | X | √ | Snippet information. | 
| index | string | X | √ | Name of the index for the search result. | 
| rank | int | X | √ | Rank (position) of the search result. The top result is number 1. | 
| hit | json | X | √ | Hit data of the search result. | 
| highlight_result | json | X | √ | Highlight information. | 
| object_id | string | X | √ | Object ID for this search result. | 
| ranking_info | json | X | √ | Ranking information for the search result. | 


