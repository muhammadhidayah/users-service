# users-service
User service is service to handle all user. this service implement clean architecture. using postgre as database server.
this service will handler 5 Task :
1. CreateUser
2. UpdateUser
3. DeleteUser
4. GetUserByPersonID
5. GetUserByPersonIDAndPassword

# how to ?
1. Install/Run Postgress Database
2. Build this apps
```
go build -o user-service *.go
```
3. Before you star application, set env variable in your terminal
```
DB_HOST=localhost
DB_USER=username
DB_NAME=yourdbname
DB_PASSWORD=yourpassword
```
4. Then run the appss
5. To test this apps use user-cli
