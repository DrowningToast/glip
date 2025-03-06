FROM oven/bun AS build

WORKDIR /app

# Cache packages installation
COPY package.json package.json
COPY bun.lock bun.lock

COPY ./src ./src
COPY ./prisma ./prisma

RUN bun install

ENV NODE_ENV=production

RUN bunx prisma generate --generator binary-targets=linux-arm64-openssl-1.1.x

RUN bun build \
	--compile \
	--minify-whitespace \
	--minify-syntax \
	--target bun \
	--outfile server \
	./src/index.ts

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=build /app/server server

ENV NODE_ENV=production

CMD ["./server"]

EXPOSE 3000