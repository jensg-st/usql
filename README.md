
# usql 1.0.0

Usql client for Direktiv

---
- #### Category: Database
- #### Image: direktiv/usql 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/usql/issues
- #### URL: https://github.com/direktiv-apps/usql
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About usql

[Usql](https://github.com/xo/usql) is a universal command-line interface for PostgreSQL, MySQL, Oracle Database, SQLite3, Microsoft SQL Server, and many other databases including NoSQL and non-relational databases. 
This service provides easy access to database from Direktiv and supports variables and interpolation as well. 

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: usql
    image: direktiv/usql
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: insert 
     type: action
     action:
       function: usql
       input:
         connection: pg://direktiv:rbRIw3I6%7B5%29%5Dda%287m.R_%5D%3Eh%5D@192.168.0.177:5432
         queries:
           - query: select * from mydata
   ```
   #### Variables
   ```yaml
   - id: insert 
     type: action
     action:
       function: usql
       input:
         connection: pg://direktiv:rbRIw3I6%7B5%29%5Dda%287m.R_%5D%3Eh%5D@192.168.0.177:5432
         queries:
           - query: select * from mydata where id = :'ID'
             args:
             - ID=3a7ff4a5-fadb-4d41-97bb-86917cdad46a
   ```
   #### Insert
   ```yaml
   - id: insert 
     type: action
     action:
       function: usql
       input:
         connection: pg://direktiv:rbRIw3I6%7B5%29%5Dda%287m.R_%5D%3Eh%5D@192.168.0.177:5432
         queries:
           - query: insert into testme values ('hello world')
   ```

### Responses
  The service responds with one JSON payload per SQL query.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
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

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| queries | [][PostOKBodyQueriesItems](#post-o-k-body-queries-items)| `[]*PostOKBodyQueriesItems` |  | |  |  |


#### <span id="post-o-k-body-queries-items"></span> postOKBodyQueriesItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` |  | |  |  |
| success | boolean| `bool` |  | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| connection | string| `string` | ✓ | | Connection string for sql commands. This will be used for all commands and has to be URL-encoded. | `pg://myuser:mypwd@MYDATABAESERVER:5432` |
| queries | [][PostParamsBodyQueriesItems](#post-params-body-queries-items)| `[]*PostParamsBodyQueriesItems` |  | | List of sql commands. Commands can use interpolation:<br/> query: select * from table where id = :'ID'<br/> args: ID=123 | `{"args":["KEY=1"],"query":"select * from mytable where key = :'KEY'"}` |


#### <span id="post-params-body-queries-items"></span> postParamsBodyQueriesItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| args | []string| `[]string` |  | |  | `ARG=myarg` |
| query | string| `string` |  | |  | `select * from table where value = :'ARG'` |

 
