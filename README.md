# bridge
[![Go Report Card](https://goreportcard.com/badge/github.com/Always-Newbie/bridge)](https://goreportcard.com/report/github.com/Always-Newbie/bridge)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Always-Newbie/bridge/blob/master/LICENSE)   
Bridge program between RSS and Slack or Discord webhook written in a pure Golang 

### About

"bridge" is RSS feeds reader.  
This app written a pure Golang.    
And RSS feeds convert into json, and execute HTTP POST your team [Slack](https://slack.com/) or [Discord](https://discordapp.com/) webhook.

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
  --・`timestamp` is latest read UTC+0000 timestamp (YYYY/MM/DD HH24:MI:SS) *AUTO UPDATE!  
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

2.Please execute command
```
//go run main.go [option:-config=(setting file path. default is "./settings.json")] 

$ ./bridge [option:-config=(setting file path. default is "./settings.json")]
```

3.Please check your Slack or Discord channel

### Affinity Usage
・cron  

### Json image for Discord Webhook

・simple message example
```json
{
    "content": "simple bot."
}
```

・embed example(hyperlink)
```json
{
    "embeds": [
        {
            "title": "bot test",
            "url": "https://www.example.com/",
            "description": "bot bot bot!",
            "type": "link",
            "thumbnail":{
                "url": ""
            }
        }
    ]
}
```

・advanced example multiple contents(content + embeds)
```json
{
    "username": "example username",
    "avatar_url": "https://xxxxxx.png",
    "content": "content sample",
    "embeds": [
        {
            "title" : "embed title",
            "url" : "https://www.example.com/",
            "author": {
                "name": "embed author name",
                "url": "https://www.example.com/",
                "icon_url": "https://xxxxxxxx.png",
            },
            "description": "[masking url text sample](https://www.google.com/)",
            "color": 15258703,
            "fields": [
                {
                    "name": "embed field first",
                    "value": "embed field first value"
                },
                {
                    "name": "embed field second",
                    "value": "embed field second value"
                }
            ],
            "footer": {
                "text": "footer text",
                "icon_url": "https://xxxxxxxx.png"
            }
        }
    ]
}
```

・a lot of field json image  
WARNING: this json is do not work  
```json
{
    "username": "example username",
    "avatar_url": "https://www.example.com/avatar_url.png",
    "content": "Content Sample Text",
    "tts": false,
    "file":  "",
    "embeds": [
        {
            "title" : "Sample Embed Title",
            "url" : "https://www.example.com",
            "description": "[Click here to see this text](https://www.google.com/embed_description_url)",
            "color": 15258703,
            "author": {
                "name": "Sample Author Name",
                "url": "https://www.example.com/author_url",
                "icon_url": "https://www.example.com/author_icon_url.png",
                "description" : "Author Description Text Example"
            },
            "thumbnail": {
                "url": "https://www.example.com/thumbnail_url.png",
            },
            "image": {
                "url": "https://www.example.com/image_url.png"
            },
            "video": {
                "url": "https://www.example.com/video_url",
            },
            "footer": {
                "text": "Footer Sample Text",
                "icon_url": "https://www.example.com/footer_icon.png"
            },
            "provider": {
                "name": "Sample Provider Name",
                "url": "http://www.example.com/provider_url"
            },
            "fields": [
                {
                    "name": "First Description Field Sample Name",
                    "value": "First Description Field Sample Value",
                    "inline": true
                },
                {
                    "name": "Second Description Field Sample Name",
                    "value": "Second Description Field Sample Value",
                    "inline": false
                }
            ],
            "footer": {
                "text": "Footer Sample Text",
                "icon_url": "https://www.example.com/footer_icon.png"
            }
        }
    ]
}
```

### Json image for Slack Webhook
[Slack Webhook json sample](https://api.slack.com/docs/messages/builder?msg=%7B%22attachments%22%3A%5B%7B%22fallback%22%3A%22Required%20plain-text%20summary%20of%20the%20attachment.%22%2C%22color%22%3A%22%2336a64f%22%2C%22pretext%22%3A%22Optional%20text%20that%20appears%20above%20the%20attachment%20block%22%2C%22author_name%22%3A%22Bobby%20Tables%22%2C%22author_link%22%3A%22http%3A%2F%2Fflickr.com%2Fbobby%2F%22%2C%22author_icon%22%3A%22http%3A%2F%2Fflickr.com%2Ficons%2Fbobby.jpg%22%2C%22title%22%3A%22Slack%20API%20Documentation%22%2C%22title_link%22%3A%22https%3A%2F%2Fapi.slack.com%2F%22%2C%22text%22%3A%22Optional%20text%20that%20appears%20within%20the%20attachment%22%2C%22fields%22%3A%5B%7B%22title%22%3A%22Priority%22%2C%22value%22%3A%22High%22%2C%22short%22%3Afalse%7D%5D%2C%22image_url%22%3A%22http%3A%2F%2Fmy-website.com%2Fpath%2Fto%2Fimage.jpg%22%2C%22thumb_url%22%3A%22http%3A%2F%2Fexample.com%2Fpath%2Fto%2Fthumb.png%22%2C%22footer%22%3A%22Slack%20API%22%2C%22footer_icon%22%3A%22https%3A%2F%2Fplatform.slack-edge.com%2Fimg%2Fdefault_application_icon.png%22%2C%22ts%22%3A123456789%7D%5D%7D)

・simple
```json
{
     "text": "Sample"
}
```

・a lot of fields sample
```json
{
    "attachments": [
        {
            "fallback": "Required plain-text summary of the attachment.",
            "color": "#36a64f",
            "pretext": "Optional text that appears above the attachment block",
            "author_name": "Bobby Tables",
            "author_link": "http://flickr.com/bobby/",
            "author_icon": "http://flickr.com/icons/bobby.jpg",
            "title": "Slack API Documentation",
            "title_link": "https://api.slack.com/",
            "text": "Optional text that appears within the attachment",
            "fields": [
                {
                    "title": "Priority",
                    "value": "High",
                    "short": false
                }
            ],
            "image_url": "http://my-website.com/path/to/image.jpg",
            "thumb_url": "http://example.com/path/to/thumb.png",
            "footer": "Slack API",
            "footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
            "ts": 123456789
        }
    ]
}
```



Thanks for Reading :)
