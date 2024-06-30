# gfdown
This project can help set Macbook as file web service,  
so that you can download file from macbook in android phone,.

## As a developer  
you can start up project in this way:  
```
go run ./cmd/myfileserver/main.go -dir=/path/to/your/files -username=admin -password=secret -port=8080
```
Each option can be ignored, it will take the default value  

## How to build execute file
```
go build -o gfdown ./cmd/gfdown  
```

## for release
```
git tag v2.0.0
git push origin v2.0.0
```
keep the version of .github/workflows/release.yml the same
