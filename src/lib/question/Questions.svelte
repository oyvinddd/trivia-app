<script lang="ts">

    import type { Question } from 'src/types/Question';

    export let questions: Array<Question>;
    export let questionIndex: number;

    let answerText: string

</script>
<div class="bg-red-400 px-6 py-8 rounded-t-2xl">
    {#if questions.length > 0 }
        <div class="font-bold flex items-stretch min-w-full justify-between">
            <div>Question { questionIndex + 1 } of { questions.length }</div>
            <div class="capitalize">{ questions[questionIndex].difficulty }</div>
        </div>
        <h3 class="text-xl font-bold text-gray-200">{ questions[questionIndex].category }</h3>
        <h2 class ="text-2xl font-bold">{@html questions[questionIndex].question }</h2>
    {/if}
</div>