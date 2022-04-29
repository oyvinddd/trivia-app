<script lang="ts">
	import type { Question } from 'src/types/Question';

	import { onMount } from 'svelte';

	let questions: Array<Question> = [];

	onMount(async () => {
		try {
			const response = await fetch('/api/questions');
			questions  = await response.json();
		} catch (error) {
			console.log(error);
		}
	});
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<div class="flex flex-col min-h-full justify-center items-center">
	<h1 class="text-3xl font-bold pb-2">DAILY TRIVIA</h1>
		<div class="min-w-full shadow-xl text-center">
			<div class="bg-yellow-400 px-6 py-8 rounded-t-2xl">
				{#if questions}
					<h2 class="text-2xl font-bold">{questions[0].category} • {questions[0].text} • { questions[0].text }</h2>
				{/if}
			</div>
			<div class="px-6 py-6 bg-white rounded-b-2xl">
				<input type="text"
						class="block w-full px-2 py-4 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
						placeholder="Write your answer..."/>
			</div>
		</div>
	<div class="pt-2">
		<p>made by <a>oyvind</a> and <a>sindre</a></p>
	</div>
</div>

<style>
</style>
