import { createWebHistory, createRouter } from "vue-router";
import AddPenyakit from "../components/AddPenyakit.vue";
import Prediction from "../components/Prediction.vue";
import History from "../components/History.vue";
import Main from "../components/Main.vue";

const routes = [
  {
    path: "/",
    name: "Main",
    component: Main,
  },
  {
    path: "/penyakit",
    name: "AddPenyakit",
    component: AddPenyakit,
  },
  {
    path: "/riwayat",
    name: "History",
    component: History,
  },
  {
    path: "/prediksi",
    name: "Prediction",
    component: Prediction,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
