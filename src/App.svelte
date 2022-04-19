<script>
	import { onMount } from "svelte"

	let questionText = ""

	// TODO: base url and API url here
	onMount(async () => {
		fetch("/api/question").then(res => res.json()).then(data => {
			questionText = data["text"]
		}).catch(err => {
			console.log(err)
		})
	})

</script>

<main>
	<h1>Daily Question</h1>
	<h2 class="blue-200 shadow-lg">{ questionText }</h2>

</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>