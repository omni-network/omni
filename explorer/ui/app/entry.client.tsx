/**
 * By default, Remix will handle hydrating your app on the client for you.
 * You are free to delete this file if you'd like to, but if you ever want it revealed again, you can run `npx remix reveal` ✨
 * For more information, see https://remix.run/file-conventions/entry.client
 */

import { RemixBrowser } from "@remix-run/react";
import { startTransition, StrictMode } from "react";
import { hydrate } from "react-dom";
import { ThemeProvider } from "@material-tailwind/react";

startTransition(() => {
  // FIXME: This is a temporary fix for the issue with the Material Tailwind theme not being applied to the app.
  // I can cause this locally by having dark reader extension enabled and then navigating to the app.
  hydrate(<RemixBrowser />, document);
});
