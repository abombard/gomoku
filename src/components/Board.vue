<template>
 	<div class="Grid">
		<div class="Row"
			v-for="(row, x) in board">
			<div class="Cell w3-border w3-center"
				v-for="(cell, y) in row"
				v-on:click="cellOnClick(x, y)">
				<div v-if="cell==1" class="CellText w3-green w3-round-xxlarge">
					{{  }}
				</div>
				<div v-else-if="cell==2" class="CellText w3-indigo w3-round-xxlarge">
   					{{  }}
   				</div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge">
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
	width: 60vh;
	display: flex;
	height: 60vh;
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
