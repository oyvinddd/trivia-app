<script lang="ts">
	import type { Question } from 'src/types/Question';
	import type { Answer } from 'src/types/Answer';
	import type { AnswerResult } from 'src/types/AnswerResult';

	import { onMount } from 'svelte';

	let questions: Array<Question> = [];
	let questionIndex: number = 0
	let answerText: string

	onMount(async () => {
		try {
			const response = await fetch('/api/questions');
			questions  = await response.json();
		} catch (error) {
			console.log(error);
		}
	});

	/**
	 * Submit an answer to a given question to the API.
	 * @param answer the user's answer
	 */
	async function submitAnswer(answer: Answer) {
		const response = await fetch('/api/questions', {
			method: 'POST',
			body: JSON.stringify(answer)
		})
		const result: AnswerResult = await response.json()
		if (result.score >= 70) {
			questionIndex++;
			answerText = "";
		}
	}
</script>

<svelte:head>
	<title>Trivia App</title>
</svelte:head>

<div class="flex flex-col min-h-full justify-center items-center">
	<h1 class="text-3xl font-bold pb-2">DAILY TRIVIA</h1>
	<div class="min-w-full shadow-xl text-center">
		<div class="bg-yellow-400 px-6 py-8 rounded-t-2xl">
			{#if questions.length > 0 }
				<h3 class="text-xl font-bold">{ questions[questionIndex].category } • { questions[questionIndex].difficulty }</h3>
				<h2 class ="text-2xl font-bold">{@html questions[questionIndex].question }</h2>
			{/if}
		</div>
		<form class="px-6 py-6 bg-white rounded-b-2xl" action="#">
			<input autofocus type="text" bind:value={answerText}
					class="block w-full px-2 py-4 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
					placeholder="Write your answer..."/>
			<button class="sr-only" type="submit" on:click={() => submitAnswer({question_id: questions[questionIndex].id, answer: answerText })}>Submit</button>
		</form>
	</div>
</div>

<style>
</style>
