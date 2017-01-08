@echo off

echo compiling...

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

: ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w ./

cd src

go build -o Gen4DP.exe
move Gen4DP.exe ../

cd ..

:end
echo finished
