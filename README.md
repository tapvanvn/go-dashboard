# go-dashboard
A simple human friendly dashboard

# Require Environment Variable
- CONNECT_STRING_WSPUBSUB : wspubsub endpoint serving dashboard's event . example: ```"ws://192.168.1.8:80/ws"```
- CONNECT_STRING_DOCUMENTDB : document base database connect string. Base on which godbengine supported.
- DOCUMENTDB_DATABASE : database name                            
- PORT : port to serve on