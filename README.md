# Back end

Run through docker by running `docker-compose up`

## Routes
|method|route|description|
|------|------------------------|-------------------------|
|POST|   /v1/auth/login         | auth user
|POST|   /v1/auth/register      | create user                 
|POST|   /v1/agenda             | create task on agenda
|GET|    /v1/agenda             | list tasks from agenda                                                                
|GET|    /v1/agenda/:id         | retrieve one task by id
|PUT|    /v1/agenda/:id         | update some task
