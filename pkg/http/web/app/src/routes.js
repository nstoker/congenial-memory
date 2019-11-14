import Vue from "vue";
import VueRouter from "vue-router";
import Auth from "@okta/okta-vue";

import Home from "./components/Home";
import Login from "./components/Login";
import GitHubRepoDetails from "./components/GitHubRepoDetails";

Vue.use(VueRouter);
Vue.use(Auth, {
  issuer: "https://dev-680481.okta.com/oauth2/default",
  client_id: "0oa1t7kmnaqyTKFvn357",
  redirect_uri: "http://localhost:8080/implicit/callback",
  scope: "openid profile email"
});

export default new VueRouter({
  mode: "history",
    routes: [
    { path: "/", component: Login },
    { path: "/me", component: Home, meta: { requiresAuth: true } },
    {
      name: "repo-details",
      path: "/repo/:id",
      component: GitHubRepoDetails,
      meta: { requiresAuth: true },
    { path: '/implicit/callback', component: Auth.handleCallback() }
    ]
});
