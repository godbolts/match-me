# This is the match me project

## Required

- **Docker**: Required for running the PostgreSQL database in a container. Tested with version 27.1.2. 
(Docker Compose is included with Docker, so no separate installation is needed).
- **Make**: For running predefined commands like `make start`, `make stop-db`, etc.
- **Node.js and NPM**: Required for the frontend. Tested with node version 23.1.0 & NPM version 10.9.0.
- **Go**: Required for the backend. Tested with version 1.23.3.
- **Concurrently**: A development dependency for running multiple commands (frontend and backend) simultaneously.


## Setting up the project

1. Clone the repository: 
```git clone https://gitea.kood.tech/karl-hendrikkahn/match-me.git```.
2. Install backend dependencies (for Go):
```cd backend```,
```go mod tidy```.
3. Install frontend dependencies:
```cd frontend```,
```npm install```.
4. Add a JWT secret key in the backend .env file ```JWT_SECRET = "a5b39010c7r7894jf8d8n98a83750t4gs6d5h54aq903831085ja1s5d6df4"```
5. Run ```docker-compose build``` in project directory.
6. Start the database, backend, and frontend servers all at the same time by running ```make start``` in the project root directory. (Run once, if exits with error then wait a few seconds, then run again).
7. Close the backend and frontend servers by pressing ```Ctrl+C``` in the terminal.
8. Close the database server with ```make stop-db```.

#### Other helpful commands
Check if the docker container is running with ```docker ps```.<br>
Start only the database server by running ```make start-db``` in the project root directory.<br>
Reset the database with ```make reset-db``` in the project root directory.<br>

## Demo bots.

By default, the backend server will be running on localhost:4000<br>
To create 100 demo users, send a request localhost:4000/spawn/bots (backend server) and wait for the request to finish.<br>

The demo users will be named as numbers. Example: 1@1.com will be user1 and the password will be 1. starting from user 0@0.com<br>

The amount of demo users created can be modified from the backend backend/db/const.go file by modifying the "DEMO_USER_COUNT" constant. <br>
As the demo accounts are created using the same paths as creating a regular user, <br>
with all the matching scores calculated for each connection, adding a lot of demo users at once can take a while.

## API

To use the API endpoints, you must first create a regular profile.

once you have created your profile, you can now access the API endpoints by logging in.

To login with a default setup

http://localhost:4000/login/api?email=1@1.com&password=1

This will return a string, this is you Bearer Token.

To access endpoints that require authentication you must include the Bearer token with any requests, <br>
including the Bearer Token in the authorization header.

**For testing purposes it is suggested to use the free postman app more information on https://web.postman.co/home**


## API Documentation

### Get All Connections
**Endpoint:** `GET /connections`

**Description:** Returns all connections for the user.

**Response:**
```json
[
    "9a5721fc-3499-4b8f-85cb-0dce0480daf6",
    "6e5e70f1-7931-40bb-b54f-374f63c72512",
    "db863918-54f0-4496-8a7e-599252ffc9ff",
    ... ,
    "ea115f4a-d64c-43fc-b09b-35e5649aea22",
    "369b46cb-20ef-44d9-965a-85ffab4ccd26",
    "4fe97a4d-ee06-4c25-a855-08a42c79b1cc"
]
```

### Get All Recommendations
**Endpoint:** `GET /recommendations`

**Description:** Returns all recommendations. It returns the same value as all connections if no filters are applied.

**Response:**
```json
[
    "ef12a23b-21e0-4f61-bfcc-5d4e39213cce",
    "ea85dc93-2fc8-4f7b-bffe-92c77ed6d7ec",
    ... ,
    "2c434578-9bd4-49bb-a264-d399d07b5c75",
    "c4c40123-d3f0-40cc-ac90-009f86109636"
]
```

### Get User Profile by ID
**Endpoint:** `GET /users/{id}/profile`

**Description:** Returns the user's profile data for the user with the included ID.

**Response:**
```json
{
    "id": 0,
    "uuid": "232e45c5-5d19-4bc0-95c8-46c32a0bad52",
    "username": "User11",
    "about_me": "I am a user 11",
    "profile_picture": "bot11.png",
    "created_at": "2025-01-19T08:07:23.124022Z",
    "birthdate": "1999-01-01T00:00:00Z"
}
```

### Get User Data by ID
**Endpoint:** `GET /users/{id}`

**Description:** Returns the user's data for the user with the included ID.

**Response:**
```json
{
    "id": "2c434578-9bd4-49bb-a264-d399d07b5c75",
    "uuid": "",
    "email": "1@1.com",
    "password_hash": "$2a$10$xlr4VDCDOMe1QSXmIijE8.y.8m.O26s3FnO1aHENqPE9JwtTTgUai",
    "created_at": "0001-01-01T00:00:00Z",
    "user_city": "",
    "user_nation": "",
    "user_region": "",
    "latitude": 59.378025,
    "longitude": 27.728493,
    "register_location": "",
    "brows_location": "",
    "is_online": false
}
```

### Get Logged-in User Profile
**Endpoint:** `GET /me/profile`

**Description:** Returns the user's profile data for the logged-in user (the data used to power recommendations).

**Response:**
```json
{
    "username": "User1",
    "email": "1@1.com",
    "created_at": "2025-01-19T08:07:22.532912Z",
    "is_online": false,
    "user_city": "Tartu",
    "user_nation": "Estonia",
    "user_region": "Tartu County",
    "about_me": "I am a user 1",
    "birthdate": "1999-01-01T00:00:00Z",


    "age": "25",


    "profile_picture": "bot1.png"
}
```

### Get Logged-in User Biographical Data
**Endpoint:** `GET /me/bio`

**Description:** Returns the user's biographical data for the logged-in user (the data used to power recommendations).

**Response Example:**
```json
{
    "Communication": [
        "voice chat "
    ],
    "Distance": [
        "100-500 km"
    ],
    "Genre": [
        "FPS",
        "Simulation"
    ],
    "Goals": [
        "Ranking",
        "Socialize",
        "For laughs"
    ],
    "Language": [
        "Estonian"
    ],
    "Platform": [
        "X-box"
    ],
    "Play Style": [
        "Enthusiast",
        "AFK"
    ],
    "Session": [
        "I can go all day, every day."
    ],
    "Vibe": [
        "Chill",
        "Laid-back"
    ]
}
```

### Get User Biographical Data by ID
**Endpoint:** `GET /users/{id}/bio`

**Description:** Returns the user's biographical data for the user with the included ID (the data used to power recommendations).

**Response:**
```json
{
    "Communication": [
        "voice chat "
    ],
    "Distance": [
        "100-500 km"
    ],
    "Genre": [
        "FPS",
        "Simulation"
    ],
    "Goals": [
        "Ranking",
        "Socialize",
        "For laughs"
    ],
    "Language": [
        "Estonian"
    ],
    "Platform": [
        "X-box"
    ],
    "Play Style": [
        "Enthusiast",
        "AFK"
    ],
    "Session": [
        "I can go all day, every day."
    ],
    "Vibe": [
        "Chill",
        "Laid-back"
    ]
}