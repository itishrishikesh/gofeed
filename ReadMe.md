# GoFeed: A RSS Feed Aggregator Backend

GoFeed is a simple and powerful backend for aggregating RSS feeds. It allows you to register users, follow RSS feeds, and read them in one place.

## Getting Started

The easiest way to get GoFeed up and running is to use `docker-compose up` command in the project directory. This will create a container for GoFeed server and a container for Postgres database.

## API Documentation

GoFeed provides the following endpoints for interacting with the backend:

- **User Management**
  - Create User (POST: `v1/user`)
  - Retrieve User (GET: `v1/user`)
- **Token Generation**
  - Generate Token (POST: `v1/token`)
- **Feed Management**
  - Create Feed (POST: `v1/feeds`)
  - Retrieve Feed (GET: `v1/feeds`)
- **Feed Follow Management**
  - Create Feed Follow (POST: `v1/feedfollow`)
  - Retrieve Feed Follow (GET: `v1/feedfollow`)
  - Delete Feed Follow (DELETE: `/feedfollow/{feedFollowId}`)
- **Post Management**
  - Retrieve Posts (GET: `v1/posts`)
- **Health Check**
  - Health (GET: `v1/healthz`)
- **Error Simulation**
  - Error (GET: `v1/error`)

## Authentication

All endpoints except "Create User", "Health", "Error", and "Token" require authentication.

GoFeed supports two authentication mechanisms:

- **API Key based authentication**: Every user is assigned an API key at the time of creation, which can be used to make subsequent API requests.
- **JWT Token based authentication**: A JWT token can also be obtained by providing username and password to the `v1/token` endpoint, and then used for authentication.

You can find examples of how to call the endpoints in the `insomnia_collection.json` file in the root of the project. This file can be imported into Insomnia or Postman.

## Motivation

This project was created as a learning exercise for Go programming language and various technologies related to it. However, the project is fully functional and can be used as a standalone backend for RSS feed aggregation.

## Contribution

Although the project is fully functional, there is a lot of room for improvement and new features. If you are interested in learning Go and contributing to an open source project, please feel free to create an issue and start working on it. I will do my best to support you by answering any questions you may have.