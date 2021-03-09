## MyLabDiscoveries Client
# Introduction
In the first phase of development, we are going to analyse 96 samples (6 dyes
each) and show the results to the end user. This results should be shown in tabular format
and graphical format. This machine will be industrial machine with Touch screen.


## Golang Boilerplate
We have used Golang boilerplate to kickstart any go api project.

## SetUp Steps
### 1. Install Go
Refer https://golang.org/doc/install, or for linux users 
```
$ sudo snap install go --classic
```
check go version by typing on a new terminal
```
$ go version
```
If this shows blank, then please ask for help

### 2. Install Postgres
Refer https://www.postgresql.org/download/
Select your OS flavour and follow steps for installation

### 3. Install npm 
(Not sure if you may skip this step)
Refer https://www.npmjs.com/get-npm for installation guide

### 4. Install yarn
Refer https://classic.yarnpkg.com/en/docs/install#debian-stable for linux users
For other users please refer internet

### 5. Set application.yml
Inside conf directory, create a clone file from application.yml.default and name it as application.yml . 
#### 5.1 Setting APP_NAME and APP_PORT
CRITICAL STEPS: 
Set APP_NAME to MyLabDiscoveries if you want to run GUI application. 
Set APP_PORT to 33001.
NOTE: Failing to setup any of the above steps will give a 404 for API responses

#### 5.2 Setting DB_URI
(You may skip this step if you have already set correct DB_URI)
Create a new database named cpagentdb

For ubuntu users type:
```
$ psql -U postgres
```
Then type your postgres password
Then type below command
```
$ CREATE DATABASE cpagentdb;
```

And thus your DB_URI should look this way 
DB_URI: "postgresql://(username):(password)@localhost:5432/cpagentdb?sslmode=disable"
You need to set username and password

### 6. Set config.yml
Inside conf directory, create a clone file from config.yml.default and name it as config.yml. And let it be. 

## Building Binary without rice
This is only backend binary
Make sure your PWD is same as this README's.
```
$ go build
```
If build fails then make sure you were on 'master' branch. 
If master branch build fails then we must have messed pretty bad, please contact us.


### Run
DEPENDENCY: Make sure that cpagent binary is built
Firstly run the migrations, make sure your PWD is same as this README's.
```
$ ./cpagent migrate
```
If that fails then conatact us.

Then type
```
$ ./cpagent start --plc simulator
```
If you have followed all steps correctly the setup should start normally.

## For Building and Running Binary with GUI
Please refer README inside web-client directory.

### Testing

Run test locally
```
$ make test
```

### DB Support

For PostgreSQL use following commands
./cpagent start --plc simulator
./cpagent migrate
./cpagent create_migration filename

### RICE

For embedding react build with Go binary

$ rice embed
$ go build
