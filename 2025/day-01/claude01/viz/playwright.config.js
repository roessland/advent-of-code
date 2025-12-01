import { defineConfig, devices } from '@playwright/test';
import net from 'net';

// Synchronously find an available port starting from the given port
function findAvailablePortSync(startPort) {
  for (let port = startPort; port < startPort + 100; port++) {
    try {
      const server = net.createServer();
      server.listen(port);
      server.close();
      return port;
    } catch {
      // Port in use, try next
    }
  }
  throw new Error('No available port found');
}

const port = findAvailablePortSync(8765);

export default defineConfig({
  testDir: './tests',
  fullyParallel: true,
  use: {
    baseURL: `http://localhost:${port}`,
    headless: true,
    screenshot: 'only-on-failure',
  },
  projects: [{
    name: 'chromium',
    use: { ...devices['Desktop Chrome'] },
  }],
  webServer: {
    command: `npx http-server . -p ${port} -c-1`,
    port: port,
    cwd: '.',
    reuseExistingServer: false,
  },
});
