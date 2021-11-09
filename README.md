## Covid data API for Geolocation co-ordiantes

### Problem Statement:
* Build an API that fetches the number of Covid-19 cases in each state and in India
and persist it in MongoDB.
* Using this data, build an API that takes the user's GPS coordinates as input and
returns the total number of Covid-19 cases in the user's state and in
India(assume India specific coordinates only) and the last update time of data.
* [Bonus Task] Deploy the app to Heroku using the Add-ons for the DBs used.
* [Bonus Task] Add a redis caching layer which caches data for 30 mins.

### Specifications:
* Language: Golang
* Web Framework: Echo
* Databases: Mongo and Redis(optional)
* Set Up Swagger Docs for the APIs
* Integrate go modules for dependency management.
* Organize the code into separate packages using any project layout pattern.
* Use any of the publicly available APIs for fetching Covid-19 data and reverse
geocoding.

