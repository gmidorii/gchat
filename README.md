# gchat

## Overview
gchat is chat server based on websocket.  

## Install
```
go get github.com/midorigreen/gchat
```

## Usage
### API Document
| param | value               |
|-------|---------------------|
| room  | chat room name.     |
| name  | register user name. |

### Start-up
```sh
% ./gchat -h
Usage of ./gchat:
  -addr string
        server port (default ":8080")
```

### Client Sample
[wscat](https://github.com/websockets/wscat) use.
ex) mac os
```
brew install wscat
```

ex) connect
```
% wscat -c "wc://localhost:3000/chat?room=gchat-usage&name=midori"
```

```
connected (press CTRL+C to quit)
> Hello!!

< [room name]: user-name
< gchat-usage:midori
 Hello!!
```
