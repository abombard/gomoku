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
          	<div v-if="error!=undefined" class="w3-red">
          		{{this.error}}
          	</div>

  		<Board
  			:board="this.board"
  			:cellOnClick="this.play"
  		/>
        <div align="center">
        <button class="w3-button w3-circle w3-black" v-on:click="previous()">-</button>
        <button class="w3-button w3-circle w3-black" v-on:click="next()">+</button>
        </div>
        <div align="left">
        	<span class="w3-badge  w3-xlarge w3-pink">{{ players[0].Score }} </span>
        </div>
        <div align="right">
        	<span class="w3-badge w3-xlarge w3-pink">{{ players[1].Score }} </span>
        </div>
        <div v-if="time!=0">
        	<p> <span class="w3-badge w3-green">{{this.time}} ms</span></p>
        </div>
        <div v-if="gameOver==true" class="w3-panel" align="center">
          	<div v-if="players[winner].Name==id" class="w3-green">
          		YOU WON !
          	</div>
          	<div v-else class="w3-red">
          		YOU LOST !
          	</div>
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
            players: [{Name:"", Score:0}, {Name:"", Score:0}],
            current: 0,
            gameOver: false,
            winner: 0,
          	time: 0,
            error : undefined
      	}
  	  },

	  components: { Board },

	  methods: {

        loadData: function () {
			Vue.http.get('/board').then(response => {
				this.updateState(response)
			}, err => {
              this.error = err.body
				console.log(`/board ${err.body}`)
			})
        },

        hint: function () {
			Vue.http.get('/hint').then(response => {
				this.updateState(response)
			}, err => {
              this.error = err.body
				console.log(`/hint ${err.body}`)
			})
      	},

        previous: function () {
			Vue.http.get('/previous').then(response => {
				this.updateState(response)
			}, err => {
              this.error = err.body
				console.log(`/previous ${err.body}`)
			})
      	},
        next: function () {
			Vue.http.get('/next').then(response => {
				this.updateState(response)
			}, err => {
              this.error = err.body
				console.log(`/next ${err.body}`)
			})
      	},
	  	updateState: function (res) {
	  	  	res.json().then(state => {
	  	 		this.board = state.Board;
	  	 		this.players = state.Players;
	  	 		this.current = state.Current;
	  	 		this.gameOver = state.GameOver;
	  	 		this.winner = state.Winner;
	 	  	  	this.time = state.Time;
	  	  	}, err => {
              this.error = err.body
				console.log(`res.json() ${err.body}`);
			});
	  	},
        noerr: function () {
          this.error = undefined
        },

	  	startGame: function (mode) {
			console.log(`startGame ${mode}`);
			Vue.http.post('/startgame', { mode:mode, player:this.id }).then(response => {
				this.updateState(response)
              	this.loop()
			}, err => {
              this.error = err.body
				console.log(`/startgame ${err.body}`)
			})
		},

		play: function (x, y) {
            if (this.gameOver) {
                return
            }
		  	Vue.http.post('/getboard', { x:x, y:y, player:this.id }).then(response => {
				this.updateState(response)
                if (this.gameOver) {
                	return
                }
                Vue.http.post('/play', { x:x, y:y, player:this.id }).then(response => {
                	this.updateState(response)
                }, err => {
              this.error = err.body
                	console.log(`/play ${err.body}`)
		  	  	})
		  	}, err => {
              this.error = err.body
			})
		},

	  	restart: function () {
			Vue.http.get('/reset').then(response => {
				this.updateState(response)
			}, err => {
              this.error = err.body
				console.log(`/reset ${err.body}`)
			})
		},
      
      	loop: function () {
            this.loadData();

            setInterval(function () {
			this.error = undefined
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
