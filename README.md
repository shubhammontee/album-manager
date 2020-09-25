# album-manager
album manager to manage albums and photos


Steps to launch the whole application

       FOR ALBUM-MANAGER

1. git clone https://github.com/suvamsingh/album-manager.git
2. cd album-manager

3. docker-compose -f .\docker-compose.yml up -d



       FOR NOTIFICATION SERVICE


1. git clone https://github.com/suvamsingh/notification.git

Launch producer api

2. cd notification 
3. go run main.go

Launch consumer

1. cd notification
2. go run worker/main.go

