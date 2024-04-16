/**
 * Fetch application metrics.
 */

import type { LoaderFunction } from "@remix-run/node";
import { register } from "prom-client";

export const loader: LoaderFunction = async () => {
  const data = await register.metrics()
  return new Response(data)
}
