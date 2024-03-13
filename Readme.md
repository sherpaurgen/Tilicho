
# ...

Project folder structure

bin - compiled binary for deployment

cmd/api - app specific codes

internal - code for talking to DB, data validation, sending email notification.Code in cmd/api will import packages in internal directory

migrations - sql migration files for db

remote - config files and setup script 



## Architecture

`HTTPLayer` : Only handles unmarshaling and marshaling request/response 

`ServiceLayer`: All business logic ->validating user has right access/authorization

`RepositoryLayer`: interacting to database and provide response to service layer

`ClientLayer`: Interface for interacting with third party api's

```
User --> Http layer -- servicelayer -- repositoryLayer --- Database

                            ||
                        Client Layer
                            ||
                        3rd Party apps    
