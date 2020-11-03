<template>
    <div id="app">
        <h1>To-Do List</h1>
        <div id="widget-container"></div>
        <form @submit.prevent="sendItem()">
            <input
                type="text"
                size="50"
                v-model="todoitem"
                placeholder="Enter new item"
            />
            <input type="submit" value="Submit" />
        </form>
        <ul>
            <li v-for="item in todolist" v-bind:key="item">{{ item }}</li>
        </ul>
        <div>{{ message }}</div>
    </div>
</template>

<script>
import axios from "axios";
import OktaSignIn from "@okta/okta-signin-widget";
import "@okta/okta-signin-widget/dist/css/okta-sign-in.min.css";
const appData = {
    todolist: ["More to do"],
    token: "",
    message: "",
};
export default {
    name: "App",
    data() {
        return appData;
    },
    methods: {
        getList: getList,
        sendItem: sendItem,
    },
    mounted: function () {
        var signIn = new OktaSignIn({
            el: "#widget-container",
            baseUrl: "https://" + process.env.VUE_APP_OKTA_DOMAIN,
            clientId: process.env.VUE_APP_OKTA_CLIENT_ID,
            redirectUri: window.location.origin,
            authParams: {
                issuer:
                    "https://" + process.env.VUE_APP_OKTA_DOMAIN + "/oauth2/default",
                responseType: ["token", "id_token"],
                display: "page",
            },
        });
        signIn.renderEl({ el: "#widget-container" });
        if (signIn.hasTokensInUrl()) {
            signIn.authClient.token.parseFromUrl().then(function success(response) {
                appData.token = response.tokens.accessToken.accessToken;
                axios.defaults.headers.common["Authorization"] =
                    "Bearer " + appData.token;
                appData.message = appData.token;
                signIn.hide();
            });
        }
        getList();
    },
};

function getList() {
    axios.get("/api/lists").then((res) => {
        appData.todolist = res.data.list;
    });
}

async function sendItem() {
    const params = new URLSearchParams();
    params.append("item", this.todoitem);
    await axios
        .post("/api/lists", params)
        .then(function () {
            getList();
        })
        .catch(function (error) {
            appData.todolist = [error.message];
        });
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
