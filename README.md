# Simple *Go* Web Server
Created from a [pluralsight course](https://app.pluralsight.com/library/courses/getting-started-with-go/table-of-contents), this example starts a web server that allows you to create, update, and delete users. No database is currently apart of this setup (will reset each time you re-start the server), but is something I will like to add, along with a HTML interface to handle the aforementioned options. 

## How to Replicate (Windows example)
### Starting from Scratch
If you want to run through the whole course, can follow along and build the code, however first you will want to run the initial setup steps below:
1. Will need to install the [*Go* msi package](https://golang.org/dl/) for Windows
2. Used [Visual Studio Code](https://code.visualstudio.com/) for creating the code, and installed the *Go* extention
3. Also installed [git](https://git-scm.com/downloads) for Windows
4. After git was installed, in Visual Studio Code run: "*ctrl + shift + p*", search for "*go: install/update tools*", select all and hit okay
5. Then can follow along with the course

### Just Downloading the .exe
Can also just download the `webservices.exe` from this repo and execute it, will create the simple web server on `localhost:3000`

After it is running, can use [Advanced REST Client](https://install.advancedrestclient.com/install) to simulate POST, GET, PUT, and DELETE commands to the web server with json commands.

#### Eamples of POST/GET/PUT/DELETE Commands
Post:
```
{
  "FirstName": "Ben",
  "LastName": "Dover"
}
```
Get:
```
http://localhost:3000/users/1
```
Put:
```
{
  "ID":1,
  "FirstName": "Seymour",
  "LastName": "Buts"
}
```
Delete:
```
http://localhost:3000/users/1
```