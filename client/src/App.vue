<template>
  <div>
  	  <button
  	  	id="StartGame"
  	  	v-on:click="startGame"
  	  >
  	  	Start Game
  	  </button>
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
  	  	var b = new Array(19);
  	  	for (var i = 0; i < 19; i ++) {
  	  		b[i] = new Array(19).fill("0");
  	  	}
  	  	return {
      	  board: b
      	}
  	  },

	  components: { Board },

	  methods: {
	  	  startGame: function (mode) {
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
