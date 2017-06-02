<template>
 	<div class="Grid">
		<div class="Row"
			v-for="(row, x) in board">
			<div class="Cell w3-border w3-center"
				v-for="(cell, y) in row"
				v-on:click="cellOnClick(x, y)">
				<div v-if="cell==1" class="CellText w3-green">
					{{  }}
				</div>
				<div v-else-if="cell==2" class="CellText w3-indigo">
   					{{  }}
   				</div>
   			</div>
   		</div>
	</div>
</template>

<script>

	export default {
		props: ['board', 'cellOnClick', 'win'],
	}

</script>

<style scoped>
.Grid {
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

.CellText {
	height:100%;
	width:100%;
}
</style>
