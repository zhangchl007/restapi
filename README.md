#Build REST API with Golang from scratch: PostgreSQL with Gorm and Gin Web Framework

## The sample todo app
-  todo app frontend
-  todo app backend

### Start PostgreSQL in localhost
``` 
 cd build
 docker-compose -f docker-compose-pgsql.yaml up -d

```
# Build todo app backend

```
cd cmd
go build -o todo  main.go
./todo
````

#### Resources 
- todo example: https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
- gin example1: https://github.com/eddycjy/go-gin-example
- gin exmaple2: https://github.com/vsouza/go-gin-boilerplate
- gin awesome: https://github.com/FlowerWrong/awesome-gin
