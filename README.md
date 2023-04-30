# golang-server
This project is a hobby project in order to learn Go in combination with api design. By running the application you will be able to read, update and add new animals into your zoo database. Final product will be a docker container that can run on EC2 and will update a database in another docker container. 

### Steps:
- [ ] Create API design
- [X] Create working server
- [X] Create working docker container
- [ ] Implement all endpoints and return json
- [ ] Implement security headers
- [ ] Create working database docker container setup
- [ ] Rewrite in memory database (map) to database
- [ ] Create API gateway
- [ ] Working situation from API gateway to EC2 instance in AWS. 

### Start server
go run main.go and the server will start on port: `3333`. 
