Slothful Soda
=============

A simple app in the GGAP Stack (Golang, Gorilla, Angular, Postgres).  Check it out because it's about sloths. And because the GGAP stack is cool. But mostly because sloths.

###What is Slothful Soda?

With all the people playing Ultimate in Boston, a group of sloths living at Cambridge's Fresh Pond decided to make a beverage for players who wanted to re-hydrate after the game slothfully.  So they invented Slothful Soda, a soda made from hibiscus flowers for a delicious fruity-minty flavor both humans and sloths will love.  The sloths fly to different sports fields in the Boston Area to sell the soda using sloth-sized quadcopters so they made a website using the GGAP Stack (Golang, Gorilla, Angular, Postgres) to tell you what fields they land their quadcopters on.

###How do I use this web app?

####What you will need is:
* Go
* A Postgres database
* A Google Maps API Key

####How to set it up:
* `go get` the following packages:
  * `github.com/coopernurse/gorp`
  * `github.com/lib/pq`
  * `github.com/gorilla/mux`
* In Line 12 of `views/index.html`, replace `YOUR_GOOGLE_MAPS_API_KEY` with your Google Maps API Key
* In Line 26 of `initDB.go`, replace `postgres://yourPostgresDatabaseLogin` with your Postgres database login information
* Run `go install`
* Start the Go program and go to `localhost:1123`

###Where do I get Slothful Soda?
Slothful Soda is not a real product. But if someone actually invents it, let me know so I can buy some!

##NOTE: Don't tell friends who are likely to believe satire news articles like the ones on The Onion about this repository.  Sloths don't live on the Fresh Pond and don't have quadcopters but everyone knows if sloths figure out how to use quadcopters it means the Apocalypse is here and we really don't need any more people theorizing about when the Apocalypse is.
