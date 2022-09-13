<template>
  <v-app secondary>
    <v-app-bar
      app
    >
      <v-toolbar-title class="ml-5" @click="$router.push('/')">wish cars</v-toolbar-title>
      <v-row justify="end" class="mr-5">
        <v-menu
        v-if="user" 
        bottom
        min-width="200px"
        rounded
        offset-y
      >
        <template v-slot:activator="{ on }">
          <v-btn
            icon
            x-large
            v-on="on"
          >
            <v-avatar color="primary" size="35">
              <img :src="user.picture" alt="John">
            </v-avatar>
          </v-btn>
        </template>
        <v-card>
          <v-list-item-content class="justify-center">
            <div class="mx-auto text-center">
              <v-avatar color="primary" size="35">
                <img :src="user.picture" alt="John">
              </v-avatar>
              <h3>{{ user.nickname }}</h3>
              <p class="text-caption mt-1">
                {{ user.email }}
              </p>
              <v-divider class="my-3"></v-divider>
              <v-btn depressed rounded text>マイページ</v-btn>
              <v-divider class="my-3"></v-divider>
              <v-btn depressed rounded text class="red" @click="logout">ログアウト</v-btn>
            </div>
          </v-list-item-content>
        </v-card>
      </v-menu>
        <button v-else class="button" @click="login">ログイン</button>
      </v-row>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <v-footer
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({
  data () {
    return {
    }
  },
  watch: {
    '$route.path': {
      handler() {
        if (!this.user) this.$router.push('/')
      }
    }
  },
  computed: {
    user() {
      return this.$auth.$state.user
    }
  },
  created() {
    if (!this.user) this.$router.push('/')
  },
  methods: {
    login () {
      this.$auth.loginWith('auth0')
    },
    logout () {
      this.$auth.logout()
    }
  }
})
</script>
