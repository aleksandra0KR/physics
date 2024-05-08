function calculation() {
    // input values
    const L = parseFloat(document.getElementById('pendulumLength').value);
    const L1 = parseFloat(document.getElementById('L1').value);
    const m = parseFloat(document.getElementById('Mass').value);
    const k = parseFloat(document.getElementById('stiffnessCoefficient').value);
    const beta = parseFloat(document.getElementById('attenuationCoefficient').value);
    const phi1 = Math.PI / 180 * parseFloat(document.getElementById('initialDeflection1').value);
    const phi2 = Math.PI / 180 * parseFloat(document.getElementById('initialDeflection2').value);
    const T = parseInt(document.getElementById('time').value);


    const g = 9.82; // constant

    const omega1 = Math.sqrt(g / L); // normal frequency 1
    const omega2 = Math.sqrt(g / L + 2 * k * L1 ** 2 / (m * L ** 2)); // normal frequency 2

    // normal coordinates
    const 𝜉1 = (phi1 + phi2) / 2;
    const 𝜉2 = (phi1 - phi2) / 2;

    // Simulation parameters
    const dt = 0.001; // Time step
    const t = []; // Array to store time values
    const phi1_t = []; // Array to store pendulum 1 angle values
    const phi2_t = []; // Array to store pendulum 2 angle values
    const v1_t = []; // Array to store pendulum 1 velocity values
    const v2_t = []; // Array to store pendulum 2 velocity values

    // Simulation loop
    for (let i = 0; i <= T / dt; i++) {
        const time = i * dt;

        // Calculate angles and velocities at current time step
        const phi12 = (𝜉1 * Math.cos(omega1 * time) + 𝜉2 * Math.cos(omega2 * time)) * Math.exp(-beta * time);
        const phi22 = (𝜉1 * Math.cos(omega1 * time) - 𝜉2 * Math.cos(omega2 * time)) * Math.exp(-beta * time);
        const v1 = (𝜉1 * omega1 * -Math.sin(omega1 * time) + 𝜉2 * omega2 * -Math.sin(omega2 * time)) * Math.exp(-beta * time);
        const v2 = (𝜉1 * omega1 * -Math.sin(omega1 * time) - 𝜉2 * omega2 * -Math.sin(omega2 * time)) * Math.exp(-beta * time);

        // Store values in respective arrays
        t.push(time);
        phi1_t.push(phi12);
        phi2_t.push(phi22);
        v1_t.push(v1);
        v2_t.push(v2);
    }

    // Draw angle plot
    drawAnglePlot(t, phi1_t, phi2_t);

    // Draw velocity plot
    VFromTPlot2(t, v1_t, v2_t);

    const omega1Paragraph = document.createElement('p');
    omega1Paragraph.textContent = `Нормальная частота omega1: ${omega1.toFixed(6)} Герц`;
    normalFrequencies.appendChild(omega1Paragraph);

    const omega2Paragraph = document.createElement('p');
    omega2Paragraph.textContent = `Нормальная частота omega2: ${omega2.toFixed(6)} Герц`;
    normalFrequencies.appendChild(omega2Paragraph);

}

// draw plot phi(t)
function drawAnglePlot(t, phi1, phi2) {

    const trace1 = {
        x: t,
        y: phi1,
        mode: 'lines',
        name: 'Угол phi1',
        line: {color: 'pink', width: 2}
    };

    const trace2 = {
        x: t,
        y: phi2,
        mode: 'lines',
        name: 'Угол phi2',
        line: {color: 'aqua', width: 2}
    };

    const layout = {
        title: 'Зависимость угла от времени для каждого маятника',
        xaxis: {
            title: 'Время (сек)'
        }, yaxis: {
            title: 'Угол (рад)'
        },
        legend: {
            x: 0,
            y: 1,
            font: {
                family: 'Arial, sans-serif',
                size: 12,
                color: 'black'
            }
        }
    };


    Plotly.newPlot('anglePlot', [trace1, trace2], {responsive: true});
}

// draw plot V(t)
function VFromTPlot2(t, v1, v2) {
    const trace1 = {
        x: t,
        y: v1,
        mode: 'lines',
        name: 'Скорость 1',
        line: {color: 'pink', width: 2}
    };

    const trace2 = {
        x: t,
        y: v2,
        mode: 'lines',
        name: 'Скорость 2',
        line: {color: 'aqua', width: 2}
    };

    const layout = {
        title: 'Зависимость скорости от времени для каждого маятника',
        xaxis: {
            title: 'Время (сек)'
        }, yaxis: {
            title: 'Скорость (рад/сек)'
        },
        legend: {
            x: 0,
            y: 1,
            font: {
                family: 'Arial, sans-serif',
                size: 12,
                color: 'black'
            }
        }
    };

    Plotly.newPlot('velocityPlot', [trace1, trace2], layout, {responsive: true});
}


const simulateButton = document.getElementById('start'); // on invent of pressing button call calculation function
simulateButton.addEventListener('click', calculation);


calculation(); // initial simulation
