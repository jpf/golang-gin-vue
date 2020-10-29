<template>
  <div id="app">
    <h1>To-Do List</h1>
    <form method="POST" @submit="sendItem()">
      <input v-model="todoitem" placeholder="Enter new item"/>
      <input type="submit" value="Submit"/>
    </form>
    <ul>
      <li v-for="item in todolist" v-bind:key="item">{{item}}</li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";
const appData = {
  todolist: ["More to do"]
}
export default {
  name: 'App',
  data() {
    return appData;
  },
  mounted: function() {
    this.getList();
  },
  methods: {
    getList: getList,
    sendItem: sendItem
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
