## MyLabDiscoveries Client

# Introduction
In the first phase of development, we are going to analyse 96 samples (6 dyes
each) and show the results to the end user. This results should be shown in tabular format
and graphical format. This machine will be industrial machine with Touch screen.

For the Extraction part we have a machine with 2 decks. Deck A and Deck B. These have motors which are controlled by a PLC. Our Industrial PC serves as a master while the PLCs serve as slaves.

## Golang Boilerplate
We have used Golang boilerplate to kickstart any go api project.

# SetUp Steps

### 1. Install Go Language
Refer https://golang.org/doc/install, or for linux users 
```
$ sudo snap install go --classic
```
check go version by typing on a new terminal
```
$ go version
```
If this shows blank, then please ask for help

### 2. Install Postgres Server
Refer https://www.postgresql.org/download/
Select your OS flavour and follow steps for installation

If password authentication fails then refer https://askubuntu.com/questions/413585/postgres-password-authentication-fails

Don't forget to restart postgresql

```
$ sudo systemctl restart postgresql
```

### 3. Install npm 
Refer https://www.npmjs.com/get-npm for installation guide
NOTE: Even if you have npm installed you need to follow this step

Go inside web-client directory
```
$ npm install
```

### 4. Install yarn
Refer https://classic.yarnpkg.com/en/docs/install#debian-stable for linux users
For other users please refer internet

### 5. Set application.yml
Inside conf directory, create a clone file from application.yml.default and name it as application.yml . 
#### 5.1 Setting APP_NAME and APP_PORT
CRITICAL STEPS: 
Set APP_NAME to MyLabDiscoveries if you want to run GUI application. 
Set APP_PORT to 33001.
Set SECRET_KEY to "123456qwerty"
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

NOTE: If you are working on backend then below credentials are highly recommended as these replicate actual machine

username: postgres
password: password
dbname: cpagentdb

### 6. Set config.yml
Inside conf directory, create a clone file from config.yml.default and name it as config.yml. And let it be. 

### 7. Accept Header
Failing to set this header will give 404 Page Not Found Error
Set this Header for Postman or for Any Other Web client
Accept:application/vnd.MyLabDiscoveries.v1

# Building Binary without GUI Support
This is only backend binary
Make sure your PWD is same as this README's.
```
$ go build
```
If build fails then make sure you were on 'master' branch. 
If master branch build fails then we must have messed pretty bad, please contact us.


# For Building with GUI Support
Run the below command in the same PWD and a binary should be created
NOTE: This will create binary for linux platform with amd64 architecture,
in case your's is different please change MakeFile's build command accordingly.

```
$ make build
```

Please refer README inside web-client directory if you are facing any issue and then escalate it to us.


# Run
DEPENDENCY: Make sure that cpagent binary is built

## FOR EXTRACTION
sudo ./run.sh extraction

## FOR RT-PCR
sudo ./run.sh rtpcr

## When there are changes in Migration/DB schema files 

Only if below make command doesn't work then go for individual statements
```
$ make migrate
```


Please drop the cpagentdb database.

For this open connection to postgresql
```
$ psql -U postgres
```
Type your password

And then drop your database by typing
```
$ DROP DATABASE cpagentdb;
```

Recreate the database
```
$ CREATE DATABASE cpagentdb;
```

You may close the connection to database

## Run the migrations, 

make sure your PWD is same as this README's.

If you have changed your branch which has differnt DB schema, please goto Step 1

```
$ ./cpagent migrate
```
If this fails then contact us.

Then type
```
$ ./cpagent start --plc simulator
```
If you have followed all steps correctly the setup should start normally.

Congrats if your setup runs, else feel free to contact us.


# Import CSV

After a successfull latest build from master, type the below command in below format to import a Recipe from CSV
As of version 1.3.0 there is also a provision to add labware details via CSV itself.

```
$ ./cpagent import --csv PATH_TO_CSV
```

PATH_TO_CSV contains name of the csv along its extension.

E.g
```
$ ./cpagent import --csv /home/josh/Downloads/CER.csv
```

# Create Zipped Artifact
```
$ make zip
```

# Create Build and then a Zipped Artifact
```
$ make baz
```

### Testing

Run test locally
```
$ make test
```

### DB Support

For PostgreSQL use following commands
```
$ ./cpagent start --plc simulator
$ ./cpagent migrate
$ ./cpagent create_migration filename
```

### RICE

For embedding react build with Go binary

```
$ rice embed
$ go build
```


### NOTE (For Mac Users Only):
If you want to run this binary locally then do the following changes:

1. In Makefile
comment out linux realted builds
e.g Comment this line in Makefile
`
 GOARCH=amd64 GOOS=linux
`
2. In case you see build error related to `-lrt`
remove `-lrt` from `TEC_1089.go` file

3. If you see build error related to `ComPort.c`
Comment out the lines that are causing trouble

4. For connecting to the database
Please refer mac guides or thy this out https://www.codementor.io/@engineerapart/getting-started-with-postgresql-on-mac-osx-are8jcopb

5. All the config files and utilites can be found on the path `$HOME/cpagent`
