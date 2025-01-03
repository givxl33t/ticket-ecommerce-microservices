import type { Metadata } from "next";
import { AppRouterCacheProvider } from "@mui/material-nextjs/v14-appRouter";
import { ThemeProvider } from "@mui/material/styles";
import { CssBaseline } from "@mui/material";
import Header from "@/components/Header";
import { AuthProvider } from "@/providers/auth";
import theme from "@/config/theme";

export const metadata: Metadata = {
  title: "Ticket Commerce",
  description: "Buy and sell tickets here!",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <AuthProvider>
          <AppRouterCacheProvider>
            <ThemeProvider theme={theme}>
              <Header />
              <CssBaseline />
              <div style={{ marginTop: "80px", padding: "1rem" }}>{children}</div>
            </ThemeProvider>
          </AppRouterCacheProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
