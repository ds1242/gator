# Blog Aggregator (Gator)
A cli tool that allows users to scrape RSS feeds and store information from those feeds to be viewed.
Links are provided to each post so that users can view the entire story.  It can be set to check automatically
over a user specified amount of time.

## Quick Start
Go Version 1.23+ required [Go Homeplage](https://go.dev/)
Postgres Version 15+ required [Postgres Homepage](https://www.postgresql.org/)

Config file instructions

## Usage

Once you have downloaded the required tools.  The following commands will let you run the programe

Register a new user:
`go run . register (name)`

Login a user
`go run . login (name)`

Follow a feed
`go run . addfeed "(feed name)" "https://feedname.com"`

Unfollow a feed
`go run . unfollow "https://feedname.com"`

Check what feeds the current logged in user is following
`go run . following`

Delete all users
`go run . reset`

Browse saved posts
`go run . browse (a number of posts)`
(defaults to 2)

Fetch all feeds based on a time frame
___
Enter any time value such as 1s 1min 1hr etc...

`go run . agg (time)`


Fetch all feeds
`go run . feeds`

List all saved users
`go run . users`

## Contributing
Feel free to clone and make it your own.  Open a PR if you want to make a contribution back.  

## Future Enhancements
I would like to redo some of the browse so the posts are only listed by posts that the current logged in user is following and not just all posts.  

