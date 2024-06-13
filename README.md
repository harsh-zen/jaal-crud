# Appointment CRUD App with Jaal

This is a simple example of crud application for managing appointments, built using the Jaal GraphQL server library for Go.

## Features

- Create, read, update, and delete appointments
- Scalar type registration for `DateTime` type

## Prerequisites

- Go (version 1.16 or later)

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/harsh-zen/jaal-crud.git
```

2. Navigate to the project directory:

```bash
cd jaal-crud
```

3. Build and run the application:

```bash
go run .
```

The server will start running on `http://localhost:9000/graphql`.

## Schema

The GraphQL schema is defined in the `schema/appointment` package. It includes the following types:

- `Appointment`: Represents an appointment with fields like `id`, `title`, `description`, `startTime`, and `endTime`.
- `CreateAppointmentRequest`: Input type for creating a new appointment.

The schema also defines the following operations:

- `Query`:
  - `appointments`: Retrieves a list of all appointments.
  - `appointment(id: ID!)`: Retrieves a single appointment by ID.
- `Mutation`:
  - `createAppointment(input: CreateAppointmentRequest!): Appointment`: Creates a new appointment.
  - `updateAppointment(id: ID!, input: CreateAppointmentRequest!): Appointment`: Updates an existing appointment.
  - `deleteAppointment(id: ID!): Appointment`: Deletes an appointment by ID.


##Usage
- Get all appointments
``` bash
query {
  appointments {
    id
    title
    description
    startTime
    endTime
  }
}
```

- Get specific appointment
```
query {
  appointment(id: "1") {
    id
    title
    description
    startTime
    endTime
  }
}
```

- Create

``` bash
mutation {
  createAppointment(input: {
    title: "New Appointment"
    startTime: "2024-06-13T10:00:00Z"
    endTime: "2024-06-13T11:00:00Z"
    description: "This is a new appointment"
  }) {
    id
    title
    description
    startTime
    endTime
  }
}
```
and so on...
