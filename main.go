package main
//TODO:Partial file

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

/*-----------------------------*/
// RSS Feeds structure
// RSS1.0
// Date format
//     RFC3339     = "2006-01-02T15:04:05Z07:00"
type Rss1 struct {
	Channel Rss1Channel `xml:"channel"`
	Item    []Rss1Item  `xml:"item"`
}

type Rss1Channel struct {
	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	Description string          `xml:"description"`
	Items       Rss1ChannelItem `xml:"items"`
	Date        string          `xml:"date"`
}

type Rss1ChannelItem struct {
	Seq Rss1ChannelItemList `xml:"Seq"`
}

type Rss1ChannelItemList struct {
	Li []string `xml:"li"`
}

type Rss1Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Creator     string `xml:"creator"`
	Date        string `xml:"date"`
}

// RSS2.0
// pubDate format
//    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
type Rss2 struct {
	Channel Rss2Channel `xml:"channel"`
}

type Rss2Channel struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Image       Rss2Image  `xml:"image"`
	PubDate     string     `xml:"pubDate"`
	Creator     string     `xml:"creator"`
	Item        []Rss2Item `xml:"item"`
}

type Rss2Image struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Url   string `xml:"url"` //logo url
}

type Rss2Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Author      string `xml;"author"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

//Atom
// Updated format
//     RFC3339     = "2006-01-02T15:04:05Z07:00"
type Atom struct {
	Id       string      `xml:"id"`
	Title    string      `xml:"title"`
	Subtitle string      `xml:"subtitle"`
	Updated  string      `xml:"updated"`
	Author   AtomAuthor  `xml:"author"`
	Logo     string      `xml:"logo"`
	Link     []string    `xml:"link"`
	Entry    []AtomEntry `xml:"entry"`
}

type AtomEntry struct {
	Id      string        `xml:"id"`
	Title   string        `xml:"title"`
	Link    AtomEntryLink `xml:"link"`
	Updated string        `xml:"updated"`
	Author  AtomAuthor    `xml:"author"`
	Content string        `xml:"content"`
	Summary string        `xml:"summary"`
}

type AtomEntryLink struct {
	Href string `xml:"href,attr"`		//get href value
}

type AtomAuthor struct {
	Name  string `xml:"name"`
	Url   string `xml:"url"`
	Email string `xml:"email"`
}
/*-----------------------------*/

/*-----------------------------*/
/* discord Webhook chatbot json*/

// nullable structure. null structure is not output, need pointer

// simple message example
/*
{
    "content": "simple bot."
}
*/

// embed example(hyperlink)
/*
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

// embed example(upload picture use multipart/form-data )
// using Attachment "Url" part
{
    "embed": {
        "image": {
            "url": "attachment://screenshot.png"
        }
    }
}

// advanced example
// multiple contents(content + embeds)
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

//a lot of field json image
// WARNING: this json is do not work
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
                    "value": "First Description Field Sample Value"
                    "inline": true
                },
                {
                    "name": "Second Description Field Sample Name",
                    "value": "Second Description Field Sample Value"
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
*/

/*-----------------------------*/
// Discord Webhook Structs
type DiscordWebhook struct {
	Content   string   `json:"content,omitempty"`    // One of content, file, embeds
	UserName  string   `json:"username,omitempty"`   // Not Required
	AvatarUrl string   `json:"avatar_url,omitempty"` // Not Required
	Tts       string   `json:"tts,omitempty"`        // Not Required
	File      string   `json:"file,omitempty"`       // One of content, file, embeds
	Embeds    []*Embed `json:"embeds,omitempty"`     // One of content, file, embeds
}

type Embed struct {
	Title       string     `json:"title,omitempty"`       // title of embed
	Type        string     `json:"type,omitempty"`        // type of ebmed (always `rich` for webhook embeds)
	Description string     `json:"description,omitempty"` // description of embed
	Url         string     `json:"url,omitempty"`         // url of embed
	Timestamp   string     `json:"timestamp,omitempty"`   // timestamp for embed content
	Color       int        `json:"color,omitempty"`       // color code of the embed (value is not HEX. value is integer. (R*256*256) * (G*256) * B = value  ex R:151 G237 B83 = 9956691（HEX is #97ED53)  )
	Footers     *Footer    `json:"footer,omitempty"`      // embed footer object footer infomation
	Images      *Image     `json:"image,omitempty"`       // embed image object image infomation
	Thumbnails  *Thumbnail `json:"thumbnail,omitempty"`   // embed thumbnail object thumbnail infomation
	Videos      *Video     `json:"video,omitempty"`       // embed video object video infomation
	Providers   *Provider  `json:"provider,omitempty"`    // embed provider object provider infomation
	Authors     *Author    `json:"author,omitempty"`      // embed author object author infomation
	Fields      []*Field   `json:"fields,omitempty"`      // array of embed field object field infomation
}

//for Embeds struct
type Thumbnail struct {
	Url      string `json:"url,omitempty"`       // source url of thumbnail (only supports http(s) and attachments)
	ProxyUrl string `json:"proxy_url,omitempty"` // a proxied url of the thumbnail
	Height   int    `json:"height,omitempty"`    // height of thumbnail
	Width    int    `json:"width,omitempty"`     // width of thumbnail
}

type Video struct {
	Url    string `json:"url,omitempty"`    // source url of video (only supports http(s) and attachments)
	Height int    `json:"height,omitempty"` // height of video
	Width  int    `json:"width,omitempty"`  // width of video
}

type Image struct {
	Url      string `json:"url,omitempty"`       // source url of image (only supports http(s) and attachments)
	ProxyUrl string `json:"proxy_url,omitempty"` // a proxied url of the image
	Height   int    `json:"height,omitempty"`    // height of image
	Width    int    `json:"width,omitempty"`     // width of image
}

type Provider struct {
	Name string `json:"name,omitempty"` // name of provider
	Url  string `json:"url,omitempty"`  // source url of provider
}

type Author struct {
	Name         string `json:"name,omitempty"`           // name of author
	Url          string `json:"url,omitempty"`            // url of author
	IconUrl      string `json:"icon_url,omitempty"`       // url of author icon (only supports http(s) and attachments)
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"` // aproxied url of author icon
}

type Footer struct {
	Text         string `json:"text,omitempty"`           // footer text
	IconUrl      string `json:"icon_url,omitempty"`       // url of footer icon (only supports http(s) and attachments)
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"` // a proxied url of footer icon
}

type Field struct {
	Name   string `json:"name,omitempty"`   // name of the field
	Value  string `json:"value,omitempty"`  // value of the field
	inline bool   `json:"inline,omitempty"` // whether or not this field should display inline
}

//type DiscordAttachment struct {
//	Id       string `json:"id"`        // attachment ID
//	Filename string `json:"filename"`  // name of file attached
//	Size     int    `json:"size"`      // size of file bytes
//	Url      string `json:"url"`       // source url of file
//	ProxyUrl string `json:"proxy_url"` // a proxiedurl of file
//	height   int    `json:"height"`    // height of file (if image)
//	width    int    `json:"width"`     // width of file (if image)
//}

// TODO:support change color
// color value define for DiscordWebhook
const (
	D_FATAL   = 16711680 // red
	D_WARNING = 16739584 // orange
	D_ALERT   = 16776960 // yellow
	D_SUCCESS = 65280    // light green
	D_INFO    = 65535    // cyan
)
/*-----------------------------*/

/*-----------------------------*/
//Slack Webhook json sample
/*
https://api.slack.com/docs/messages/builder?msg=%7B%22attachments%22%3A%5B%7B%22fallback%22%3A%22Required%20plain-text%20summary%20of%20the%20attachment.%22%2C%22color%22%3A%22%2336a64f%22%2C%22pretext%22%3A%22Optional%20text%20that%20appears%20above%20the%20attachment%20block%22%2C%22author_name%22%3A%22Bobby%20Tables%22%2C%22author_link%22%3A%22http%3A%2F%2Fflickr.com%2Fbobby%2F%22%2C%22author_icon%22%3A%22http%3A%2F%2Fflickr.com%2Ficons%2Fbobby.jpg%22%2C%22title%22%3A%22Slack%20API%20Documentation%22%2C%22title_link%22%3A%22https%3A%2F%2Fapi.slack.com%2F%22%2C%22text%22%3A%22Optional%20text%20that%20appears%20within%20the%20attachment%22%2C%22fields%22%3A%5B%7B%22title%22%3A%22Priority%22%2C%22value%22%3A%22High%22%2C%22short%22%3Afalse%7D%5D%2C%22image_url%22%3A%22http%3A%2F%2Fmy-website.com%2Fpath%2Fto%2Fimage.jpg%22%2C%22thumb_url%22%3A%22http%3A%2F%2Fexample.com%2Fpath%2Fto%2Fthumb.png%22%2C%22footer%22%3A%22Slack%20API%22%2C%22footer_icon%22%3A%22https%3A%2F%2Fplatform.slack-edge.com%2Fimg%2Fdefault_application_icon.png%22%2C%22ts%22%3A123456789%7D%5D%7D

// simple
{
     text: "Sample"
}

// a lot of fields sample
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
*/

/*-----------------------------*/
// Slack webhook strucuture
type SlackWebhook struct {
	Text        string             `json:"text"`
	UserName    string             `json:"username"`
	IconEmoji   string             `json:"icon_emoji"`
	IconUrl     string             `json:"icon_url"`
	Channel     string             `json:"channel"`
	Attachments []*SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Fallback   string                  `json:"fallback"`
	Color      string                  `json:"color"`   //HEX. (ex: #FFFFFF)
	Pretext    string                  `json:"pretext"` //text position.
	AuthorName string                  `json:"author_name"`
	AuthorLink string                  `json:"author_link"`
	AuthorIcon string                  `json:"author_icon"`
	Title      string                  `json:"title"`
	TitleLink  string                  `json:"title_link"`
	Text       string                  `json:"text"`
	Fields     []*SlackAttachmentField `json:"fields"`
	ImageUrl   string                  `json:"image_url"`
	ThumbUrl   string                  `json:"thumb_url"`
	Footer     string                  `json:"footer"`
	FooterIcon string                  `json:"footer_icon"`
	Ts         int                     `json:"ts"` //timestamp(number)
}

type SlackAttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"` //display inline or block
}

// TODO:support change color
// color value define for SlackWebhook
const (
	S_FATAL   = "#f44242" // red
	S_WARNING = "#f48641" // orange
	S_ALERT   = "#f4f141" // yellow
	S_SUCCESS = "#41f449" // light green
	S_INFO    = "#41eef4" // cyan
)
/*-----------------------------*/
// default date format define
const (
	YYYYMMDDHH24MISS = "2006/01/02 15:04:05"
)

// setting file path
const DEFAULT_SETTINGS = "./settings.json"
/*-----------------------------*/

/*-----------------------------*/
// local setting json file structure
type Config struct {
	Deliveries []*DeliveryList `json:"deliveries"`
	Feeds      []*FeedList     `json:"feeds"`
}

type DeliveryList struct {
	UserName string `json:"username"` //display user name
	Icon     string `json:"icon"`     //display user icon (slack is emoji, discord is avatar_url value)
	Url      string `json:"url"`      //webhook url
	Type     string `json:"type"`     //webhook type `slack` or `discord`
	Post     bool   `json:"post"`     //using webhook
}

type FeedList struct {
	Url       string `json:"url"`       //RSS feed url
	Timestamp string `json:"timestamp"` //latest readed timestamp (YYYY/MM/DD HH21:MI:DD)
	Read      bool   `json:"read"`      //using RSS feed url
}
/*-----------------------------*/

/*-----------------------------*/
// command lin args
var configFile string
/*-----------------------------*/

func main() {
	flag.StringVar(&configFile, "config", DEFAULT_SETTINGS, "using config file path.")
	flag.Parse()

	c := new(Config)
	c.readSettingsJson()
	t := time.Now().UTC().Format(YYYYMMDDHH24MISS)

	for hooks, _ := range c.Deliveries {
		if c.Deliveries[hooks].Post {
			for idx, _ := range c.Feeds {
				if c.Feeds[idx].Read {
					t, _ := time.Parse(YYYYMMDDHH24MISS, c.Feeds[idx].Timestamp)
					readFeed(c.Feeds[idx].Url, t, c.Deliveries[hooks].Url, c.Deliveries[hooks].Type, c.Deliveries[hooks].UserName, c.Deliveries[hooks].Icon)
				}
				c.Feeds[idx].Timestamp = t
			}
		}
	}
	c.writeSettingsJson()
}

// Read RSS Feed
func readFeed(url string, latestReaded time.Time, webhookUrl string, webhookType string, username string, icon string) {

	xmlStr := getXml(url)

	// case RSS1.0
	if strings.Contains(xmlStr, "<rdf:RDF") {
		data := new(Rss1)
		if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
			fmt.Println("XML Unmarshal error:", err)
			return
		}

		switch webhookType {
		case "discord":
			d := new(DiscordWebhook)
			d.parseRss1(data, latestReaded, username, icon)
			postWebhook(webhookUrl, webhookType, d.marshal())
		case "slack":
			s := new(SlackWebhook)
			s.parseRss1(data, latestReaded, username, icon)
			postWebhook(webhookUrl, webhookType, s.marshal())
		}
	}

	// case RSS2.0
	if strings.Contains(xmlStr, "<rss") {
		data := new(Rss2)
		if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
			fmt.Println("XML Unmarshal error:", err)
			return
		}

		switch webhookType {
		case "discord":
			d := new(DiscordWebhook)
			d.parseRss2(data, latestReaded, username, icon)
			postWebhook(webhookUrl, webhookType, d.marshal())
		case "slack":
			s := new(SlackWebhook)
			s.parseRss2(data, latestReaded, username, icon)
			postWebhook(webhookUrl, webhookType, s.marshal())
		}
	}

	// case Atom
	if strings.Contains(xmlStr, "<feed") {
		data := new(Atom)
		if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
			fmt.Println("XML Unmarshal error:", err)
			return
		}

		switch webhookType {
		case "discord":
			d := new(DiscordWebhook)
			d.parseAtom(data, latestReaded, username, icon)
			postWebhook(webhookUrl, webhookType, d.marshal())
		case "slack":
			s := new(SlackWebhook)
			s.parseAtom(data, latestReaded, username, icon)
			postWebhook(webhookUrl, webhookType, s.marshal())
		}
	}
}

// get xml document use url
func getXml(url string) (html string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(body)

	html = buf.String()
	return
}

// innerText html part delete
func purgeHTML(eval string) (ret string) {
	if e, _ := regexp.MatchString(`<(".*?"|'.*?'|[^'"])*?>`, eval); e {
		r := regexp.MustCompile(`<(".*?"|'.*?'|[^'"])*?>`)
		ret = r.ReplaceAllLiteralString(eval, "")
	} else {
		ret = eval
	}
	ret = strings.TrimSpace(strings.Replace(ret, "\"", "'", -1))

	return
}

// Parse RSS1.0 feed for DiscordWebhook structure
func (d *DiscordWebhook) parseRss1(data *Rss1, latestReaded time.Time, username string, icon string) {
	for idx, _ := range data.Item {
		thisEmbedPublished, _ := time.Parse(YYYYMMDDHH24MISS, dateFormat(data.Item[idx].Date))
		if latestReaded.Before(thisEmbedPublished) {
			if d.Content == "" {
				d.UserName = username
				d.AvatarUrl = icon
				d.Content = fmt.Sprintf("%s New publish! (from: %s)\n%s", dateFormat(data.Channel.Date), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Link))
			}
			emb := new(Embed)
			emb.Title = purgeHTML(data.Item[idx].Title)
			emb.Description = purgeHTML(data.Item[idx].Description)
			emb.Url = purgeHTML(data.Item[idx].Link)
			emb.Color = D_INFO

			foot := new(Footer)
			foot.Text = fmt.Sprintf("%s\n%s\n%s", dateFormat(data.Item[idx].Date), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Description))
			emb.Footers = foot

			d.Embeds = append(d.Embeds, emb)
		}
	}
}

// Parse RSS2.0 feed for DiscordWebhook structure
func (d *DiscordWebhook) parseRss2(data *Rss2, latestReaded time.Time, username string, icon string) {
	for idx, _ := range data.Channel.Item {
		thisEmbedPublished, _ := time.Parse(YYYYMMDDHH24MISS, dateFormat(data.Channel.Item[idx].PubDate))
		if latestReaded.Before(thisEmbedPublished) {
			if d.Content == "" {
				d.UserName = username
				d.AvatarUrl = icon
				d.Content = fmt.Sprintf("%s New publish! (from: %s)\n%s", dateFormat(data.Channel.PubDate), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Link))
			}
			emb := new(Embed)
			emb.Title = purgeHTML(data.Channel.Item[idx].Title)
			emb.Description = purgeHTML(data.Channel.Item[idx].Description)
			emb.Url = purgeHTML(data.Channel.Item[idx].Link)
			emb.Color = D_INFO

			foot := new(Footer)
			foot.Text = fmt.Sprintf("%s\n%s\n%s\n%s", dateFormat(data.Channel.Item[idx].PubDate), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Description), purgeHTML(data.Channel.Creator))
			emb.Footers = foot

			d.Embeds = append(d.Embeds, emb)
		}
	}
}

// Parse Atom feed for DiscordWebhook structure
func (d *DiscordWebhook) parseAtom(data *Atom, latestReaded time.Time, username string, icon string) {
	for idx, _ := range data.Entry {
		thisEmbedPublished, _ := time.Parse(YYYYMMDDHH24MISS, dateFormat(data.Entry[idx].Updated))
		if latestReaded.Before(thisEmbedPublished) {
			if d.Content == "" {
				d.UserName = username
				d.AvatarUrl = icon
				d.Content = fmt.Sprintf("%s New publish! (from: %s)\n%s", dateFormat(data.Updated), purgeHTML(data.Title), purgeHTML(data.Author.Url))
			}
			emb := new(Embed)
			emb.Title = purgeHTML(data.Entry[idx].Title)
			emb.Description = purgeHTML(data.Entry[idx].Content)
			emb.Url = purgeHTML(data.Entry[idx].Link.Href)
			emb.Color = D_INFO

			foot := new(Footer)
			foot.Text = fmt.Sprintf("%s\n%s\n%s", dateFormat(data.Entry[idx].Updated), purgeHTML(data.Title), purgeHTML(data.Subtitle))
			emb.Footers = foot

			d.Embeds = append(d.Embeds, emb)
		}
	}
}

// Parse RSS1.0 feed for SlackWebhook structure
func (s *SlackWebhook) parseRss1(data *Rss1, latestReaded time.Time, username string, icon string) {
	for idx, _ := range data.Item {
		thisEmbedPublished, _ := time.Parse(YYYYMMDDHH24MISS, dateFormat(data.Item[idx].Date))
		if latestReaded.Before(thisEmbedPublished) {
			if s.Text == "" {
				s.UserName = username
				s.IconEmoji = icon
				s.Text = fmt.Sprintf("%s New publish! (from: %s)\n%s", dateFormat(data.Channel.Date), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Link))
			}
			emb := new(SlackAttachment)
			emb.Title = purgeHTML(data.Item[idx].Title)
			emb.Text = purgeHTML(data.Item[idx].Description)
			emb.TitleLink = purgeHTML(data.Item[idx].Link)
			emb.Color = S_INFO
			emb.Footer = fmt.Sprintf("%s\n%s\n%s", dateFormat(data.Item[idx].Date), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Description))

			s.Attachments = append(s.Attachments, emb)
		}
	}
}

// Parse RSS2.0 feed for SlackWebhook structure
func (s *SlackWebhook) parseRss2(data *Rss2, latestReaded time.Time, username string, icon string) {
	for idx, _ := range data.Channel.Item {
		thisEmbedPublished, _ := time.Parse(YYYYMMDDHH24MISS, dateFormat(data.Channel.Item[idx].PubDate))
		if latestReaded.Before(thisEmbedPublished) {
			if s.Text == "" {
				s.UserName = username
				s.IconEmoji = icon
				s.Text = fmt.Sprintf("%s New publish! (from: %s)\n%s", dateFormat(data.Channel.PubDate), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Link))
			}
			emb := new(SlackAttachment)
			emb.Title = purgeHTML(data.Channel.Item[idx].Title)
			emb.Text = purgeHTML(data.Channel.Item[idx].Description)
			emb.TitleLink = purgeHTML(data.Channel.Item[idx].Link)
			emb.Color = S_INFO
			emb.Footer = fmt.Sprintf("%s\n%s\n%s\n%s", dateFormat(data.Channel.Item[idx].PubDate), purgeHTML(data.Channel.Title), purgeHTML(data.Channel.Description), purgeHTML(data.Channel.Creator))

			s.Attachments = append(s.Attachments, emb)
		}
	}
}

// Parse Atom feed for SlackWebhook structure
func (s *SlackWebhook) parseAtom(data *Atom, latestReaded time.Time, username string, icon string) {
	for idx, _ := range data.Entry {
		thisEmbedPublished, _ := time.Parse(YYYYMMDDHH24MISS, dateFormat(data.Entry[idx].Updated))
		if latestReaded.Before(thisEmbedPublished) {
			if s.Text == "" {
				s.UserName = username
				s.IconEmoji = icon
				s.Text = fmt.Sprintf("%s New publish! (from: %s)\n%s", dateFormat(data.Updated), purgeHTML(data.Title), purgeHTML(data.Author.Url))
			}
			emb := new(SlackAttachment)
			emb.Title = purgeHTML(data.Entry[idx].Title)
			emb.Text = purgeHTML(data.Entry[idx].Content)
			emb.TitleLink = purgeHTML(data.Entry[idx].Link.Href)
			emb.Color = S_INFO
			emb.Footer = fmt.Sprintf("%s\n%s\n%s", dateFormat(data.Entry[idx].Updated), purgeHTML(data.Title), purgeHTML(data.Subtitle))

			s.Attachments = append(s.Attachments, emb)
		}
	}
}

// publish date format method  ( out string YYYY/MM/DD hh24:mi:ss )
func dateFormat(str string) (ret string) {
	t := time.Now() // t init
	if _, err := time.Parse(time.UnixDate, str); err == nil {
		t, _ = time.Parse(time.UnixDate, str)
	}
	if _, err := time.Parse(time.RubyDate, str); err == nil {
		t, _ = time.Parse(time.RubyDate, str)
	}
	if _, err := time.Parse(time.RFC822, str); err == nil {
		t, _ = time.Parse(time.RFC822, str)
	}
	if _, err := time.Parse(time.RFC822Z, str); err == nil {
		t, _ = time.Parse(time.RFC822Z, str)
	}
	if _, err := time.Parse(time.RFC850, str); err == nil {
		t, _ = time.Parse(time.RFC850, str)
	}
	if _, err := time.Parse(time.RFC1123, str); err == nil {
		t, _ = time.Parse(time.RFC1123, str)
	}
	if _, err := time.Parse(time.RFC1123Z, str); err == nil {
		t, _ = time.Parse(time.RFC1123Z, str)
	}
	if _, err := time.Parse(time.RFC3339, str); err == nil {
		t, _ = time.Parse(time.RFC3339, str)
	}
	if _, err := time.Parse(time.RFC3339Nano, str); err == nil {
		t, _ = time.Parse(time.RFC3339Nano, str)
	}
	ret = t.Format(YYYYMMDDHH24MISS)
	return
}

//JSON Marshal for DiscordWebhook
func (d *DiscordWebhook) marshal() (jsonBytes []byte) {

	jsonBytes, err := json.Marshal(*d)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	//	out := new(bytes.Buffer)
	//	json.Indent(out, jsonBytes, "", "    ")
	//	fmt.Println(out.String())
	return
}

//JSON Marshal for SlackWebhook
func (s *SlackWebhook) marshal() (jsonBytes []byte) {

	jsonBytes, err := json.Marshal(*s)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	//	out := new(bytes.Buffer)
	//	json.Indent(out, jsonBytes, "", "    ")
	//	fmt.Println(out.String())
	return
}

// read setting json file
func (c *Config) readSettingsJson() {
	f, _ := filepath.Abs(configFile)
	r, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(r, &c); err != nil {
		panic(err)
	}
}

// write setting json file
func (c *Config) writeSettingsJson() {
	enc, _ := json.MarshalIndent(*c, "", "    ")

	if err := ioutil.WriteFile(configFile, enc, 0660); err != nil {
		panic(err)
	}
}

// Webhook Post Method(do not support File Upload)
func postWebhook(webhookUrl string, webhookType string, p []byte) {

	var payloadName string
	switch webhookType {
	case "slack":
		payloadName = "payload"
	case "discord":
		payloadName = "payload_json"
	}
	resp, _ := http.PostForm(
		webhookUrl,
		url.Values{payloadName: {string(p)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}

/*------------------------------------------------------------*/
/*----------------*/
/* test code part */
/*----------------*/

// read/write setting json file test
func testReadSettingJson() {
	f, _ := filepath.Abs("./settings.json")
	s, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	var c Config

	err = json.Unmarshal(s, &c)
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Deliveries[0].Url)
	fmt.Println(c.Deliveries[0].Post)
	fmt.Println(c.Deliveries[1].Url)
	fmt.Println(c.Deliveries[1].Post)
	fmt.Println(c.Feeds[0].Url)
	fmt.Println(c.Feeds[0].Timestamp)
	fmt.Println(c.Feeds[0].Read)
	fmt.Println(c.Feeds[1].Url)
	fmt.Println(c.Feeds[1].Timestamp)
	fmt.Println(c.Feeds[1].Read)

}

// TODO:Webhook Post Method FileUpload
func testPostWebhook(p bytes.Buffer) {
	resp, _ := http.PostForm(
		"",
		url.Values{"payload_json": {p.String()}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}
