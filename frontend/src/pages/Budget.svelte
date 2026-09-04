<script>
    import Chart from 'chart.js/auto';
    import {onMount} from "svelte";

    import { GetCategoryInfo } from '../../wailsjs/go/main/App';

    let donutCanvas;
    let chartInstance;
    let info;
    let data;

    const colors = [
        'rgb(255, 99, 132)',
        'rgb(54, 162, 235)',
        'rgb(255, 205, 86)',
        'rgb(75, 192, 192)',
        'rgb(153, 102, 255)',
        'rgb(255, 159, 64)'
    ];

    async function getCategoryInfo() {
        const info = await GetCategoryInfo();

        const data = {
            labels: info.map(c => c.Category),
            datasets: [{
                label: 'Spending by Category',
                data: info.map(c => c.Amount),
                backgroundColor: info.map((_, i) => colors[i % colors.length]),
                hoverOffset: 4
            }]
        };

        chartInstance = new Chart(donutCanvas, {
            type: "doughnut",
            data: data,
        });
    }

    onMount(() => {
        getCategoryInfo()
    });

</script>
<div class="page">
    <div class="donut">
        <canvas bind:this={donutCanvas}></canvas>
    </div>
</div>

<style>
    .page {
        height: 75vh;
        width: 90vw;
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: stretch;
    }
    .donut {
        position: relative; /* Chart.js absolutely-positions the canvas inside its parent when responsive */
        width: 100%;
        height: 100%;
        /*height: 500px; !* a concrete height — height: 100% only works if every ancestor up the tree also has one *!*/
    }
</style>
