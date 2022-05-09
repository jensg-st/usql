


# usql
Usql client for Direktiv
  

## Informations

### Version

1.0.0

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | / | [delete](#delete) |  |
| POST | / | [post](#post) |  |
  


## Paths

### <span id="delete"></span> delete (*Delete*)

```
DELETE /
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Direktiv-ActionID | `header` | string | `string` |  |  |  | On cancel Direktiv sends a DELETE request to
the action with id in the header |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-200) | OK |  |  | [schema](#delete-200-schema) |

#### Responses


##### <span id="delete-200"></span> 200
Status: OK

###### <span id="delete-200-schema"></span> Schema

### <span id="post"></span> post (*Post*)

```
POST /
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Direktiv-ActionID | `header` | string | `string` |  |  |  | direktiv action id is an UUID. 
For development it can be set to 'development' |
| Direktiv-TempDir | `header` | string | `string` |  |  |  | direktiv temp dir is the working directory for that request
For development it can be set to e.g. '/tmp' |
| body | `body` | [PostParamsBody](#post-params-body) | `models.PostParamsBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-200) | OK | The service responds with one JSON payload per SQL query. |  | [schema](#post-200-schema) |
| [default](#post-default) | | generic error response | ✓ | [schema](#post-default-schema) |

#### Responses


##### <span id="post-200"></span> 200 - The service responds with one JSON payload per SQL query.
Status: OK

###### <span id="post-200-schema"></span> Schema
   
  

[PostOKBody](#post-o-k-body)

###### Examples
    
**query**
```json
{
  "queries": [
    {
      "result": [
        {
          "id": "c5c20864-2c1b-47e2-94ab-185a3e8cf9ee",
          "name": 123
        }
      ],
      "success": true
    },
    {
      "result": "INSERT 1",
      "success": true
    }
  ]
}
```

##### <span id="post-default"></span> Default Response
generic error response

###### <span id="post-default-schema"></span> Schema

  

[Error](#error)

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Direktiv-ErrorCode | string | `string` |  |  |  |
| Direktiv-ErrorMessage | string | `string` |  |  |  |

## Models

### <span id="error"></span> error


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| errorCode | string| `string` | ✓ | |  |  |
| errorMessage | string| `string` | ✓ | |  |  |



### <span id="post-o-k-body"></span> postOKBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| queries | [][PostOKBodyQueriesItems](#post-o-k-body-queries-items)| `[]*PostOKBodyQueriesItems` |  | |  |  |



### <span id="post-o-k-body-queries-items"></span> postOKBodyQueriesItems


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` |  | |  |  |
| success | boolean| `bool` |  | |  |  |



### <span id="post-params-body"></span> postParamsBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| connection | string| `string` | ✓ | | Connection string for sql commands. This will be used for all commands and has to be URL-encoded. | `pg://myuser:mypwd@MYDATABAESERVER:5432` |
| queries | [][PostParamsBodyQueriesItems](#post-params-body-queries-items)| `[]*PostParamsBodyQueriesItems` |  | | List of sql commands. Commands can use interpolation:<br/> query: select * from table where id = :'ID'<br/> args: ID=123 | `{"args":["KEY=1"],"query":"select * from mytable where key = :'KEY'"}` |



### <span id="post-params-body-queries-items"></span> postParamsBodyQueriesItems


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| args | []string| `[]string` |  | |  | `ARG=myarg` |
| query | string| `string` |  | |  | `select * from table where value = :'ARG'` |


