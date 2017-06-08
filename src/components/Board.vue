<template>
 	<div class="Grid">
		<div class="Row" v-for="(row, x) in board">
			<div v-if="x==0 && y==0" class="TopLeftCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="CellText w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="CellText w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="x== 0 && y > 0 && y < 18" class="LeftCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="x==0 && y == 18" class="BotLeftCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="x > 0 && x < 18 && y==0" class="TopCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="x== 18 && y == 0" class="TopRightCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="x == 18 && y < 18 && y > 0" class="RightCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="x== 18 && y == 18" class="BotRightCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="x > 0 && y > 0 && x < 18 && y < 18" class="Cell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
   			</div>
            <div v-if="y == 18 && x > 0 && x < 18" class="BotCell" v-for="(cell, y) in row" v-on:click="cellOnClick(x, y)" >
				<div v-if="cell==1" class="White w3-green w3-round-xxlarge"></div>
				<div v-else-if="cell==2" class="Black w3-indigo w3-round-xxlarge"></div>
				<div v-else-if="cell==-1" class="CellText w3-red w3-round-xxlarge"></div>
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
    background-image: url("maple.jpg");
    background-size: contain;
}

.Row {
	flex: 1;
	display: flex;
	flex-flow: column nowrap;
}

.Cell {
	flex: 1;
    background-image: url("cross.png");
    background-size: contain;
}

.LeftCell {
	flex: 1;
    background-image: url("leftborder.png");
    background-size: contain;
}
.RightCell {
	flex: 1;
    background-image: url("rightborder.png");
    background-size: contain;
}
.TopCell {
	flex: 1;
    background-image: url("topborder.png");
    background-size: contain;
}
.BotCell {
	flex: 1;
    background-image: url("bottomborder.png");
    background-size: contain;
}
.TopLeftCell {
	flex: 1;
    background-image: url("topleftcorner.png");
    background-size: contain;
}
.TopRightCell {
	flex: 1;
    background-image: url("toprightcorner.png");
    background-size: contain;
}
.BotLeftCell {
	flex: 1;
    background-image: url("botleftcorner.png");
    background-size: contain;
}
.BotRightCell {
	flex: 1;
    background-image: url("botrightcorner.png");
    background-size: contain;
}

.White {
    background-image: url("white.png");
    background-size: contain;
	height:100%;
	width:100%;
}
.Black {
    background-image: url("black.png");
    background-size: contain;
	height:100%;
	width:100%;
}
.CellText {
	height:100%;
	width:100%;
}
</style>
