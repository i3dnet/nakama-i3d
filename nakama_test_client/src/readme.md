## Nakama client
The client connects with the Nakama server. You need to run the docker-compose file that is in the _infra folder.

The compose file starts the following services:
- Nakama
- PostgreSQL
- Mock server
- api-client

In the docker-compose file, we also run an client. When you after the compose has started, run the client locally than 
automatically the client will connect to the Nakama server. In your logs you will see the following:
```
2025-04-18 09:05:53 DBG Notifications: notifications:{id:"9b3f04fe-ff3d-4624-a035-fbc9c65a8328" subject:"Game Session Created" content:"{\"IpAddress\":\"127.0.0.1\",\"Port\":7777}" code:2 sender_id:"00000000-0000-0000-0000-000000000000" create_time:{seconds:1744959953} persistent:true}
2025-04-18 09:05:53 DBG MatchmakerMatched: ticket:"997b9f32-b2f7-4dd5-835d-73e997bc8e65" token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ5NTk5ODMsIm1pZCI6ImQxNmNmMGJkLTY5NDYtNGE4ZC05YmZiLWMyZmI5NjI5YzhjYS4ifQ.rn8ReGmwoFJH8RXbm7NrPsjeJz1747BtPLfFUA0ROPE" users:{presence:{user_id:"2505a149-76dc-4212-a6c3-31eac6ae81d2" session_id:"8c705ea9-1c23-11f0-81c2-006100a0eb06" username:"d00vjica1a8q0c2lqc80"}} users:{presence:{user_id:"56c2d8ef-fae5-4a39-be5b-73cca89ece85" session_id:"88a8e6f8-1c23-11f0-81c2-006100a0eb06" username:"d00vjgl443niinotvvi0"}} self:{presence:{user_id:"2505a149-76dc-4212-a6c3-31eac6ae81d2" session_id:"8c705ea9-1c23-11f0-81c2-006100a0eb06" username:"d00vjica1a8q0c2lqc80"}}
2025-04-18 09:05:53 DBG Notifications: notifications:{id:"a337668e-5e69-40e9-b13d-797bbcc02968" subject:"Game Session Created" content:"{\"IpAddress\":\"127.0.0.1\",\"Port\":7777}" code:2 sender_id:"00000000-0000-0000-0000-000000000000" create_time:{seconds:1744959953} persistent:true}
2025-04-18 09:05:53 DBG MatchmakerMatched: ticket:"c64bb9a4-6067-48be-b347-3f04e5366716" token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ5NTk5ODMsIm1pZCI6ImMwYzMwMWJmLTYzMzMtNGM5NC1hYzI4LTkzZTc5ZWRhNDM4Ny4ifQ.kzMySt6E11I-I-EK1QVZq2HHEg4RpBO4YAb8KpKHf9o" users:{presence:{user_id:"2505a149-76dc-4212-a6c3-31eac6ae81d2" session_id:"8c705ea9-1c23-11f0-81c2-006100a0eb06" username:"d00vjica1a8q0c2lqc80"}} users:{presence:{user_id:"56c2d8ef-fae5-4a39-be5b-73cca89ece85" session_id:"88a8e6f8-1c23-11f0-81c2-006100a0eb06" username:"d00vjgl443niinotvvi0"}} self:{presence:{user_id:"2505a149-76dc-4212-a6c3-31eac6ae81d2" session_id:"8c705ea9-1c23-11f0-81c2-006100a0eb06" username:"d00vjica1a8q0c2lqc80"}}
```

### SETUP
You need to go in your console to this folder: src/nakama-test-client
Then run the following command:
```
go mod tidy
go mod download
go run main.go
```
This will start the client and connect to the Nakama server. You can then use the client to send requests to the Nakama server and receive responses.
