# Simple program to generate some results in Frisk

## _Args_
--app [name of the app]
--project [project name]
--posts [number of posts to do]
--entries [number of entries per post]
--host [host with port number]

## Example
frisk_test --app=dude --project=frisk --posts=10 --entries=10 --host=localhost:8080

## Query Frisk with PostMan
If used the example above, the following will work in postman

### Search json in body of request
* METHOD: GET
* URL: http://[frisk host ip]:[frisk host port]/v1/frisk
* HEADER: Content-Type/application/json
* BODY:
    * `{
 	"name": "frisk search",
     "searches": [
         {
             "name": "pass",
             "tags": {
                 "tags.app": "dude",
                 "tags.result": "pass"
             }
         },
         {
             "name": "total",
             "tags": {
                 "tags.app": "dude"
             }
         }
     ],
     "search_results": {},
     "formula_temp": "({{index . \"pass\"}} * 100 / {{index . \"total\"}})",
     "operator": "gt",
     "threshold": 39
 }`
 