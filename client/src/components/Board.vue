<template>
 	<div class="Board">
		<div id="Line"
			v-for="(line, x) in board">
			<div id="Cell"
				v-for="(cell, y) in line"
				v-on:click="cellOnClick(x, y)">
				this is a Cell
   			</div>
   		</div>
	</div>
</template>

<script>

	export default {
		props: ['board', 'cellOnClick'],
	}

</script>
