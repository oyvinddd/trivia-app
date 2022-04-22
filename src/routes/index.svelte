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
	<div class="px-4 py-2">
		<h1 class="text-3xl font-bold">Daily Trivia</h1>
		{#if question}
			<h2>{question?.text}</h2>
		{/if}
	</div>
</div>

<style>
</style>
