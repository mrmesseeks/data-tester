# Data Tester checks data in a database

## Install
```console
git clone https://github.com/mixo/data-tester.git
cd data-tester
docker build -t data-tester:latest .
```  

## Run
To run you need an env file with the following env vars:  
- data_tester_db_driver
- data_tester_db_host
- data_tester_db_port
- data_tester_db_user
- data_tester_db_password
- data_tester_db_name

A run example:  
```console
docker run --env-file=/path/to/.env data-tester:latest day-fluctuation -tn datatester_fixture -dc date -nc int_param,float_param -gc group_param -di -1 -md 40 -nd 10
docker run --env-file=/path/to/.env clickhouse_data_tester.go day-fluctuation --tn yandex_direct_report -dc date -nc impressions,clicks -gc criterion_type -di -2 -md 20 -nd 10
```

sudo docker run --env-file=.env.dev data-tester:latest day-fluctuation --tn adsByCountry -dc date -nc impressions,clicks -gc type -di -2 -md 20 -nd 10
~~~~
// cant do it
go get github.com/olekukonko/tablewriter

go run data_tester.go day-fluctuation --tn adsByCountry -dc date -nc impressions,clicks,installs,downloads -gc type -di -2 -md 20 -nd 10

go run data_tester.go day-fluctuation --tn yandex_direct_report -dc date -nc impressions,clicks -gc account_id -di -2 -md 20 -nd 10
