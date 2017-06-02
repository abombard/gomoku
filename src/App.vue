<template>
	<div v-if="board == undefined" class="w3-section">
  		<button
  			class="w3-button w3-ripple w3-red w3-padding"
  	    	v-on:click="startGame('solo')"
  	  	>
  	  		Solo
  	  	</button>
  	  	<button
  	  	  	class="w3-button w3-ripple w3-purple"
  	  	  	v-on:click="startGame('multi')"
  	  	>
  	  		Multi
  	  	</button>
  	</div>
	<div v-else id="App">
  		<Board
  			:board="this.board"
  			:cellOnClick="this.play"
  		/>
  	  	<button
  	  	  	class="w3-button w3-ripple w3-purple"
  	  	  	v-on:click="restart()"
  	  	>
  	  		Restart
  	  	</button>
  	</div>
</template>

<script>
	import Board from './components/Board.vue';

	export default {

  	  data () {
  	  	return {
  	  		id: (Math.random() % 255).toString(),
  	  		board: undefined
      	}
  	  },

	  components: { Board },

	  methods: {
        loadData: function () {
			Vue.http.get('/board').then(response => {
				this.updateBoard(response)
			}, err => {
				console.log(`/board ${err.body}`)
			})
              },
	  	  updateBoard: function (res) {
	  	  	  res.json().then(newBoard => {
	  	  	  	  this.board = newBoard;
	  	  	  }, err => {
				  console.log(`res.json() ${err.body}`);
			  });
	  	  },
	  	  startGame: function (mode) {
			console.log(`startGame ${mode}`);
			Vue.http.post('/startgame', { mode:mode, player:this.id }).then(response => {
				this.updateBoard(response)
              this.loop()
			}, err => {
				console.log(`/startgame ${err.body}`)
			})
		  },
		  play: function (x, y) {
		  	  Vue.http.post('/getboard', { x:x, y:y, player:this.id }).then(response => {
				this.updateBoard(response)
                Vue.http.post('/play', { x:x, y:y, player:this.id }).then(response => {
                  this.updateBoard(response)
                }, err => {
                  console.log(`/play ${err.body}`)
		  	  })
		  	  }, err => {
		  	  	console.log(`/getboard ${err.body}`)
		  	  })
		  },
	  	  restart: function () {
			Vue.http.get('/reset').then(response => {
				this.updateBoard(response)
			}, err => {
				console.log(`/reset ${err.body}`)
			})
		  },
      loop: function () {
            this.loadData();

            setInterval(function () {
                    this.loadData();
                  }.bind(this), 3000); 
          }
	  }
	}

</script>

<style scoped>
.App {
	display: flex
}
</style>
