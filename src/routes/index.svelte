<script lang="ts">
	import ConfettiGenerator from "confetti-js";

	import type { Question } from 'src/types/Question';
	import type { Answer } from 'src/types/Answer';
	import type { AnswerResult } from 'src/types/AnswerResult';

	import { onMount } from 'svelte';

	const maxQuestions: number = 5

	let questions: Array<Question> = [];
	let questionIndex: number = 0
	let soundSource: string = "sound/crying1.mp3"
	let audioElement: HTMLAudioElement
	let answerText: string

	onMount(fetchDailyQuestions);

	/**
	 * Fetch a list of 5 daily trivia questions from the API.
	 */
	async function fetchDailyQuestions() {
		try {
			const response = await fetch('/api/questions')
			questions = await response.json()
		} catch (error) {
			console.log(error);
		}
	}

	/**
	 * Submit an answer to a given question to the API.
	 * @param answer the user's answer
	 */
	async function submitAnswer(answer: Answer) {
		try {
			const response = await fetch('/api/questions', {
				method: 'POST',
				body: JSON.stringify(answer)
			})
			const result: AnswerResult = await response.json()
			if (result.correct) {
				renderConfetti();
			} else {
				await audioElement.play();
			}
			questionIndex++;
			answerText = "";
		} catch (error) {
			console.log(error);
		}
	}

	function renderConfetti() {
		// FIXME: try initializing this library once, globally
		// initialize confetti library with custom configuration
		const settings = { target: 'confetti-canvas', max: 300, size: 1.5, clock: 35, rotate: false, respawn: false };
		const confetti = new ConfettiGenerator(settings);
		confetti.render();
	}
</script>

<svelte:head>
	<title>Daily Trivia</title>
</svelte:head>

<canvas id="confetti-canvas"></canvas>
<audio bind:this={audioElement} src={soundSource}></audio>
<div class="flex flex-col min-h-full justify-between">
	<!-- MAIN CARD START -->
	<main class="mb-auto justify-center text-center items-center">
		<h1 class="text-3xl font-bold pb-2">DAILY TRIVIA</h1>
		<div class="min-w-full shadow-xl">
			<div class="bg-yellow-400 px-6 py-8 rounded-t-2xl">
				{#if questions.length > 0 }
					<!-- QUESTION DETAILS START -->
					<div class="flex items-stretch min-w-full bg-blue-200">
						<div class="bg-green-200 font-bold">Question { questionIndex + 1 } of { maxQuestions }</div>
						<div class="bg-red-200 font-bold">{ questions[questionIndex].difficulty }</div>
					</div>
					<!-- QUESTION DETAILS END -->
					<h3 class="text-xl font-bold text-gray-200">{ questions[questionIndex].category }</h3>
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
	</main>
	<!-- MAIN CARD END -->
	<!-- FOOTER START -->
	<footer class="bg-blue-500 text-center">
		<h4>Footer</h4>
	</footer>
	<!-- FOOTER END -->
</div>

<style>
	#confetti-canvas {
		position: fixed;
		inset: 0;
		pointer-events: none;
	}
</style>
