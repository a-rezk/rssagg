RSS Aggregator API 🚀
A backend service built with Go that allows users to subscribe to RSS feeds and aggregate posts from multiple sources into a personalized feed. The service features a background scraper that concurrently fetches new posts on a regular schedule.

Features
User Management: Simple user creation and authentication via auto-generated API keys.

Feed Management: Authenticated users can add new RSS feeds to the system.

Feed Subscriptions: Users can follow and unfollow feeds to customize their content.

Concurrent Scraping: A background worker process constantly fetches the latest posts from all tracked feeds.

Personalized Feeds: Users can retrieve a list of the most recent posts from only the feeds they follow.

Technology Stack 🔧
Go: The core language for the backend server.

MySQL: SQL database for storing all data.

go-chi/chi: For lightweight, powerful HTTP routing and middleware.

sqlc: For generating type-safe Go code from raw SQL queries.

goose: For managing database schema migrations.

joho/godotenv: For loading environment variables from a .env file.

google/uuid: For generating unique IDs.

API Endpoints
The service exposes the following RESTful API endpoints. All routes are prefixed with /v1.

Method

Endpoint

Requires Auth?

Description

POST

/users

No

Creates a new user account.

GET

/users

Yes

Retrieves the authenticated user's info.

POST

/feeds

Yes

Adds a new RSS feed to the system.

GET

/feeds

No

Retrieves all feeds in the system.

POST

/feed_follows

Yes

Follows a feed for the authenticated user.

GET

/feed_follows

Yes

Retrieves all feeds the user follows.

DELETE

/feed_follows/{feedFollowID}

Yes

Unfollows a specific feed.

GET

/posts

Yes

Gets the most recent posts for the user.

GET

/healthz

No

Readiness check for the service.


Export to Sheets
Authentication: Authenticated endpoints require an Authorization header with the format ApiKey YOUR_API_KEY.

🚀 Getting Started
To get a local copy up and running, follow these simple steps.

Prerequisites
Go (version 1.22 or newer)

A running SQL database (e.g., PostgreSQL, MariaDB)

sqlc and goose installed

Bash

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
Installation
Clone the repo

Bash

git clone https://github.com/a-rezk/rssagg
cd rssagg
Create your environment file
Create a .env file in the root of the project by copying the example.

Bash

cp .env.example .env
Then, edit the .env file and add your database connection string and a port number.

Install dependencies

Bash

go mod tidy
Generate SQL-to-Go code

Bash

sqlc generate
Run the database migrations
Replace "your_db_connection_string" with the value from your .env file.

Bash

goose -dir ./sql/schema mysql "your_db_connection_string" up
Run the server

Bash

go run main.go
The server should now be running on the port you specified.

Environment Variables
The following environment variables are required to run the application. Create a .env file in the project root.

Variable

Description

Example

PORT

The port for the web server to run on.

8080

DATABASE_URL

The connection string for your SQL database.

root:YourPassword@tcp(localhost:3306)/rssagg?parseTime=true

`parseTime=true` For specific instruction for the Go MySQL driver. It tells the driver, "When you see a TIMESTAMP or DATETIME column, please parse it into a standard Go time.Time object for me.

Export to Sheets
📄 License
Distributed under the MIT License. See LICENSE for more information.