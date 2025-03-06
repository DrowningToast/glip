# Elysia with Bun runtime

## Getting Started

```sh
copy .env.example .env

docker-compose -f docker-compose.dev.yml up -d

bun install
bun db:migrate
bun db:generate

bun run dev

```

Open http://localhost:3000/ with your browser to see the result.

Swagger: http://localhost:3000/swagger