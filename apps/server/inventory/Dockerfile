FROM oven/bun AS build

RUN apt-get update -y && apt-get install -y openssl

WORKDIR /app

COPY . .

RUN bun install

RUN bun db:generate

EXPOSE 3000

CMD ["bun", "run", "src/runner.ts"]