#!/usr/bin/env bash
docker run -d \
  -e "ACCEPT_EULA=Y" \
  -e "MSSQL_SA_PASSWORD=yourStrongP@ssword" \
  -e "MSSQL_PID=Express" \
  -p 1433:1433 \
   mcr.microsoft.com/mssql/server:2022-latest

docker exec $(docker ps -alq) /opt/mssql-tools/bin/sqlcmd \
  -S localhost -U SA -P 'yourStrongP@ssword' -Q "CREATE DATABASE cloudquery;"