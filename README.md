# yolofigolang
yololife


==============================


Let's first make a GoLang file to push some data to our MongoDB Atlas cluster. After that we will think about deployement.

I will use this readme file as a documentation also as most of the web resources I find on GoLang+MongoDB combination even at their official docs is outdated or deprecated.

---------------------

Below is the Mongolang setup in windows and connecting it to Atlas Cluster, thereafter installing goland driver and get it runnin and communicating with Atlas cluster:

for connecting to mongoDB Atlas cluster:
mongosh "mongodb+srv://cluster0.wghwm.mongodb.net/myFirstDatabase" --username THISISUSERNAME

installing MongoDB driver for golang:
dep ensure -add "go.mongodb.org/mongo-driver/mongo@~1.7.0"

Initialize this dep in the root folder of the new project:
dep init

setup GOPATH in env VAriables to your current project path: or yo no Gucci
and if you don't wanna place all the projects in one folder at one place (like me):
            Remove the Gopkg.toml and Gopkg.lock (if you have it).
            Run,
            a. go mod init <project-name> Replace <project-name> with the name of your project.
            b. Run go mod tidy and it'll add all the dependencies you are using in your project.
            c. Run go build once to make sure your project still builds. If it doesn't, You can add the missing dependencies manually in go.mod.
            Commit go.mod and go.sum(if you need deterministic builds).
  Now you Gucci. GoLang setup is not good. Need to make it better.
  
  ========================
            
 
  
  
