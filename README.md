# RSS Aggregator API 🚀
A backend service built with Go that allows users to subscribe to RSS feeds and aggregate posts from multiple sources into a personalized feed. The service features a background scraper that concurrently fetches new posts on a regular schedule.

## Features
  - User Management: Simple user creation and authentication via auto-generated API keys.

  - Feed Management: Authenticated users can add new RSS feeds to the system.

  - Feed Subscriptions: Users can follow and unfollow feeds to customize their content.

  - Concurrent Scraping: A background worker process constantly fetches the latest posts from all tracked feeds.

  - Personalized Feeds: Users can retrieve a list of the most recent posts from only the feeds they follow.

## Technology Stack 🔧
  - **Go**: The core language for the backend server.

  - **MySQL**: SQL database for storing all data.

  - **go-chi/chi**: For lightweight, powerful HTTP routing and middleware.

  - **sqlc**: For generating type-safe Go code from raw SQL queries.

  - **goose**: For managing database schema migrations.

  - **joho/godotenv**: For loading environment variables from a .env file.

  - **google/uuid**: For generating unique IDs.

## API Endpoints
The service exposes the following RESTful API endpoints. All routes are prefixed with /v1.

| Method   | Endpoint                       | Auth?  | Description                              |
| :------- | :----------------------------- | :------| :--------------------------------------- |
| `GET`    | `/healthz`                     |   No   | TO check the service is live.            |
| `GET`    | `/err`                         |   No   | To check the error handling.             |
| `GET`    | `/users`                       |   Yes  | Gets the user's info.                    |
| `POST`   | `/users`                       |   No   | Creates a new user account.              |
| `GET`    | `/feeds`                       |   No   | Retrieves all feeds in the system.       |
| `POST`   | `/feeds`                       |   Yes  | Adds a new RSS feed.                     |
| `DELETE` | `/feeds/{feedID}`              |   Yes  | Deletes a feed from the system.          |
| `GET`    | `/feed_follows`                |   Yes  | Retrieves all feeds the user follows.    |
| `POST`   | `/feed_follows`                |   Yes  | Follows a feed for user.                 |
| `DELETE` | `/feed_follows/{feedFollowID}` |   Yes  | Unfollows a specific feed.               |
| `GET`    | `/posts`                       |   Yes  | Gets the most recent posts for the user. |


### Prerequisites
Go (version 1.22 or newer)

A running SQL database (e.g., PostgreSQL, MySQL)

[sqlc](https://github.com/sqlc-dev/sqlc) and [goose](https://github.com/pressly/goose) installed

### Installation

- Clone the repo
```Bash
git clone https://github.com/a-rezk/rssagg
cd rssagg
```

- Install Dependencies
```Bash
go mod tidy
go mod vendor
```

- Configure Environment
```Bash
echo 'PORT="8080" 
DB_URL="root:YourPassword@tcp(127.0.0.1:3306)/rssagg"' > .env
```

- Run Database Migrations then Generate Database Code
```Bash
# Replace "mysql" with "postgres" if needed
goose mysql "DB_URL" up && sqlc generate
```

- Run the Server
```Bash
go build && ./rssagg
# Alternatively, for development, you can just run:
go run main.go
```

*The server should now be running on the port you specified in your .env file*
