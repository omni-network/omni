/**
 * Fetch application metrics.
 */

import type { LoaderFunction } from '@remix-run/node'
import { register, Gauge  } from 'prom-client'

const gauge = new Gauge({ name: 'up', help: 'up is always set to 1 when the application is running' })
gauge.set(1)
register.registerMetric(gauge)

export const loader: LoaderFunction = async () => {
  const data = await register.metrics()
  return new Response(data)
}
