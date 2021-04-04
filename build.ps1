$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o ./build/main main.go
~\Go\Bin\build-lambda-zip.exe -o ./build/main.zip ./build/main
Remove-Item ./build/main