# yolofigolang
                                                                                    YOLOfi
                                                                                    
                                                                             You live only once

Concept Description:
Made for humans to know each other better.
The service gamifies human interactions, making them fun. 

--------------------

Prototype:
The prototye will aim at Kood Johvi students: and make a version of Guess who the person is.
Each student gets a set of 100 questions, they can choose to answer any 20 (or more). This gives flexibility to the students as per how much information they want to share.

Each student ID now contains a list of questions answered, a list of questions unanswered, and a list of answers to the questions that have been answered.
In MongoDB this all can be part of a single Json document inside a collection (collection is table equivalent of MySQL in MongoDB, however unlike tables, one can do all kind of data in a single collection as well as within a single document. A document in MongoDB corresponds to a row in table in mysql)



--------------------

Commercial Product: Yolofi Teams: 

Yolofi Teams Product Pitch:= 
  Q: How strong is the bonding in your team? 
  Q: Is there a matrix for measuring bonding?

With Yolofi:
 => Measure team bonding.
 => Improve team performance.
 => Form unbreakable friendships.
 => Made for people-first companies.
 => Decrease attrition.
 
 
 The storyline:
 When a new employee joins a company, they find themselves with the unknowns.
 
 They don't know the tastes of other people in their teams, the chilling factor and the vibe factors, the eyebrow raising topics... the sensitivity and culture of the company as a whole and their teammates as individuals.
 
 It takes an average of 2 years for a new employee to become one with the company cutlture and make effortless team work.
    What if we could reduce it to 2 weeks.
    
 With Yolofi Teams, you don't imagine the productivity boost, you live it.


  

Pilot:
The test pilot of Yolofi Teams will happen in one of the sponsors of Kood Johvi prgram. Either Bolt of Transferwise or some other sponsor could try it out with a team who have recruited a new member. 

Could you use it for remote teams:
Yes, even more so. Because remote teams chit chat less and the culture formation and introduction, especially in new recruits becomes more challenging.
---------------------

---------------------
Trying to get technical side below this line.

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
            
 
  
  
