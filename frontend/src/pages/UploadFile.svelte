<script>
// @ts-nocheck

    import { ParseCsv } from '../../wailsjs/go/main/App';
    import { ClearTransactionsTable } from "../../wailsjs/go/main/App";

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
                error = await ParseCsv(reader.result);
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
<button on:click={ClearTransactionsTable}>ClearTransactions</button>
{#if error}
    <p class="error">{error}</p>
    {:else}
    <p>No Errors with Upload</p>
{/if}



