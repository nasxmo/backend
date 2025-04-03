# ORY Fullstack

This repository provides a full-stack implementation of authentication and authorization using ORY tools, including Kratos, Hydra, and Oathkeeper. It consists of a backend written in Go and a frontend built with Next.js.

## Project Structure

```plaintext
ory-fullstack/
├── docker-compose.yml # Defines multi-container setup
├── backend/           # Backend services in Go
│   ├── Dockerfile     # Docker configuration for backend
│   ├── go.mod         # Go module file
│   ├── go.sum         # Dependency checksums
│   ├── cmd/
│   │   └── api/
│   │       └── main.go # Main entry point for the API
│   ├── internal/
│   │   ├── auth/      # Authentication integrations with ORY tools
│   │   │   ├── kratos.go
│   │   │   ├── hydra.go
│   │   │   └── oathkeeper.go
│   │   ├── handlers/  # Request handlers
│   │   │   ├── api/v1/
│   │   │   │   ├── auth.go
│   │   │   │   └── user.go
│   │   │   └── web/
│   │   │       └── oauth.go
│   │   └── models/   # Database models
│   │       └── user.go
│   └── pkg/          # Utility packages
│       ├── config/   # Configuration settings
│       └── database/ # Database connection
└── frontend/         # Frontend application with Next.js
    ├── Dockerfile    # Docker configuration for frontend
    ├── package.json  # Dependencies and scripts
    ├── public/       # Static assets
    ├── src/          # Source code
    │   ├── app/      # Application components
    │   │   ├── (auth)/       # Authentication pages
    │   │   │   ├── login/
    │   │   │   │   └── page.tsx
    │   │   │   ├── signup/
    │   │   │   │   └── page.tsx
    │   │   │   └── callback/
    │   │   │       └── page.tsx
    │   │   ├── dashboard/    # Dashboard components
    │   │   │   ├── layout.tsx
    │   │   │   ├── page.tsx
    │   │   │   └── settings/
    │   │   │       ├── security/
    │   │   │       │   ├── page.tsx
    │   │   │       │   ├── password/
    │   │   │       │   │   └── page.tsx
    │   │   │       │   └── mfa/
    │   │   │       │       └── page.tsx
    │   │   │       └── profile/
    │   │   │           └── page.tsx
    │   │   └── api/           # API routes
    │   │       ├── auth/
    │   │       │   └── route.ts
    │   │       └── user/
    │   │           └── route.ts
    │   ├── components/        # UI components
    │   │   ├── auth/
    │   │   │   ├── SocialAuthButtons.tsx
    │   │   │   └── EmailPasswordForm.tsx
    │   │   ├── dashboard/
    │   │   │   ├── SidebarNav.tsx
    │   │   │   └── SecurityCard.tsx
    │   │   └── ui/
    │   │       ├── Alert.tsx
    │   │       └── Loading.tsx
    │   ├── lib/               # Helper utilities
    │   │   ├── auth.ts
    │   │   ├── ory.ts
    │   │   └── api.ts
    │   └── styles/            # Global styles
    │       └── globals.css
    └── next.config.js         # Next.js configuration
```

## Getting Started

Prerequisites. Ensure you have the following installed:

- Docker
- Go
- Node.js

### Setup

Clone the repository:

    git clone https://github.com/nasxmo/ory-fullstack.git
    cd ory-fullstack

Initialize the Go module

    cd backend
    go mod init github.com/yourname/ory-fullstack

Install dependencies:

     npm install

Run the application using Docker:

    docker-compose up --build

## Features

- Secure authentication with ORY Kratos
- Authorization with ORY Hydra
- API gateway and request validation using ORY Oathkeeper
- Fully containerized using Docker
- Frontend built with Next.js

## Contributing

Feel free to submit issues and pull requests! 🚀

## License

This project is licensed under the MIT License
