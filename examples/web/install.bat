@echo off
rem run this script as admin

if not exist web.exe (
    echo Build the example before installing by running "go build"
    goto :exit
)

sc create static-web binpath= "%CD%\web.exe" start= auto DisplayName= "static-web"
sc description static-web "static-web"
net start static-web
sc query static-web

:exit
