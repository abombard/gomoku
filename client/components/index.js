import Vue, { set } from 'vue'
import Vuex from 'vuex';

Vue.use(Vuex);

const reset = () => (
  players: Array(2).fill(0),
  current: 0,
  cells: Array(19*19).fill(),
}

const isWinner = (player) => (

