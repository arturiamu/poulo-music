import { createRouter, createWebHashHistory } from "vue-router";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      component: () => import("../components/Search.vue"),
      name: "search",
      path: "/search",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("@/pages/Hot/index.vue"),
      name: "hot",
      path: "/hot",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("../components/Recommend.vue"),
      name: "recommend",
      path: "/recommend",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("../components/Recent.vue"),
      name: "recent",
      path: "/recent",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("../components/Artist.vue"),
      name: "artist",
      path: "/artist",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("../components/Album.vue"),
      name: "album",
      path: "/album",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("../components/Song.vue"),
      name: "song",
      path: "/song",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("../components/Playlist.vue"),
      name: "playlist",
      path: "/playlist",
      meta: {
        keep: true,
      },
    },
    {
      component: () => import("../components/Setting.vue"),
      name: "setting",
      path: "/setting",
      meta: {
        keep: true,
      },
    },
    {
      redirect: "/hot",
      path: "/",
    },
  ],
});
export default router;
