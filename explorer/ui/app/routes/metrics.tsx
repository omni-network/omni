import type { LoaderFunction } from '@remix-run/node'
import { Gauge, Registry } from 'prom-client'

// Create a new blank registry to register the metrics not to expose system metrics
const reg = new Registry()

const gauge = new Gauge({ name: 'up', help: 'up is always set to 1 when the application is running' })
gauge.set(1)
reg.registerMetric(gauge)

export const loader: LoaderFunction = async () => {
  const data = await reg.metrics()
  return new Response(data)
}
