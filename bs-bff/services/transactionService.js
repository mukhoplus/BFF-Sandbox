const axios = require("axios");

exports.getTransactionsByAccountId = async (accountId) => {
  const res = await axios.get(
    `http://localhost:3003/transactions/account/${accountId}`
  );
  return res.data;
};
