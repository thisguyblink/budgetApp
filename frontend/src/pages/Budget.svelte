<script>
    import Chart from 'chart.js/auto';
    import {onMount} from "svelte";

    import { GroupTransactionsCategory } from '../../wailsjs/go/main/App';

    let donutCanvas;
    let chartInstance;
    let categoryInfo;

    const data = {
        labels: [
            'Red',
            'Blue',
            'Yellow'
        ],
        datasets: [{
            label: 'My First Dataset',
            data: [300, 50, 100],
            backgroundColor: [
                'rgb(255, 99, 132)',
                'rgb(54, 162, 235)',
                'rgb(255, 205, 86)'
            ],
            hoverOffset: 4
        }]
    };

    onMount(() => {
        chartInstance = new Chart(donutCanvas, {
            type: "doughnut",
            data: data,
        })
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
