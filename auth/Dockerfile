FROM oven/bun:latest AS common-build-stage

WORKDIR /app

# Cache packages installation
COPY package.json package.json
COPY bun.lockb bun.lockb

RUN bun install

COPY . ./

FROM common-build-stage AS development-build-stage

ENV NODE_ENV=development

CMD ["bun", "run", "dev"]

FROM common-build-stage AS production-build-stage

ENV NODE_ENV=production

RUN bun build \
	--compile \
	--minify-whitespace \
	--minify-syntax \
	--target bun \
	--outfile server \
	./src/server.ts

FROM gcr.io/distroless/base AS production-runtime-stage

WORKDIR /app

COPY --from=production-build-stage /app/server server

ENV NODE_ENV=production

EXPOSE 3000

ENTRYPOINT ["./server"]
