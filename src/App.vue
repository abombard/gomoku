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
      	}
  	  },

	  components: { Board },

	  methods: {
	  	  startGame: function (mode) {
			console.log(`startGame ${mode}`);
			Vue.http.post('/startgame', { mode:mode }).then(response => {
				this.board = response.board;
			}, response => {
				console.log('/startgame Error response')
			})
		  },
		  play: function (x, y) {
		  	  Vue.http.post('/play', { x:x, y:y }).then(response => {
		  	  	  this.board = response.board;
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
