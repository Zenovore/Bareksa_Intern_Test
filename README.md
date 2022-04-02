# Bareksa_Intern_Test

Created for Bareksa Backend Engineer Intern Selection Test.
## Tech Stack
- Golang (Echo Framework)
- PostgreSQL
- Pgadmin
- Docker
## How to Run
### Windows
Simply run `.\run.bat` on terminal.
### Others
- `docker build -t bareksa-news .`
- `docker-compose -f docker-compose.yaml up`

## Ports
- `8080` : API
- `5050` : Pgadmin

## Endpoints News

### Create News
- POST `{{url}}/news/create`
- Request Body:
```
{
  "title":"",
  "content":""
}
```
### Search One News
- POST `{{url}}/news/search_one`
- Request Body:
```
{
  "guid":"",
  "title":"",
  "is_published_valid":0/1,
  "is_published":0/1,
  "is_deleted_valid":0/1,
  "is_deleted":0/1  
}
```
- Notes : tidak ada field yang wajib diisi pada body request.

### Search Many News
- POST `{{url}}/news/search_many`
- Request Body:
```
{
  "guid":"",
  "title":"",
  "is_published_valid":0/1,
  "is_published":0/1,
  "is_deleted_valid":0/1,
  "is_deleted":0/1  
}
```
- Notes : tidak ada field yang wajib diisi pada body request.

### Filter News By Topic
- POST `{{url}}/news/filter_topic`
- Request Body:
```
{
  "topic":""
}
```
- Notes : String pada field topic adalah nama topic bukan guid topic.
### Update News
- POST `{{url}}/news/update`
- Request Body:
```
{
  "guid":"",
  "title":"",
  "content":"",
  "added_tags":[],
  "deleted_tags":[],
  "is_published_valid":0/1,
  "is_published":0/1
}
```
- Notes : field yang wajib diisi hanya `guid`, string pada `added_tags` dan `deleted_tags` adalah nama tag bukan guid tag.

### Delete News
- POST `{{url}}/news/delete`
- Request Body:
```
{
  "guid":""
}
```

### Delete News
- POST `{{url}}/news/delete`
- Request Body:
```
{
  "guid":""
}
```
## Endpoints Tags
### Create Tag
- POST `{{url}}/tags/create`
- Request Body:
```
{
  "name":""
}
```
### Search One Tag
- POST `{{url}}/tags/search_one`
- Request Body:
```
{
  "guid":"",
  "name":""
}
```
- Notes: Tidak ada field yang wajib diisi pada body request.

### Update Tag
- POST `{{url}}/tags/update`
- Request Body:
```
{
  "guid":"",
  "name":""
}
```

### Delete Tag
- POST `{{url}}/tags/update`
- Request Body:
```
{
  "guid":""
}
```

