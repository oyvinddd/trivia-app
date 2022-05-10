<script lang="ts">

    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    export let answerText = ""

    function submit() {
        dispatch('submit_answer', answerText)
    }

</script>

<form class="px-6 py-6 bg-white rounded-b-2xl" action="#">
    <input autofocus type="text" bind:value={ answerText }
           class="block w-full px-2 py-4 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
           placeholder="Write your answer..."/>
    <button class="sr-only" type="submit" on:click={ submit }>Submit</button>
</form>