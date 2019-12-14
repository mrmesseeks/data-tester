# Data Tester checks data in a database

## Install
```console
git clone https://github.com/mixo/data-tester.git
cd data-tester
docker build -t data-tester:latest .
```  

## Run
To run I need a env file wth the following env vars:  
- export data_tester_db_driver
- export data_tester_db_host
- export data_tester_db_port
- export data_tester_db_user
- export data_tester_db_password
- export data_tester_db_name

A run example:  
```console
docker run --env-file=/path/to/.env data-tester:latest day-fluctuation datatester_fixture date -1 20 10 int_param,float_param group_param a
```
