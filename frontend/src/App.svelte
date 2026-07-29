<script>
  import Home from './pages/Home.svelte'
  import UploadFile from './pages/UploadFile.svelte'
  import Budget from './pages/Budget.svelte'
  import Loans from './pages/Loans.svelte'
  import Investment from './pages/Investments.svelte'
  import Sankey from './pages/Sankey.svelte'
  import mockData from './assets/mockBudgetData.json'
  // import { parseCsv } from "../wailsjs/go/main/App";

  let currentPage = "home"
</script>

<div>
  <div>Sidebar</div>

  <button on:click={() => currentPage = "home"}>Home</button>
  <button on:click={() => currentPage = "uploadFile"}>Upload Files</button>
  <button on:click={() => currentPage = "budget"}>Personal Budget</button>
  <button on:click={() => currentPage = "loans"}>Loans</button>
  <button on:click={() => currentPage = "investments"}>Investments</button>

  {#if currentPage === "home"}
    <Home />
  {:else if currentPage === "uploadFile"}
    <UploadFile />
  {:else if currentPage === "budget"}
    <Budget />
  {:else if currentPage === "loans"}
    <Loans />
  {:else if currentPage === "investments"}
    <Investment />
  {/if}

</div>