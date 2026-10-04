go mod init main

go run math.go

go clean -testcache

go test -v -run ^TestSoma$


Set-Content -Path .\go.mod -Value 'module projeto-ci-cd
>>
>> go 1.22'

go test -v
