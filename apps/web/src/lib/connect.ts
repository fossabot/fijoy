import { env } from "@/env";
import { createConnectTransport } from "@connectrpc/connect-web";

export const finalTransport = createConnectTransport({
  baseUrl: env.VITE_BACKEND_URL,
  credentials: "include",
  useHttpGet: true,
});
