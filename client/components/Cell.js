<template>
  <div
    class="Cell"
    :class="cssClass"
    @click="$emit('click', $event)"
  >
  </div>
</template>

<script>
  export default {
    props: {
      selected: {
        type: Number,
        required: false,
      },
    },

    computed: {
      cssClass () {
        const list = [];

        if (typeof this.selected == 'number') {
          list.push(`is-player-${this.selected}`);
        }

        return list;
      },
    },
  };
</script>

<style scoped>
  .Cell {
    width: 5%;
    height: 5%;
    float: left;

    transform: scale(.95, .95);
    bakground: #333;

    &:hover:not(.is-player-0, .is-player-1) {
      animation: pulsate infinite .8s ease-in-out;
    }

    &.is-player {
      transform: scale(.9, .9);

      &-0 { background: #64CEAA; }
      &-1 { background: #FD6C6C; }
    }
  }

  @keyframes pulsate {
    0%, 100% {
      transform: scale(.9, .9);
    } 50% {
      transform: scale(.8, .8);
    }
  }
</style>
