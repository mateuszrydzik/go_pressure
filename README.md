# go pressure;
#### gathering pressure data

i wondered whether my random headaches and changes in well-being were caused by sudden changes in atmospheric pressure. instead of checking the current weather conditions, i wrote a Go application that fetches the current pressure value from the IMGW API and then saves it to a database.

the following data is recorded in the database:

- station: selected measurement station name
- pressure: pressure value
- date: measurement day, month, and year
- hour: measurement hour

## setup

this repo provides docker-compose file with go and postgresql containers. their use is purely for testing. you can use your local go installation to compile the app, and an existing postgresql database to connect to.

### 1. environment variables

create a `.env` file in the main directory, based on the `.envtemplate` file, and fill it with values.

fill the `.env` file with the appropriate values to connect the application to the target database. if you want to use the application and database inside the containers, you only need to fill database name, user name and password.

### 2. running the application

if you decide to use the provided docker-compose file, you can run the following command to start the application:
`docker-compose up --build`

next time you want to start the application, you can use the following command:
`docker-compose up`

### 3. building the application

if you decide to use go container, you can access the container's bash by running the following command:
`docker exec -it {CONTAINER_NAME}_app bash`

with the default value of CONTAINER_NAME, the command will look like this:
`docker exec -it pressure_app bash`


you can build the application by running the following command:
`go build main.go`

if you build the file in the container, but want to run it in specific environment, you need to modify the command with the appropriate values for the target operating system and architecture. for example, for macOS with apple m1, the command will look like this:
`GOOS=darwin GOARCH=arm64 go build main.go`.

the list of available GOOS and GOARCH can be found at:
https://gist.github.com/zfarbp/121a76d5a3fde562c3955a606a9d6fcc

executable file will be built in the application directory.

## what does the app do?

when you run the application, it will make a request to the imgw api, and then save the data to the database. the main use case and purpose of this application is to add a cron job to regularly fetch the data (e.g. three times a day).
