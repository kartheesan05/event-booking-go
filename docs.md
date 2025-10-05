## Event Booking API Documentation

### Base URL

- **Local**: `http://localhost:8080`

### Authentication

- **Login** returns a JWT in the `token` field.
- For authenticated endpoints, send the token as-is in the `Authorization` header (no `Bearer ` prefix).
  - Example: `Authorization: <jwt-token>`

### Data Models

- **Event**
  - `id`: number
  - `name`: string (required)
  - `description`: string (required)
  - `location`: string (required)
  - `dateTime`: string (ISO 8601 / RFC3339)
  - `userId`: number (owner)

### Endpoints

#### Signup

- **POST** `/signup`
- **Auth**: Not required
- **Request (application/json)**
```json
{
  "email": "user@example.com",
  "password": "strongpassword"
}
```
- **Responses**
  - 201 Created: `{ "message": "User creates succesfully." }`
  - 400 Bad Request: `{ "message": "Bad Request." }`
  - 500 Server Error: `{ "message": "Could not save user." }`
- **curl**
```bash
curl -X POST http://localhost:8080/signup \
  -H 'Content-Type: application/json' \
  -d '{"email":"user@example.com","password":"strongpassword"}'
```

#### Login

- **POST** `/login`
- **Auth**: Not required
- **Request (application/json)**
```json
{
  "email": "user@example.com",
  "password": "strongpassword"
}
```
- **Responses**
  - 200 OK: `{ "message": "Login Successful", "token": "<jwt>" }`
  - 400 Bad Request: `{ "message": "Bad Request." }`
  - 401 Unauthorized: `{ "message": "Credentials Invalid" }` or `{ "message": "Credential Invalid" }`
  - 500 Server Error: `{ "message": "Server Error. Try again later" }`
- **curl**
```bash
curl -X POST http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"user@example.com","password":"strongpassword"}'
```

#### List Events

- **GET** `/events`
- **Auth**: Not required
- **Response (200 OK)**: `Event[]`
- **curl**
```bash
curl http://localhost:8080/events
```

#### Get Event by ID

- **GET** `/events/:id`
- **Auth**: Not required
- **Path Params**: `id` (number)
- **Responses**
  - 200 OK: `Event`
  - 400 Bad Request: `{ "message": "Bad Request" }`
  - 500 Server Error: `{ "message": "Could not fetch Event. Try again Later." }`
- **curl**
```bash
curl http://localhost:8080/events/1
```

#### Create Event

- **POST** `/events`
- **Auth**: Required (`Authorization: <jwt-token>`, token from Login)
- **Request (application/json)**
```json
{
  "name": "My Conference",
  "description": "A great event",
  "location": "Chennai",
  "dateTime": "2026-01-01T18:00:00Z"
}
```
- The `userId` is taken from the authenticated user; do not send it in the body.
- **Responses**
  - 201 Created: `{ "message": "Event Successfully Created.", "event": Event }`
  - 400 Bad Request: `{ "message": "Could Not Parse Request Data." }`
  - 500 Server Error: `{ "message": "Could not create Event. Try again Later." }`
- **curl**
```bash
TOKEN="<jwt-token>"
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{"name":"My Conference","description":"A great event","location":"Chennai","dateTime":"2026-01-01T18:00:00Z"}'
```

#### Update Event

- **PUT** `/events/:id`
- **Auth**: Required; only the event owner may update
- **Path Params**: `id` (number)
- **Request (application/json)**: same shape as Create Event
- **Responses**
  - 200 OK: `{ "message": "Event updated successfully." }`
  - 400 Bad Request: `{ "message": "Bad Request" }`
  - 401 Unauthorized: `{ "message": "Not Authorized to update this event." }`
  - 500 Server Error: `{ "message": "Could not fetch Event. Try again Later." }` or `{ "message": "Could not update Event. Try again Later." }`
- **curl**
```bash
TOKEN="<jwt-token>"
curl -X PUT http://localhost:8080/events/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{"name":"Updated Title","description":"Updated","location":"Chennai","dateTime":"2026-01-02T18:00:00Z"}'
```

#### Delete Event

- **DELETE** `/events/:id`
- **Auth**: Required; only the event owner may delete
- **Path Params**: `id` (number)
- **Responses**
  - 200 OK: `{ "message": "Event deleted successfully." }`
  - 400 Bad Request: `{ "message": "Bad Request" }`
  - 401 Unauthorized: `{ "message": "Not Authorized to delete this event." }`
  - 500 Server Error: `{ "message": "Could not fetch Event. Try again Later." }` or `{ "message": "Could not delete Event. Try again Later." }`
- **curl**
```bash
TOKEN="<jwt-token>"
curl -X DELETE http://localhost:8080/events/1 \
  -H "Authorization: $TOKEN"
```

#### Register for Event

- **POST** `/events/:id/register`
- **Auth**: Required
- Registers the authenticated user for the event.
- **Responses**
  - 201 Created: `{ "message": "Registered Successfully." }`
  - 400 Bad Request: `{ "message": "Bad Request" }`
  - 400 Bad Request (invalid event): `{ "message": "Event does not exist." }`
  - 500 Server Error: `{ "message": "Could Not Register." }`
- **curl**
```bash
TOKEN="<jwt-token>"
curl -X POST http://localhost:8080/events/1/register \
  -H "Authorization: $TOKEN"
```

#### Cancel Event Registration

- **DELETE** `/events/:id/register`
- **Auth**: Required
- Cancels the authenticated user’s registration for the event.
- **Responses**
  - 201 Created: `{ "message": "Registeration Cancelled Successfully." }`
  - 400 Bad Request: `{ "message": "Bad Request" }`
  - 500 Server Error: `{ "message": "Could Not Cancel Registeration." }`
- **curl**
```bash
TOKEN="<jwt-token>"
curl -X DELETE http://localhost:8080/events/1/register \
  -H "Authorization: $TOKEN"
```

### Notes

- Timestamps must be RFC3339 format, e.g., `2026-01-01T18:00:00Z`.
- Auth middleware expects the raw JWT in `Authorization` (no `Bearer` prefix).
- Owner-only actions (`PUT /events/:id`, `DELETE /events/:id`) require the token to belong to the `userId` that created the event.
