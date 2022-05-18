# kill -9 $(lsof -ti:50010)
# kill -9 $(lsof -ti:50060)
# kill -9 $(lsof -ti:8090)

# go run ../pkg/api/app/main.go
go run ./pkg/app/main.go

