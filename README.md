

# Timble

## Details

- Postman collection: `file/Timble.postman_collection.json`

## How to run

Create postgre instance, can use docker
```
docker run --name some-postgres -v postgre-vol:/var/lib/postgresql/data -e POSTGRES_PASSWORD=abc -d -p 5432:5432 postgres
```

Extract .env file to environment variables

Run server, with command 
```
go run main.go
```


