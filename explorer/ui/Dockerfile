FROM node:21-slim AS base

ENV NODE_ENV production
RUN corepack enable

RUN mkdir /app
WORKDIR /app

# ADD package.json pnpm-lock.yaml ./
COPY . .

# RUN pnpm install --production=false
# RUN pnpm run build

EXPOSE 3000
# ENTRYPOINT ["pnpm", "run", "start"]
