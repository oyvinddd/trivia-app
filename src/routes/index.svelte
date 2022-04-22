<script lang="ts">
	import type { Question } from 'src/types/Question';

	import { onMount } from 'svelte';

	// let questions: Array<Question> = [];
	let question: Question | undefined;

	onMount(async () => {
		try {
			const response = await fetch('/api/question');
			const { id, question: text } = await response.json();

			question = {
				id,
				text
			};
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
			<div class="bg-yellow-400 px-6 py-8 rounded-t-lg">
				<h2 class="text-2xl font-bold">When was IBM founded? (hardcoded question)</h2>
				{#if question}
					<h2>{question?.text}</h2>
				{/if}
			</div>
			<div class="px-6 py-6 bg-white rounded-b-lg">
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
