<template>
  <div
    class="Board"
    :class="cssClass"
  >
    <Cell
      v-for="(selected, id) in cells"
      :selected="selected"
      @click="selectCell(id)"
    />

    <div
      class="Board__reset"
      v-id="isGameOver"
      @click="reset"
    >
    </div>
  </div>
</template>

<script>
  import Cell from 'Cell';
  
  export default {
    components: {
      Cell,
    },

    computed: {
      cssClass () {
        const list = [];

        if (this.$store.getters.isWinnerPresent) {
          list.push('is-winner');
        }

        return list;
      },

      cells () {
        return this.$store.state.cells;
      }

      isGameOver () {
        return this.$store.getters.isGameOver;
      },
    },

    methods: {
      selectCell (cellId) {
        this.$store.dispatch('selectCell', cellId);
      },

      reset () {
        this.$store.dispatch('reset');
      },
    },
  };
</script>

<style scoped>
  .Board {
    position: relative;

    width: 100%;
    height: 100%;
  }

  .Board.is-winner .Cell {
    animation: winner 1.3s infinite ease-in-out;
  }

  .Board__reset {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }

  @keyframes winner {
    0%, 70%, 100% {
      transform: scale(.9, .9);
    } 35% {
      transform: scale(0, 0);
    }
  }
</style>
