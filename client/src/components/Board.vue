<template>
 	<div class="Board">
		<div class="Row"
			v-for="(row, x) in board">
			<div class="Cell"
				v-for="(cell, y) in row"
				v-on:click="cellOnClick(x, y)">
   				{{ cell }}
   			</div>
   		</div>
	</div>
</template>

<script>

	export default {
		props: ['board', 'cellOnClick'],
	}

</script>

<style scoped>
.Board {
	border: 3px solid #fff;
	width: 70vh;
	display: flex;
	height: 70vh;
	margin: auto;
	flex-flow: row nowrap;
}

.Row {
	flex: 1;
	display: flex;
	flex-flow: column nowrap;
}

.Cell {
	border: 1px solid #000;
	flex: 1;
}
</style>
