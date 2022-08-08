<template>
  <div class="wrap">
    <p
      v-if="username"
      class="message"
    >
      ようこそ {{ username }} さん
    </p>
    <button
      v-if="loggedIn"
      class="button"
      @click="logout"
    >
      ログアウト
    </button>
    <button
      v-else
      class="button"
      @click="login"
    >
      ログイン
    </button>
  </div>
</template>

<script>
export default {
  name: 'IndexPage',
  computed: {
    username () {
      if (this.$auth.$state.user &&
          this.$auth.$state.user.name) {
        return this.$auth.$state.user.name
      } else {
        return null
      }
    },
    loggedIn () {
      return this.$auth.loggedIn
    }
  },
  methods: {
    login () {
      this.$auth.loginWith('auth0')
    },
    logout () {
      this.$auth.logout()
    }
  }
}
</script>

<style scoped>
.wrap {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.message {
  font-size: 2.5rem;
}

.button {
  font-size: 2rem;
}
</style>