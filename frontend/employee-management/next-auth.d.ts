import NextAuth from "next-auth";

declare module "next-auth" {
    interface Session {
        user : {
            employee_id: string,
            name : string,
            email: string,
            role: string,
            token: string
        }
    }
}