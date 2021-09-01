# Voting APP
This is a small project to showcase voting app with basic functionality to add, delete candidates and caste votes.

At the end of casting votes winner will be declared which has maximum votes.

## pre-requisites
1. Install Docker
2. Install `task`
3. Port 7070 must be free. You can change the port if you want.

## How to start the application
To start the application quickly
```
task build deploy test
```

## APIs available
1. `/healthz` - Returns health status
2. `/candidate/list/` - List the candidates that are standing in the elections
3. `/candidate/add/` - To add any candidate
4. `/candidate/delete/` - To remove a candidate
5. `/candidate/vote/` - To caste a vote for a candidate
6. `/voting/status` - To get the live status of all the votes that have been casted
7. `/voting/results` - To get the final result of elections

## Key observations
1. For simplicity sake all the APIs are of `GET` method.
2. This is simple application which stores data in-memory so your data will be lost if you close the application.


