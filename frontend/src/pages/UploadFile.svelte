<script>
    import { ParseCsv } from '../../wailsjs/go/main/App';

    let transactions = []
    let error

    function handleFileChange(event) {
        const file = event.target.files[0];
        if (!file) return;

        const reader = new FileReader();
        reader.onload = async () => {
            try {
                error = '';
                if (typeof reader.result !== 'string') {
                    error = 'Failed to read file as text';
                    return;
                }
                transactions = await ParseCsv(reader.result);
                console.log(transactions);
            } catch (err) {
                error = err.message || String(err);
            }
        };
        reader.onerror = () => {
            error = 'Failed to read file';
        };
        reader.readAsText(file);
    }
</script>

<input type="file" id="csvFileInput" accept=".csv" on:change={handleFileChange} />

{#if error}
    <p class="error">{error}</p>
{/if}
{#if transactions.length}
    <ul>
        {#each transactions as t}
            <li>{t.Date} — {t.Description} ({t.Category}): {t.Amount.toFixed(2)}</li>
        {/each}
    </ul>
{/if}
