import type { Metadata } from "next";
import { Inter , Roboto , Kanit } from "next/font/google";
import "./globals.css";
import { getServerSession } from "next-auth";
import NextAuthProvider from "@/providers/NextAuthProvider";
import { authOptions } from "./api/auth/[...nextauth]/route";
import { ThemeProvider } from "@/providers/ThemeProvider";
import ReduxProvider from "@/providers/ReduxProvider";
import FontProvider from "@/providers/FontProvider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};



export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  const session = await getServerSession(authOptions);
  return (
    <html lang="en">
      <body className={
        inter.className
        
        }>
          <ThemeProvider
            attribute="class"
            defaultTheme="system"
            enableSystem
            disableTransitionOnChange
          >
            <NextAuthProvider session={session}>
              <ReduxProvider>
                <FontProvider>
                  {children}
                </FontProvider>
              </ReduxProvider>
            </NextAuthProvider>
          </ThemeProvider>
      </body>
    </html>
  );
}
