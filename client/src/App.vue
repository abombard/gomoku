<template>
  <div id="app">
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
			this.$http.post('/startgame', { mode:mode }).then(response => {
				this.board = response.board;
			}, response => {
				Console.log('/startgame Error response')
			})
		  },
		  play: function (x, y) {
		  	  this.$http.post('/play', { x: "x", y: "y" }).then(response => {
		  	  	  this.board = response.board;
		  	  }, response => {
		  	  	  Console.log('/play Error response')
		  	  })
		  },
	  }
	}

</script>

<style lang="scss">
</style>
