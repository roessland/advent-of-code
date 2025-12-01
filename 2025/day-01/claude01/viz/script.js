// AoC 2025 Day 1 - Dial Visualization

// Input data (will be loaded)
let inputData = '';

// Parsed rotations: array of {dir: 'L'|'R', dist: number}
let rotations = [];

// Cumulative positions at each step (including start at 50)
// positions[0] = 50 (start)
// positions[i] = position after rotation i-1
let positions = [];

// Zero crossings: array of {step: number, fraction: number, y: number}
// step is which rotation, fraction is 0-1 within that rotation
let crossings = [];

// Envelope
let minY = 0;
let maxY = 0;
let totalSteps = 0;

// Animation state
let animationStart = 0;
const ANIMATION_DURATION = 60000; // 60 seconds
let currentStep = 0;
let currentFraction = 0;

// Canvas and sizing
const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');
let dpr = window.devicePixelRatio || 1;
let WIDTH = 800;
let HEIGHT = 600;
let MARGIN = { top: 30, right: 30, bottom: 30, left: 50 };
let plotWidth = WIDTH - MARGIN.left - MARGIN.right;
let plotHeight = HEIGHT - MARGIN.top - MARGIN.bottom;

// Resize canvas to fill container with retina support
function resizeCanvas() {
    const wrapper = document.getElementById('canvas-wrapper');
    const rect = wrapper.getBoundingClientRect();

    dpr = window.devicePixelRatio || 1;
    WIDTH = rect.width;
    HEIGHT = rect.height;

    // Set canvas size in pixels (scaled for retina)
    canvas.width = WIDTH * dpr;
    canvas.height = HEIGHT * dpr;

    // Scale context to match
    ctx.setTransform(dpr, 0, 0, dpr, 0, 0);

    // Update plot dimensions
    MARGIN = { top: 30, right: 30, bottom: 30, left: 50 };
    plotWidth = WIDTH - MARGIN.left - MARGIN.right;
    plotHeight = HEIGHT - MARGIN.top - MARGIN.bottom;
}

// Parse input data
function parseInput(input) {
    rotations = [];
    const lines = input.trim().split('\n');
    for (const line of lines) {
        const trimmed = line.trim();
        if (!trimmed) continue;
        const dir = trimmed[0];
        const dist = parseInt(trimmed.slice(1), 10);
        rotations.push({ dir, dist });
    }
    totalSteps = rotations.length;
}

// Calculate cumulative positions and find crossings
function calculatePositions() {
    positions = [50]; // Start at 50
    crossings = [];

    let cumulative = 50;

    for (let i = 0; i < rotations.length; i++) {
        const { dir, dist } = rotations[i];
        const startPos = cumulative;
        const delta = dir === 'R' ? dist : -dist;
        const endPos = startPos + delta;

        // Find zero crossings during this rotation
        // Zero occurs at cumulative positions that are multiples of 100
        // Going from startPos to endPos

        if (delta > 0) {
            // Moving right (increasing)
            let nextZero = Math.ceil(startPos / 100) * 100;
            if (nextZero === startPos) nextZero += 100; // Don't count starting position
            while (nextZero <= endPos) {
                const fraction = (nextZero - startPos) / delta;
                crossings.push({ step: i, fraction, y: nextZero });
                nextZero += 100;
            }
        } else if (delta < 0) {
            // Moving left (decreasing)
            let nextZero = Math.floor(startPos / 100) * 100;
            if (nextZero === startPos) nextZero -= 100;
            while (nextZero >= endPos) {
                const fraction = (startPos - nextZero) / (-delta);
                crossings.push({ step: i, fraction, y: nextZero });
                nextZero -= 100;
            }
        }

        cumulative = endPos;
        positions.push(cumulative);
    }

    // Calculate envelope
    minY = Math.min(...positions);
    maxY = Math.max(...positions);

    // Add some padding
    const yRange = maxY - minY;
    minY -= yRange * 0.05;
    maxY += yRange * 0.05;
}

// Convert data coordinates to canvas coordinates
function toCanvasX(step, fraction = 0) {
    const x = (step + fraction) / totalSteps;
    return MARGIN.left + x * plotWidth;
}

function toCanvasY(y) {
    const normalized = (y - minY) / (maxY - minY);
    return MARGIN.top + (1 - normalized) * plotHeight;
}

// Draw zero lines (horizontal lines where dial position mod 100 = 0)
function drawZeroLines() {
    ctx.strokeStyle = 'rgba(255, 100, 100, 0.2)';
    ctx.lineWidth = 0.5;

    // Show labels every 500 instead of every 100
    const minZero = Math.ceil(minY / 500) * 500;
    const maxZero = Math.floor(maxY / 500) * 500;

    ctx.font = '10px monospace';
    ctx.fillStyle = 'rgba(255, 100, 100, 0.5)';

    for (let z = minZero; z <= maxZero; z += 500) {
        const y = toCanvasY(z);
        ctx.beginPath();
        ctx.moveTo(MARGIN.left, y);
        ctx.lineTo(WIDTH - MARGIN.right, y);
        ctx.stroke();

        // Label
        ctx.fillText(`${z}`, 5, y + 3);
    }
}

// Draw axes
function drawAxes() {
    ctx.strokeStyle = '#555';
    ctx.lineWidth = 0.5;

    // Y axis
    ctx.beginPath();
    ctx.moveTo(MARGIN.left, MARGIN.top);
    ctx.lineTo(MARGIN.left, HEIGHT - MARGIN.bottom);
    ctx.stroke();

    // X axis
    ctx.beginPath();
    ctx.moveTo(MARGIN.left, HEIGHT - MARGIN.bottom);
    ctx.lineTo(WIDTH - MARGIN.right, HEIGHT - MARGIN.bottom);
    ctx.stroke();

    // Labels
    ctx.fillStyle = '#888';
    ctx.font = '11px monospace';
    ctx.fillText('Step', WIDTH / 2, HEIGHT - 8);

    ctx.save();
    ctx.translate(12, HEIGHT / 2);
    ctx.rotate(-Math.PI / 2);
    ctx.fillText('Cumulative Position', 0, 0);
    ctx.restore();
}

// Draw the path up to current animation point
function drawPath(upToStep, fraction) {
    if (positions.length < 2) return;

    ctx.strokeStyle = '#00ff88';
    ctx.lineWidth = 1.5; // Visible but not too thick
    ctx.lineJoin = 'round';
    ctx.lineCap = 'round';
    ctx.beginPath();

    ctx.moveTo(toCanvasX(0), toCanvasY(positions[0]));

    for (let i = 1; i <= upToStep && i < positions.length; i++) {
        ctx.lineTo(toCanvasX(i), toCanvasY(positions[i]));
    }

    // Draw partial segment for current step
    if (upToStep < positions.length - 1 && fraction > 0) {
        const startY = positions[upToStep];
        const endY = positions[upToStep + 1];
        const currentY = startY + (endY - startY) * fraction;
        ctx.lineTo(toCanvasX(upToStep, fraction), toCanvasY(currentY));
    }

    ctx.stroke();
}

// Draw crossing markers up to current point
function drawCrossings(upToStep, fraction) {
    ctx.fillStyle = '#ffcc00';

    const dotRadius = Math.max(0.8, Math.min(1.5, WIDTH / 800)); // Much smaller dots

    for (const crossing of crossings) {
        // Check if this crossing has been reached
        if (crossing.step > upToStep) break;
        if (crossing.step === upToStep && crossing.fraction > fraction) continue;

        const x = toCanvasX(crossing.step, crossing.fraction);
        const y = toCanvasY(crossing.y);

        ctx.beginPath();
        ctx.arc(x, y, dotRadius, 0, Math.PI * 2);
        ctx.fill();
    }
}

// Count crossings up to current point
function countCrossings(upToStep, fraction) {
    let count = 0;
    for (const crossing of crossings) {
        if (crossing.step > upToStep) break;
        if (crossing.step === upToStep && crossing.fraction > fraction) continue;
        count++;
    }
    return count;
}

// Get current position
function getCurrentPosition(step, fraction) {
    if (step >= positions.length - 1) {
        return positions[positions.length - 1];
    }
    const startY = positions[step];
    const endY = positions[step + 1];
    return startY + (endY - startY) * fraction;
}

// Main render function
function render() {
    ctx.clearRect(0, 0, WIDTH, HEIGHT);

    drawZeroLines();
    drawAxes();
    drawPath(currentStep, currentFraction);
    drawCrossings(currentStep, currentFraction);

    // Draw current position marker
    const currentY = getCurrentPosition(currentStep, currentFraction);
    const currentX = toCanvasX(currentStep, currentFraction);

    const markerRadius = Math.max(3, Math.min(5, WIDTH / 200));
    ctx.fillStyle = '#00ffff';
    ctx.beginPath();
    ctx.arc(currentX, toCanvasY(currentY), markerRadius, 0, Math.PI * 2);
    ctx.fill();

    // Update info display
    document.getElementById('step').textContent = currentStep;
    document.getElementById('total-steps').textContent = totalSteps;
    document.getElementById('position').textContent = (((Math.round(currentY) % 100) + 100) % 100);
    document.getElementById('crossings').textContent = countCrossings(currentStep, currentFraction);
}

// Animation loop
function animate(timestamp) {
    if (!animationStart) animationStart = timestamp;

    const elapsed = timestamp - animationStart;
    const progress = (elapsed % ANIMATION_DURATION) / ANIMATION_DURATION;

    // Map progress to step + fraction
    const totalProgress = progress * totalSteps;
    currentStep = Math.floor(totalProgress);
    currentFraction = totalProgress - currentStep;

    // Clamp to valid range
    if (currentStep >= totalSteps) {
        currentStep = totalSteps - 1;
        currentFraction = 1;
    }

    render();
    requestAnimationFrame(animate);
}

// Handle window resize
function handleResize() {
    resizeCanvas();
    render();
}

// Load input and start
async function init() {
    try {
        const response = await fetch('./input.txt');
        if (!response.ok) throw new Error('Failed to load');
        inputData = await response.text();
    } catch (e) {
        console.error('Failed to load input.txt, using example data');
        inputData = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`;
    }

    parseInput(inputData);
    calculatePositions();

    console.log(`Loaded ${totalSteps} rotations`);
    console.log(`Position range: ${minY} to ${maxY}`);
    console.log(`Total zero crossings: ${crossings.length}`);

    // Initial resize
    resizeCanvas();

    // Listen for resize events
    window.addEventListener('resize', handleResize);

    requestAnimationFrame(animate);
}

// Export for testing
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { parseInput, calculatePositions, rotations, positions, crossings };
}

init();
