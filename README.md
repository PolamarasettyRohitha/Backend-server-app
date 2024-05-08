
# Backend Developer

An application for a library of books. It will have 2 types of users. Admin user and a Regular user.
Admin user has the privilege to add or delete books.



To run the project, download the zip file, open cmd promt and navigate to the the below mentioned folder

```
/BackendDeveloper/cmd/web
```
run the web.exe file






## Usage/Examples
Open another command promt and run the following cmds.

### Login 
```http
curl "http://localhost:4000?username=user&password=user@123"
```
Login returns a token which should be included as a http header when calling other api's like home, addBook, deleteBook
Use  the file users.txt to find list of users and their detials available to use. 


*Note* - token is valid for 5 minutes, later it expires. You have to login again to generate a new token and use it.

### Home
Include the token returned by /login api as a http Header for calling /home api as shown below

```
curl -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTcxNTE1ODU5MiwicGFzc3dvcmQiOiJ1c2VyQDEyMyIsInVzZXJuYW1lIjoidXNlciJ9.V7dPVFCXY7M0Nb1QhdJIIUmePxGvBUc9HyXFqwl0Ssw" http://localhost:4000/home
```

This will list the books present in regularUser.csv  and adminUser.csv according to the user

*Note* - here eyJhbGci... is the token returned by /login api.

### addBook
Add Book requires few params names Book Name, Author, Publication Year to be sent when calling /addBook api. 
Please refer the following cmd.

```http
curl -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTcxNTE1ODU5MiwicGFzc3dvcmQiOiJ1c2VyQDEyMyIsInVzZXJuYW1lIjoidXNlciJ9.V7dPVFCXY7M0Nb1QhdJIIUmePxGvBUc9HyXFqwl0Ssw" -d "Book Name=jinglebells&Author=Santa&Publication Year=1990" http://localhost:4000/addBook

```
### deleteBook
Add Book requires  Book Name param to be sent when calling /deleteBook api. 
Please refer the following cmd.
```http
curl -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTcxNTE1ODU5MiwicGFzc3dvcmQiOiJ1c2VyQDEyMyIsInVzZXJuYW1lIjoidXNlciJ9.V7dPVFCXY7M0Nb1QhdJIIUmePxGvBUc9HyXFqwl0Ssw" -d "Book Name=jinglebells" http://localhost:4000/deleteBook

```


*Note* - if you want to add users, please update cmd/web/users.go as accordingly

```go
var regularUsers = map[string]user{
	"user":  user{name: "user", password: "user@123"},
	"user1": user{name: "user1", password: "user@123"},
}

var adminUsers = map[string]user{
	"admin":  user{name: "admin", password: "admin@123", admin: true},
	"admin1": user{name: "admin1", password: "admin1@123", admin: true},
}
```



