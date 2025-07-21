const Router = require("koa-router");

const customerService = require("../services/customerService");
const accountService = require("../services/accountService");
const transactionService = require("../services/transactionService");

const router = new Router();

// GET /bank/:id
router.get("/:id", async (ctx) => {
  const customerId = ctx.params.id;

  const customer = await customerService.getCustomer(customerId);
  const accounts = await accountService.getAccountsByCustomerId(customerId);

  // 각 계좌의 거래내역 병합
  for (let account of accounts) {
    const transactions = await transactionService.getTransactionsByAccountId(
      account.id
    );
    account.transactions = transactions;
  }

  ctx.body = {
    customer,
    accounts,
  };
});

module.exports = router;
