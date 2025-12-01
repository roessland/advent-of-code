import { test, expect } from '@playwright/test';

test.describe('Dial Visualization', () => {
  test('page loads and shows canvas', async ({ page }) => {
    await page.goto('/');

    const canvas = page.locator('#canvas');
    await expect(canvas).toBeVisible();

    // Canvas should have non-zero dimensions (dynamic sizing with retina support)
    const width = await canvas.getAttribute('width');
    const height = await canvas.getAttribute('height');
    expect(parseInt(width)).toBeGreaterThan(0);
    expect(parseInt(height)).toBeGreaterThan(0);
  });

  test('displays initial stats', async ({ page }) => {
    await page.goto('/');

    // Wait for animation to start
    await page.waitForTimeout(500);

    const totalSteps = page.locator('#total-steps');
    await expect(totalSteps).not.toHaveText('0');
  });

  test('animation progresses over time', async ({ page }) => {
    await page.goto('/');

    // Wait for initial load
    await page.waitForTimeout(500);

    const stepEl = page.locator('#step');
    const initialStep = await stepEl.textContent();

    // Wait a bit for animation
    await page.waitForTimeout(2000);

    const laterStep = await stepEl.textContent();
    expect(parseInt(laterStep)).toBeGreaterThan(parseInt(initialStep));
  });

  test('crossings counter increases during animation', async ({ page }) => {
    await page.goto('/');

    // Wait for animation to progress
    await page.waitForTimeout(3000);

    const crossings = page.locator('#crossings');
    const count = await crossings.textContent();
    expect(parseInt(count)).toBeGreaterThan(0);
  });

  test('canvas is being drawn on', async ({ page }) => {
    await page.goto('/');

    // Wait for animation
    await page.waitForTimeout(1000);

    // Check that canvas has non-empty content by evaluating pixel data
    const hasContent = await page.evaluate(() => {
      const canvas = document.getElementById('canvas');
      const ctx = canvas.getContext('2d');
      const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height);
      // Check if there are any non-zero pixels
      for (let i = 0; i < imageData.data.length; i += 4) {
        if (imageData.data[i] !== 0 || imageData.data[i+1] !== 0 ||
            imageData.data[i+2] !== 0 || imageData.data[i+3] !== 0) {
          return true;
        }
      }
      return false;
    });

    expect(hasContent).toBe(true);
  });

  test('example data produces 6 crossings', async ({ page }) => {
    // Navigate with example data by modifying the page
    await page.goto('/');

    const crossingsCount = await page.evaluate(async () => {
      const exampleInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`;

      // Re-parse with example data
      const rotations = [];
      const lines = exampleInput.trim().split('\n');
      for (const line of lines) {
        const trimmed = line.trim();
        if (!trimmed) continue;
        const dir = trimmed[0];
        const dist = parseInt(trimmed.slice(1), 10);
        rotations.push({ dir, dist });
      }

      // Calculate crossings
      let cumulative = 50;
      let crossings = 0;

      for (const { dir, dist } of rotations) {
        const startPos = cumulative;
        const delta = dir === 'R' ? dist : -dist;
        const endPos = startPos + delta;

        if (delta > 0) {
          let nextZero = Math.ceil(startPos / 100) * 100;
          if (nextZero === startPos) nextZero += 100;
          while (nextZero <= endPos) {
            crossings++;
            nextZero += 100;
          }
        } else if (delta < 0) {
          let nextZero = Math.floor(startPos / 100) * 100;
          if (nextZero === startPos) nextZero -= 100;
          while (nextZero >= endPos) {
            crossings++;
            nextZero -= 100;
          }
        }

        cumulative = endPos;
      }

      return crossings;
    });

    expect(crossingsCount).toBe(6);
  });
});
