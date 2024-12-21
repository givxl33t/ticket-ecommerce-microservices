import app from "@infrastructure/app";

import config from "@/config";

app.listen(config.app.port, () => {
  console.log(`🦊 Current Environment: ${config.app.env} `);
  console.log(`🍔 Bun Elysia is running at ${app.server?.hostname}:${app.server?.port}`);
});
