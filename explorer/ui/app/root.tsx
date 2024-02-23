import { cssBundleHref } from '@remix-run/css-bundle'
import type { LinksFunction } from '@remix-run/node'
import stylesheet from '~/tailwind.css'
import { Links, LiveReload, Meta, Outlet, Scripts, ScrollRestoration } from '@remix-run/react'
import { Client, Provider, cacheExchange, fetchExchange } from 'urql';

export const links: LinksFunction = () => [{ rel: 'stylesheet', href: stylesheet }]

const client = new Client({
  url: 'http://localhost:8080/query',
  exchanges: [fetchExchange, cacheExchange],
});

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
  return (
    <Provider value={client}>
      <App />
    </Provider>
  )
}
