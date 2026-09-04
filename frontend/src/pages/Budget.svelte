<script>
    import Chart from 'chart.js/auto';

    import { GetCategoryTotals } from '../../wailsjs/go/main/App';
    import { GetTotalSpending } from '../../wailsjs/go/main/App';

    let donutCanvas;
    let chartInstance;
    let totalSpending = 0;
    let categoryTotals = {};
    let loading = false;
    let error = null;

    const colors = [
        'rgb(255, 99, 132)',
        'rgb(54, 162, 235)',
        'rgb(255, 205, 86)',
        'rgb(75, 192, 192)',
        'rgb(153, 102, 255)',
        'rgb(255, 159, 64)'
    ];

    // sort by amount descending, recalculates whenever categoryTotals changes
    $: sortedCategories = Object.entries(categoryTotals).sort((a, b) => b[1] - a[1]);

    async function getCategoryInfo() {
        loading = true;
        error = null;
        try {
            const totals = await GetCategoryTotals();
            totalSpending = await GetTotalSpending();
            categoryTotals = totals;

            // use the same sorted order for the chart so slice colors match the list
            const sorted = Object.entries(totals).sort((a, b) => b[1] - a[1]);

            const data = {
                labels: sorted.map(([category]) => category),
                datasets: [{
                    label: 'Spending by Category',
                    data: sorted.map(([, amount]) => amount),
                    backgroundColor: sorted.map((_, i) => colors[i % colors.length]),
                    hoverOffset: 4
                }]
            };

            if (chartInstance) {
                chartInstance.destroy();
            }

            chartInstance = new Chart(donutCanvas, {
                type: "doughnut",
                data: data,
                options: {
                    plugins: {
                        legend: {
                            display: false
                        }
                    }
                }
            });
        } catch (err) {
            error = err;
            console.error(err);
        } finally {
            loading = false;
        }
    }
</script>

<div class="page">
    <h1>Personal Budget</h1>
    <button on:click={getCategoryInfo} disabled={loading}>
        {loading ? 'Loading...' : 'Load Data'}
    </button>
    {#if error}
        <p class="error">{error}</p>
    {:else}
        <p>Total Spending: ${totalSpending.toFixed(2)}</p>
    {/if}

    <div class="donut">
        <canvas bind:this={donutCanvas}></canvas>
    </div>

    <ul class="category-list">
        {#each sortedCategories as [category, amount], i}
            <li>
                <span class="swatch" style="background-color: {colors[i % colors.length]}"></span>
                <span class="label">{category}</span>
                <span class="amount">${amount.toFixed(2)}</span>
            </li>
        {/each}
    </ul>
</div>

<style>
    .page {
        height: 75vh;
        width: 90vw;
        display: flex;
        flex-direction: column;
        align-items: stretch;
        overflow-y: auto;
    }
    .donut {
        position: relative;
        flex: 1;
        width: 100%;
        min-height: 300px;
    }
    .category-list {
        list-style: none;
        margin: 0;
        padding: 0;
    }
    .category-list li {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        padding: 0.4rem 0;
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    }
    .swatch {
        width: 12px;
        height: 12px;
        border-radius: 50%;
        flex-shrink: 0;
    }
    .label {
        flex: 1;
    }
    .amount {
        font-variant-numeric: tabular-nums;
    }
</style>