import { StrictMode } from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createRouter } from "@tanstack/react-router";

// Import the generated route tree
import { routeTree } from "./routeTree.gen";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import "./index.css";
import { ThemeProvider } from "./components/theme-provider";
import { AuthProvider, useAuth } from "./auth";
import { Toaster } from "./components/ui/sonner";

const queryClient = new QueryClient();

import { TransportProvider, useTransport } from "@connectrpc/connect-query";
import { finalTransport } from "./lib/connect";

// Create a new router instance
const router = createRouter({
  routeTree,
  defaultPreload: "intent",
  context: {
    auth: undefined!, // will be set after we wrap the app in AuthProvider
    transport: undefined!,
    queryClient,
  },
  // defaultPreloadStaleTime: 0,
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

function InnerApp() {
  const auth = useAuth();

  const transport = useTransport();

  return <RouterProvider router={router} context={{ auth, transport }} />;
}

// Render the app
const rootElement = document.getElementById("app")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <TransportProvider transport={finalTransport}>
        <QueryClientProvider client={queryClient}>
          <AuthProvider>
            <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
              <InnerApp />
              <Toaster />
            </ThemeProvider>
          </AuthProvider>
        </QueryClientProvider>
      </TransportProvider>
    </StrictMode>,
  );
}
