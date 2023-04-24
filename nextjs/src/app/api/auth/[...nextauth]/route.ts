import NextAuth from 'next-auth';
import KeycloakProvider from 'next-auth/providers/keycloak';

export const authConfig = {
    providers: [
        KeycloakProvider({
            clientId: "nextjs",
            clientSecret: "lmoJ2CPErZiPXPFNxMTU9RvFDH32w0b1",
            issuer: "http://host.docker.internal:9000/realms/master",
        })
    ],
};

const handler = NextAuth(authConfig);

export { handler as GET, handler as POST };
