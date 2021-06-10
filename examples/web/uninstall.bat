@echo off
rem run this script as admin

net stop static-web
sc delete static-web
