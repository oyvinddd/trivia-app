<script lang="ts">

	import { onMount } from 'svelte';
	import ConfettiGenerator from "confetti-js";

	// import svelte components
	import Questions from '../lib/question/Questions.svelte';
	import UserInput from '../lib/question/UserInput.svelte';
	import Summary from '../lib/summary/Summary.svelte';

	// import entities
	import type { Question } from 'src/types/Question';
	import type { Answer } from 'src/types/Answer';
	import type { AnswerResult } from 'src/types/AnswerResult';

	// constants
	const maxQuestions: number = 5
	const enum TriviaState {
		Intro, Questions, Summary
	}

	let questions: Array<Question> = [];
	let results: Array<AnswerResult> = [];
	let questionIndex: number = 0
	let currentState: TriviaState = TriviaState.Questions
	let soundSource: string = "sound/crying1.mp3"
	let audioElement: HTMLAudioElement

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
	async function submitAnswer(answerText: string, questionId: number) {
		// exit early if user submits an empty string as an answer
		const answer: Answer = {answer: answerText, question_id: questionId};
		answer.answer = answer.answer.trim();
		if (answer.answer.length == 0) return;
		try {
			const response = await fetch('/api/questions', {
				method: 'POST',
				body: JSON.stringify(answer)
			})
			const result: AnswerResult = await response.json();
			results.push(result);

			if (result.correct) {
				renderConfetti();
			} else {
				await audioElement.play();
			}
			if (questionIndex < maxQuestions - 1) {
				questionIndex++;
				//answerText = "";
			}
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
<audio bind:this={ audioElement } src={ soundSource }></audio>
<div class="flex flex-col min-h-full justify-between">
	<!-- MAIN CARD START -->
	<main class="mb-auto justify-center text-center items-center">
		<h1 class="text-3xl font-bold pb-2">DAILY TRIVIA</h1>
		<div class="min-w-full shadow-xl">
			{#if currentState === TriviaState.Questions }
				<!-- QUESTION PART -->
				<Questions { questions } { questionIndex } />
				<!-- USER INPUT PART -->
				<UserInput on:submit_answer={ (event) => submitAnswer(event.detail, questions[questionIndex].id) } />
			{:else if currentState === TriviaState.Summary }
				<Summary />
			{/if}

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
