<template>
  <div v-if="board === undefined">
  	  <button v-on:click="startGame('solo')">
  	  	  Solo
  	  </button>
  	  <button v-on:click="startGame('multi')">
  	  	  Multi
  	  </button>
  </div>
  <div v-else>
  	  <Board
  		:board="this.board"
  		:cellOnClick="this.play"
  	  />
  </div>
</template>

<script>
	import Board from './components/Board.vue';

	export default {

  	  data () {
  	  	return {
  	  		board: undefined
      	}
  	  },

	  components: { Board },

	  methods: {
	  	  updateBoard: function (res) {
	  	  	  res.json().then(newBoard => {
	  	  	  	  this.board = newBoard;
	  	  	  }, err => {
				console.log("res.json() error");
			  });
	  	  },
	  	  startGame: function (mode) {
			console.log(`startGame ${mode}`);
			Vue.http.post('/startgame', { mode:mode }).then(response => {
				this.updateBoard(response)
			}, response => {
				console.log('/startgame Error response')
			})
		  },
		  play: function (x, y) {
		  	  Vue.http.post('/play', { x:x, y:y, player:"toto" }).then(response => {
				this.updateBoard(response)
		  	  }, response => {
		  	  	console.log('/play Error response')
		  	  })
		  },
	  }
	}

</script>

<style>
.App {
	display: flex;
}

.Board {
	border: 3px solid #fff;
	width: 70vh;
	display: flex;
	height: 70vh;
	margin: auto;
	flex-flow: row nowrap;
}
</style>
