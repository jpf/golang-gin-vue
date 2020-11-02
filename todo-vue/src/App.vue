<template>
  <div id="app">
    <h1>To-Do List</h1>
    <form @submit.prevent="customLogin()">
      Email: <input id="email" v-model="email" type="email" size="50"/>
      Password: <input id="password" v-model="password" type="password"/>
      <input type="submit" value="Login"/>
    </form>
    <form @submit.prevent="sendItem()">
      <input type="text" size="50" v-model="todoitem" placeholder="Enter new item"/>
      <input type="submit" value="Submit"/>
    </form>
    <ul>
      <li v-for="item in todolist" v-bind:key="item">{{item}}</li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";
import { OktaAuth } from '@okta/okta-auth-js'
const appData = {
  todolist: ["More to do"],
  token: ''
}
export default {
  name: 'App',
  data() {
    return appData;
  },
  methods: {
    getList: getList,
    sendItem: sendItem,
    customLogin: customLogin
  },
  mounted: function() {
    getList();
  }
}

function getList() {
  axios.get("/api/lists").then( res => {
    appData.todolist = res.data.list
  });
}

async function sendItem() {
  const params = new URLSearchParams();
  params.append('item', this.todoitem);
  await axios.post("/api/lists", params);
  getList()
}

function customLogin() {
  const authClient = new OktaAuth({issuer: 'https://' + process.env.VUE_APP_OKTA_DOMAIN});
  authClient.signIn({
    username: this.email,
    password: this.password
  }).then(transaction => {
    authClient.token.getWithoutPrompt({
      clientId: process.env.VUE_APP_OKTA_CLIENT_ID,
      responseType: ['token'],
      scopes: ['openid', 'email', 'profile'],
      sessionToken: transaction.sessionToken,
      redirectUri: window.location.origin + '/login/callback'
    }).then(response => {
      this.token = response.tokens.accessToken.accessToken;
      alert('Token:' + JSON.stringify(this.token));
    })
  })
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
