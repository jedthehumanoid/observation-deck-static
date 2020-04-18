import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
	    cards: data.cards,
	    boards: data.boards
	}
});

export default app;
