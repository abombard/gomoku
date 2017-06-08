<template>

<div>
    <!-- display some infos -->
    <div class="w3-panel" style="height:20px">
        <div v-if="AppState=='game' && GameOver==true" align="center">
          	<div v-if="Players[Winner].Name==UserID" class="w3-green">
          		YOU WON !
          	</div>
          	<div v-else class="w3-red">
          		YOU LOST !
          	</div>
        </div>
    	<div v-else-if="AppState=='game' && GameOver!=true && Error!=undefined" class="w3-red">
        	{{this.Error}}
    	</div>
    </div>

	<!-- Home -->
	<div v-if="AppState=='home'" class="w3-section">
      <div align="center">
  		<button class="w3-button w3-ripple w3-red w3-padding w3-round-xxlarge w3-xxlarge"
  	    		v-on:click="soloOnClick()"
		>
  	  		Solo
  	  	</button>
  	  	<button class="w3-button w3-ripple w3-purple w3-padding w3-round-xxlarge w3-xxlarge"
  	  	  		v-on:click="multiOnClick"
  	  	>
  	  		Multi
  	  	</button>
      </div>
  	</div>

  	<!-- Rooms -->
  	<div v-else-if="AppState=='rooms'">
  		<p>Available rooms:</p>
  		<div class="w3-section">
  			<div v-for="(name, key) in Rooms">
				<button class="w3-button w3-ripple w3-red w3-padding"
  		 				v-on:click="joinRoom(name)"
  		 		>
  		 			{{ name }}
  		 		</button>
  			</div>
  		</div>
  		<input v-model="createroomname" placeholder="Enter a room name">
  		<button class="w3-button w3-ripple w3-red w3-padding"
  	   			v-on:click="createRoom(createroomname, 'multi')"
  	    >
  	    	Create Rooms
  	    </button>
  		<div class="w3-section">
  			<button class="w3-button w3-ripple w3-red w3-padding"
  	   				v-on:click="updateRooms"
  	    	>
  	    		Refresh
  	    	</button>
  	    </div>
  	</div>

  	<!-- Game -->
	<div v-else-if="AppState=='game'" id="App" class="w3-container">
  		<Board
  			:board="this.Board"
  			:cellOnClick="this.cellOnClick"
  		/>
        <div align="center">
        <button class="w3-button"
        		v-on:click="previousOnClick">
        	<img src="https://d30y9cdsu7xlg0.cloudfront.net/png/6398-200.png"
        		 style="width:30%; height:30%">
        </button>
        <button class="w3-button"
        		v-on:click="nextOnClick">
        	<img src="https://d30y9cdsu7xlg0.cloudfront.net/png/6402-200.png"
        		 style="width:30%; height:30%">
        </button>
        </div>
        <div align="left">
        	<span class="w3-badge  w3-large w3-pink">{{ Players[0].Score }} </span>
        </div>
        <div align="right">
        	<span class="w3-badge w3-large w3-pink">{{ Players[1].Score }} </span>
        </div>
        <div v-if="Time!=0">
        	<p> <span class="w3-badge w3-green">{{this.Time}} ms</span></p>
        </div>
        <div align="center">
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
  	</div>
	<div v-else>
  		Unknown AppState {{ AppState }}
  	</div>
    <!-- leave some space -->
    <div class="w3-panel">
    </div>
</div>

</template>

<script>
	import Board from './components/Board.vue';

	const createRandomId = function () {
		var text = ""
		var sample = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

		for ( var i=0; i < 8; i ++ ) {
			text += sample.charAt(Math.floor(Math.random() * sample.length))
		}

		return text
	}

	export default {

  	  data () {
  	  	return {
  	  		UserID: createRandomId(),
  	  		AppState: "home",
  	  		Rooms: undefined,
  	  		RoomName: "",
  	  		Board: undefined,
            Players: [{Name:"", Score:0, Index: 0}, {Name:"", Score:0, Index: 0}],
            Current: 0,
            GameOver: false,
            Winner: 0,
          	Time: 0,
          	LoopStatus: 0,
            Error : undefined,
      	}
  	  },

	  components: { Board },

	  methods: {

		// App States
		AppStateHome: function () {
			this.AppState = "home"
			this.LoopStatus = 0
		},
		AppStateRooms: function () {
			this.AppState = "rooms"
			this.LoopStatus = 0
		},
		AppStateGame: function () {
			this.AppState = "game"
			this.LoopStatus = 1
			this.loop()
		},

		// App State Home
		soloOnClick: function () {
			var RoomName = createRandomId()

			this.createRoom(RoomName, "solo")
			this.joinRoom(RoomName)
		},

		multiOnClick: function () {

			this.AppStateRooms()

			this.updateRooms()
		},

		// App State Room
		updateRooms: function () {
			Vue.http.get('/RoomsRequest').then(res => {
				res.json().then(rooms => {
					this.Rooms = rooms;
				}, err => {
					console.log(`/RoomsRequest res.json() ${err.body}`)
				})
			}, err => {
				this.Error = err.body
				console.log(`/RoomsRequest ${err.body}`)
			});
		},

		createRoom: function (RoomName, roomMode) {
			if (RoomName == "") {
				return
			}
			Vue.http.post('/CreateRoomRequest', { RoomName:RoomName, RoomMode:roomMode }).then(res => {
				this.updateRooms()
			}, err => {
				this.Error = err.body
				console.log(`/createRoom ${err.body}`)
			})
		},

		joinRoom: function (RoomName) {

			this.RoomName = RoomName

			Vue.http.post('/JoinRoomRequest', { UserName:this.UserID, RoomName:RoomName }).then(res => {
				this.updateGameState(res)
				this.AppStateGame()

			}, err => {
              	this.Error = err.body
				console.log(`/joinRoomRequest ${err.body}`)
			})
		},

		// App State Game
	  	updateGameState: function (res) {
	  	  	res.json().then(Game => {
	  	 		this.Board = Game.Board;
	  	 		this.Players = Game.Players;
	  	 		this.Current = Game.Current;
	  	 		this.GameOver = Game.GameOver;
	  	 		this.Winner = Game.Winner;
	 	  	  	this.Time = Game.Time;
	  	  	}, err => {
              	this.Error = err.body
				console.log(`/updateGameState res.json() ${err.body}`);
			})
		},

        noerr: function () {
          	this.Error = undefined
        },

        updateData: function () {
			Vue.http.post('/GameStateRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
				this.Error = err.body
				console.log(`/GameStateRequest ${err.body}`)
			})
        },

		startGame: function () {
			Vue.http.post('/StartGameRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
              	this.Error = err.body
				console.log(`/StartGameRequest ${err.body}`)
			})
		},

	  	restart: function () {
			Vue.http.post('/StartGameRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
              	this.Error = err.body
				console.log(`/restartGame ${err.body}`)
			})
		},

		cellOnClick: function (x, y) {
            if (this.GameOver) {
                return
            }
		  	Vue.http.post('/PlayerMoveRequest', { UserName:this.UserID, RoomName:this.RoomName, x:x, y:y }).then(res => {
				this.updateGameState(res)
                if (this.GameOver) {
                	return
                }
                Vue.http.post('/IAMoveRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
                	this.updateGameState(res)
                }, err => {
              		this.Error = err.body
                	console.log(`/IAMoveRequest ${err.body}`)
		  	  	})
		  	}, err => {
              	this.Error = err.body
		  		console.log(`/PlayerMoveRequest ${err.body}`)
			})
		},

        hint: function () {
			Vue.http.post('/HintRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
              	this.Error = err.body
				console.log(`/hintRequest ${err.body}`)
			})
      	},

        previousOnClick: function () {
			Vue.http.post('/HistoryPrevRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
              	this.Error = err.body
				console.log(`/HistoryPrevRequest ${err.body}`)
			})
      	},

        nextOnClick: function () {
			Vue.http.post('/HistoryNextRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
              	this.Error = err.body
				console.log(`/HistoryNextRequest ${err.body}`)
			})
      	},

      	loop: function () {
            this.updateData();

            setInterval(function () {

				this.Error = undefined

				if (this.LoopStatus == 0) {
					return
				}

            	this.updateData();

            }.bind(this), 3000); 
        }
	  }
	}

</script>

<style scoped>
#App {

}
</style>
