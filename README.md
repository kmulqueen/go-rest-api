# Go REST API - Event Booking System

A comprehensive REST API built with Go for managing event bookings. This project serves as a practical learning exercise for Go development while creating a reusable foundation for future event management applications.

## Table of Contents

- [Go REST API - Event Booking System](#go-rest-api---event-booking-system)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Running the API](#running-the-api)
  - [API Documentation](#api-documentation)
    - [Authentication](#authentication)
    - [Events Endpoints](#events-endpoints)
    - [User Endpoints](#user-endpoints)

## Features

- **Event Management**: Create, read, update, and delete events
- **User Authentication**: JWT-based authentication system
- **Event Registration**: Users can register for and cancel event bookings
- **Authorization**: Role-based access control for event operations
- **RESTful Design**: Clean and intuitive API endpoints

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-rest-api.git
   cd go-rest-api
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

### Running the API

```bash
go run main.go
```

The API will start on `http://localhost:8080` by default.

## API Documentation

### Authentication

Most endpoints require authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

### Events Endpoints

| Method   | Endpoint       | Description              | Auth Required | Notes              |
| -------- | -------------- | ------------------------ | ------------- | ------------------ |
| `GET`    | `/events`      | Get all available events | No            | Public endpoint    |
| `GET`    | `/events/{id}` | Get event by ID          | No            | Public endpoint    |
| `POST`   | `/events`      | Create a new event       | Yes           | Event creator only |
| `PUT`    | `/events/{id}` | Update an event          | Yes           | Event creator only |
| `DELETE` | `/events/{id}` | Delete an event          | Yes           | Event creator only |

### User Endpoints

| Method   | Endpoint                | Description               | Auth Required | Notes             |
| -------- | ----------------------- | ------------------------- | ------------- | ----------------- |
| `POST`   | `/signup`               | Create a new user account | No            | Returns user data |
| `POST`   | `/login`                | Authenticate user         | No            | Returns JWT token |
| `POST`   | `/events/{id}/register` | Register for an event     | Yes           | User registration |
| `DELETE` | `/events/{id}/register` | Cancel event registration | Yes           | User cancellation |
