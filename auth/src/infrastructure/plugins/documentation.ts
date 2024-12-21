import swagger from "@elysiajs/swagger";

import config from "@/config";

const swaggerConfig = swagger({
  scalarConfig: {
    servers: ["/"],
  },
  path: "/documentation",
  documentation: {
    info: {
      title: "Bun Elysia Sandbox API Docs",
      version: config.app.version,
    },
    tags: [
      {
        name: "Auth",
        description: "Authentication routes",
      },
    ],
    components: {
      securitySchemes: {
        cookieAuth: {
          type: "apiKey",
          in: "cookie",
          name: "session",
        },
      },
    },
  },
});

export default swaggerConfig;
