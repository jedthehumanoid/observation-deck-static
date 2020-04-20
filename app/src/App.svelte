<script>
  export let data;
  let selected;
  let board = { title: "", decks: [] };
  let showboards = false;
  let params = new URLSearchParams(window.location.search);

  let boardname = params.get("board") || "";

  for (let b of data.boards) {
    if (b.title == boardname) {
      board = b;
    }
  }

  function toggleboards() {
    showboards = !showboards;
  }
</script>

<style>
  .deck {
    margin: 5px;
    padding: 5px;
    width: 22em;
    border: 1px solid black;
    float: left;
    border-radius: 5px;
  }

  .card {
    width: 20em;
    border: 1px solid black;
    border-radius: 5px;
    margin: 10px auto 10px auto;
    padding: 10px;
  }

  #sidebar {
    float: left;
  }
</style>

<main>

  <button on:click={toggleboards}>Boards</button>

  <div id="board">
    {#if showboards}
      <div id="sidebar">
        {#each data.boards as board}
          <a href="?board={board.title}">{board.title}</a>
          <br />
        {/each}
        <br />
        Folders:
        <br />
        {#each data.folders as folder}
          <a href="?folder={folder}">{folder}</a>
          <br />
        {/each}
        <br />
        Labels:
        <br />
        {#each data.labels as label}
          <a href="?label={label}">{label}</a>
          <br />
        {/each}
      </div>
    {/if}
    {#each board.decks as deck}
      <div class="deck">
        <strong>{deck.title}</strong>
        <br />
        {#each deck.cards as card}
          <div class="card">{card}</div>
        {/each}
      </div>
    {/each}

  </div>

</main>
