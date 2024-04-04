import type { LinksFunction } from '@remix-run/node'
import stylesheet from '~/tailwind.css'
import { Links, LiveReload, Meta, Outlet, Scripts, ScrollRestoration, json, useLoaderData } from '@remix-run/react'
import { Client, Provider, cacheExchange, fetchExchange } from 'urql'

export const links: LinksFunction = () => [{ rel: 'stylesheet', href: stylesheet }]

const client = new Client({
  url: 'http://localhost:21335/query',
  exchanges: [fetchExchange, cacheExchange],
})

function GetGraphQLHost(): string {
  const v = (window as any).ENV.GRAPHQL_HOST
  console.log('testing fetching env variables')
  return 'http://localhost:21335/query'
}

export async function loader() {
  console.log('loader called')
  console.log('TEST=' + process.env.TEST?.toString());
  console.log('TEST2=' + process.env.TEST2?.toString());
  console.log('GRAPHQL_HOST=' + process.env.GRAPHQL_HOST?.toString());
  console.log('FOO=' + process.env.FOO?.toString());
  console.log('loader done')
  return json({
    ENV: {
      GRAPHQL_HOST: process.env.GRAPHQL_HOST,
      FOO: process.env.FOO,
      TEST: process.env.TEST,
      TEST2: process.env.TEST2,
    },
  });
}

function App() {
  const data = useLoaderData<typeof loader>();
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
        <script
          dangerouslySetInnerHTML={{
            __html: `window.ENV = ${JSON.stringify(
              data.ENV
            )}`,
          }}
        />
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
