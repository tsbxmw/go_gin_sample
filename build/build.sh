mkdir dist
go build -o dist/main ../app/main.go
docker build -f Dockerfile -t $1 .
docker push $1