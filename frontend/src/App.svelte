<script>
  import Home from './pages/Home.svelte'
  import UploadFile from './pages/UploadFile.svelte'
  import Budget from './pages/Budget.svelte'
  import Loans from './pages/Loans.svelte'
  import Investment from './pages/Investments.svelte'
  import mockData from './assets/mockBudgetData.json'
  // import { parseCsv } from "../wailsjs/go/main/App";

  let currentPage = "home"
</script>

<div class="app">
  <div class="page-header">
  <button on:click={() => currentPage = "home"}>Home</button>
  <button on:click={() => currentPage = "uploadFile"}>Upload Files</button>
  <button on:click={() => currentPage = "budget"}>Personal Budget</button>
  <button on:click={() => currentPage = "loans"}>Loans</button>
  <button on:click={() => currentPage = "investments"}>Investments</button>
  </div>
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

<style>
  :global(html, body) {
    margin: 0;
    padding: 0;
    height: 100%;
  }
  .app {
    height: 100vh;
    width: 100vw;
    display: flex;
    flex-direction: column;
    overflow: hidden; /* prevents the whole app from scrolling; individual panels can opt into their own scroll */
  }
  .page-header {
    height: 10vh;
    width: 90vw;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    border: aquamarine 2px solid;
  }

</style>