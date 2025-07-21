const Koa = require("koa");
const Router = require("koa-router");
const bankRoutes = require("./routes/bank");

const app = new Koa();
const router = new Router();

router.use("/bank", bankRoutes.routes());

app.use(router.routes()).use(router.allowedMethods());

app.listen(4000, () => {
  console.log("Server is running on port 4000");
});
