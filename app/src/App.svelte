<script>
  export let cards;
  export let boards;
  let selected;
  let board = { title: "", decks: [] };

  let params = new URLSearchParams(window.location.search);

  let boardname = params.get("board") || "";

  for (let b of boards) {
    if (b.title == boardname) {
      board = b;
    }
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
</style>

<main>
  {#each boards as board}
    <a href="?board={board.title}">{board.title}</a>
    <br />
  {/each}
  {boardname}
  <div id="boards">

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
