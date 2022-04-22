<script>
	import { onMount } from "svelte"

	class Question {
		constructor(data) {}
	}

	let question = new Question(null)

	onMount(async () => {
		fetch("/api/question").then(res => res.json()).then(data => {
			question.text = data["question"]
			question.id = data["id"]
		}).catch(err => {
			console.log(err)
		})
	})

</script>

<main>
	<div class="bg-white">
		<h1>Daily Trivia</h1>
		<h2>{ question.text }</h2>
	</div>
</main>

<!--
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
-->
