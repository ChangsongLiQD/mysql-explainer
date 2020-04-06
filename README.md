# mysql-explainer
A checking tool of mysql is for explaining the sql query.
 It makes sure the select queries have a proper efficiency.
  You can use it for your project mysql queries before 
  your project is released.
  
----
### How to use?

````
// download project
git clone https://github.com/ChangsongLiQD/mysql-explainer.git

// your-project-path/config/config.yaml, change to your own configuration.
log: "log" # log file path
db: # database configuration
  username: "root" # account username
  password: "123456" # account password
  dbname: "chuman_user" # database name
  host: "127.0.0.1" # database host
  port: "3306" # database port
rule: # mysql banned explain fields' value
  selecttype: # mysql explain select_type banned values
    - "SUBQUERY"
    - "DERIVED"
    - "UNION"
    - "UNION RESULT"
  type: # mysql explain type banned values
    - "ALL"
    - "index"
  extra: # mysql explain type banned values
    - "Using filesort"
    - "Using temporary"

// go run main.go [--conf your-own-config-path/yaml-config.yaml]
[--sql your-sql-file-path/sql-file.sql] [--report your-report-file-path/report-file.txt]

`Check your report file to see whether there is a problem in your sqls!`

````

Default Location:

sql-file: your-project-path/sql/example.sql

report-file: your-project-path/report/report.txt

config-file: your-project-path/config/config.yaml

----
### Report Example:

````
SELECT * FROM user.user_phone where phone = xxxx;
----------------------------------
(1) Type: ALL is not valid. 
````

Because the filed phone has no index, it will be an
 all table table searching. It will be more efficient
 if phone field is an index.


