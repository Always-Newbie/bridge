# bridge
[![Go Report Card](https://goreportcard.com/badge/github.com/Always-Newbie/bridge)](https://goreportcard.com/report/github.com/Always-Newbie/bridge)  
Bridge program between RSS and webhook written in a pure golang

### About

"bridge" is RSS feeds reader.  
This app is written a pure Golang.    
And RSS feeds convert into json, and execute HTTP POST your team (Discord or Slack)webhook.

### Support webhooks

・Slack  
・Discord

### Install  

```
go get -u github.com/Always-Newbie/bridge　　
```

### How to use
0.Please "go build"　　

```
//cd $GOPATH/src/github.com/Always-Newbie/bridge

go build
```

1.Please preparation for settings.json  
・file encoding is "UTF-8".  
・Grant read / write permission.  
・"settings.json" path is binary same directory (default settings file)  
・`deliveries` is webhooks setting.  
  --・`url` is webhook url.  
  --・`username` is post username.  
  --・`icon` is post avatar icon.  
  --・`type` is webhook type. ("discord" or "slack")  
  --・`post` is can do post flag.  
  ・`feeds` is RSS feeds setting.  
  --・`url` is RSS feed url.  
  --・`timestamp` is latest readed UTC+0000 timestamp (YYYY/MM/DD HH24:MI:SS)  
  --・`read` is can do read flag.

```json
{
    "deliveries":[
        {
            "url": "https://discordapp.com/api/webhooks/xxxxxxxxxxxxxx",
            "username": "discord bot",
            "icon": "",
            "type": "discord",
            "post": true
        },
        {
            "url": "https://hooks.slack.com/services/yyyyyyyyyyyyyyy",
            "username": "slack bot",
            "icon": ":crossed_swords:",
            "type": "slack",
            "post": true
        }
    ],
    "feeds":[
        {
            "url": "http://www.nba.com/rss/nba_rss.xml",
            "timestamp": "2017/03/03 16:12:12",
            "read": true
        }
    ]
}
```

2.Execute command
```
$ ./bridge [option:-config=(setting file path. default is "./settings.json")]
```

3.Please check your Slack or Discord channel

### Affinity Usage
・cron  

### License
MIT license.  


Thanks for Reading :)
