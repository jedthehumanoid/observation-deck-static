<script>
	export let cards;
	export let boards;
	let selected;
	let board = {title: "",decks:[]}
	
 function setBoard(title) {
   for (let b of boards) {
       if (b.title==title) {
       	 board=b;
       }
    }
 }

$: setBoard(selected)

</script>

<main>
	<label for="boards">Board:</label><select id="boards" bind:value={selected}>
	{#each boards as board}
	       <option value="{board.title}">{board.title}</option>	
	{/each}
	</select>

	{selected}

	
<div id="boards">


{#each board.decks as deck}
<div class="deck"><strong>{deck.title}</strong><br/>
{#each deck.cards as card}
<div class="card">{card}</div>


{/each}
</div>
{/each}



</div>

</main>

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