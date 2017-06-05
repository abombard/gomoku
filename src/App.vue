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
        <div v-if="time!=0">
        <p> <span class="w3-badge w3-green">{{this.time}} ms</span></p>
        </div>
          <div v-if="win==true" class="w3-panel w3-green" align="center">
          YOU WON !
        </div>
          <div v-if="lost==true" class="w3-panel w3-red" align="center">
          YOU LOST !
        </div>
  	  	<button
  	  	  	class="w3-button w3-ripple w3-purple"
  	  	  	v-on:click="restart()"
  	  	>
  	  		Restart
  	  	</button>
  	  	<button
  	  	  	class="w3-button w3-ripple w3-blue"
  	  	  	v-on:click="hint()"
  	  	>
  	  		Hint
  	  	</button>
  	</div>
</template>

<script>
	import Board from './components/Board.vue';

	export default {

  	  data () {
  	  	return {
  	  		id: (Math.random() % 255).toString(),
  	  		board: undefined,
            win: false,
          lost: false,
          time: 0
      	}
  	  },

	  components: { Board },

	  methods: {
        loadData: function () {
			Vue.http.get('/board').then(response => {
				this.updateBoard(response)
                if (response.status === 202 && !this.win) {
                  this.lost = true
                }
			}, err => {
				console.log(`/board ${err.body}`)
			})
              },
        hint: function () {
			Vue.http.get('/hint').then(response => {
				this.updateBoard(response)
                if (response.status === 202 && !this.win) {
                  this.lost = true
                }
			}, err => {
				console.log(`/hint ${err.body}`)
			})
              },
	  	  updateBoard: function (res) {
	  	  	  res.json().then(newBoard => {
	  	  	  	  this.board = newBoard;
	  	  	  }, err => {
				  console.log(`res.json() ${err.body}`);
			  });
	  	  },
	  	  updateBoardAndTime: function (res) {
	  	  	  res.json().then(newBoard => {
	  	  	  	  this.board = newBoard.Board;
	  	  	  	  this.time = newBoard.Time;
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
            if (this.win || this.lost) {
                return
            }
		  	  Vue.http.post('/getboard', { x:x, y:y, player:this.id }).then(response => {
                if (response.status === 201) {
                  this.win = true
                }
				this.updateBoard(response)
                Vue.http.post('/play', { x:x, y:y, player:this.id }).then(response => {
                  this.updateBoardAndTime(response)
                }, err => {
                  console.log(`/play ${err.body}`)
		  	  })
		  	  }, err => {
		  	  })
		  },
	  	  restart: function () {
            this.win = false
            this.lost = false
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
