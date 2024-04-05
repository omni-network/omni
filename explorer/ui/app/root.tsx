import type { LinksFunction, SerializeFrom } from '@remix-run/node'
import stylesheet from '~/tailwind.css'
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  json,
  useLoaderData,
} from '@remix-run/react'
import { Client, Provider, cacheExchange, fetchExchange } from 'urql'
import { useEnv } from './lib/use-env'

export const links: LinksFunction = () => [{ rel: 'stylesheet', href: stylesheet }]
export type LoaderData = SerializeFrom<typeof loader>

const client = new Client({
  url: 'http://localhost:21335/query',
  exchanges: [fetchExchange, cacheExchange],
})

// This loads in our environment variables from the .env file
export function loader() {
  const ENV = {
    GRAPHQL_HOST: process.env.GRAPHQL_HOST ?? 'http://localhost:8080/query',
  }
  return json({ ENV })
}

function GetGraphQLHost() {
  const ENV = useEnv()
  return ENV.GRAPHQL_HOST
}

function App() {
  const data = useLoaderData<typeof loader>()
  GetGraphQLHost()
  return (
    <html lang="en" data-theme="light">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        <Outlet />
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  )
}

export default function AppWithProviders() {
  return (
    <Provider value={client}>
      <App />
    </Provider>
  )
}
