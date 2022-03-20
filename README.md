# Base Structure
This structure used for initial a golang micro service project which has been implemented mongo db and redis by default. you can improve the structure and create a pull request it will be pleasure for me to accept them. directories used as below description.

## /app
This directory contains app starter struct with it's methods to start application and also gracefull shutdown, all configurations initialized in this directory. routers and middlewares for web services defined in this directory.

## /cmd
cobra package used in this package for execute application and define additional commands such as database migrations or seeds.

## /config
applicatons config struct and also global config variable named map defined in this directory, it's also contains setter and getter for debug mode and load config from file methods.

## /controller
The Controller's job is to translate incoming requests into outgoing responses. In order to do this, the controller must take request data and pass it into the Service layer. 

## /domain
Domain Model Layer: Responsible for representing concepts of the business, information about the business situation, and business rules. State that reflects the business situation is controlled and used here, even though the technical details of storing it are delegated to the infrastructure. This layer is the heart of business software. Domain entities should not have any direct dependency (like deriving from a base class) on any data access infrastructure framework.

## /internal

## /service

## /pkg
