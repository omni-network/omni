FROM node:21-slim AS base

# Install all node_modules, including dev dependencies
FROM base as deps

RUN mkdir /app
WORKDIR /app

ADD package.json pnpm-lock.yaml ./
RUN corepack enable
RUN pnpm install --production=false

# Setup production node_modules
FROM base as production-deps

ENV NODE_ENV production
RUN corepack enable

RUN mkdir /app
WORKDIR /app

COPY --from=deps /app/node_modules /app/node_modules
ADD package.json pnpm-lock.yaml ./

RUN pnpm prune --production

# Build the app
FROM base as build

RUN mkdir /app
WORKDIR /app

COPY --from=deps /app/node_modules /app/node_modules
RUN corepack enable

ADD . .
RUN pnpm run build

# Finally, build the production image with minimal footprint
FROM base

ENV NODE_ENV production
RUN corepack enable

RUN mkdir /app
WORKDIR /app

COPY --from=deps /app/node_modules /app/node_modules
COPY --from=build /app/build /app/build
COPY --from=build /app/public /app/public
COPY --from=build /app/package.json /app/package.json
COPY --from=build /app/pnpm-lock.yaml /app/pnpm-lock.yaml
ADD . .

EXPOSE 3000
ENTRYPOINT ["pnpm", "run", "start"]
