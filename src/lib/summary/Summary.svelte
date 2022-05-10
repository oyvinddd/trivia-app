<script lang="ts">
    import type { AnswerResult } from 'src/types/AnswerResult';

    export let results: Array<AnswerResult>;

    const urGood: string = "✅"
    const urBad: string = "💩"

</script>

<table class="w-full text-left">
    {#each results as result, i }
        <tr class="{ i % 2 === 0 ? 'bg-gray-100' : '' }">
            <td class="p-4">Question #{ i + 1 }</td>
            <td class="p-2">{ result.score }%</td>
            <td class="p-2">{ result.correct ? urGood : urBad }</td>
        </tr>
    {/each}
</table>