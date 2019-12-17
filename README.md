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
docker run --env-file=/path/to/.env data-tester:latest day-fluctuation datatester_fixture date -1 20 10 int_param,float_param group_param a
```
