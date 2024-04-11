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

export function loader() {
  const ENV = {
    GRAPHQL_URL: process.env.GRAPHQL_URL,
  }
  return json({ ENV })
}

function App() {
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
  useLoaderData<typeof loader>()

  const ENV = useEnv()
  let client = new Client({
    url: ENV.GRAPHQL_URL ?? '',
    exchanges: [fetchExchange, cacheExchange],
  })
  console.log(ENV.GRAPHQL_URL)
  return (
    <Provider value={client}>
      <App />
    </Provider>
  )
}
