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
import Navbar from './components/shared/navbar'
import {Footer} from './components/shared/footer'
import { gqlClient } from './entry.server'

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
    <html lang="en" data-theme="dark" className="h-full bg-surface">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="bg-surface flex flex-col justify-start content-start  h-full">
        <Navbar />
        <Outlet />
        <div className="grow"></div>
        <Footer />
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

  return (
    <Provider value={gqlClient}>
      <App />
    </Provider>
  )
}
