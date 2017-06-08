<template>
	<div v-if="AppState=='home'" class="w3-section">
  		<button class="w3-button w3-ripple w3-red w3-padding"
  	    		v-on:click="soloOnClick()"
		>
  	  		solo
  	  	</button>
  	  	<button class="w3-button w3-ripple w3-purple"
  	  	  		v-on:click="multiOnClick"
  	  	>
  	  		multi
  	  	</button>
  	</div>
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
	<div v-else-if="AppState=='game'" id="App">
  		<Board
  			:board="this.Board"
  			:cellOnClick="this.cellOnClick"
  		/>
        <div align="left">
        	<span class="w3-badge  w3-large w3-pink">{{ Players[0].Score }} </span>
        </div>
        <div align="right">
        	<span class="w3-badge w3-large w3-pink">{{ Players[1].Score }} </span>
        </div>
        <div v-if="Time!=0">
        	<p> <span class="w3-badge w3-green">{{this.Time}} ms</span></p>
        </div>
        <div v-if="GameOver==true" class="w3-panel" align="center">
          	<div v-if="Players[Winner].Name==UserID" class="w3-green">
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
  	<div v-else>
  		Unknown AppState {{ AppState }}
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
            Players: [{Name:"", Score:0}, {Name:"", Score:0}],
            Current: 0,
            GameOver: false,
            Winner: 0,
          	Time: 0,
          	LoopStatus: 0,
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

			this.startGame()
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
				console.log(`/createRoom ${err.body}`)
			})
		},

		joinRoom: function (RoomName) {

			this.AppStateGame()

			this.RoomName = RoomName
			Vue.http.post('/JoinRoomRequest', { RoomName:RoomName, UserName:this.UserID }).then(res => {
				this.updateGameState(res)
			}, err => {
				console.log(`/joinRoom ${err.body}`)
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
				console.log(`/updateGameState res.json() ${err.body}`);
			});
	  	},

        updateData: function () {
			if (this.AppState != "game") {
				return
			}
			Vue.http.post('/GameStateRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
				console.log(`/GameStateRequest ${err.body}`)
			})
        },

		startGame: function () {
			Vue.http.post('/StartGameRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
				console.log(`/startGame ${err.body}`)
			})
		},

	  	restart: function () {
			Vue.http.post('/StartGameRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
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
                	console.log(`/IAMoveRequest ${err.body}`)
		  	  	})
		  	}, err => {
		  		console.log(`/PlayerMoveRequest ${err.body}`)
			})
		},

        hint: function () {
			Vue.http.post('/HintRequest', { UserName:this.UserID, RoomName:this.RoomName }).then(res => {
				this.updateGameState(res)
			}, err => {
				console.log(`/hintRequest ${err.body}`)
			})
      	},

      	loop: function () {
            this.updateData();

			if (this.LoopStatus == 0) {
				return
			}

            setInterval(function () {
            	this.updateData();
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
