const axios = require("axios");

exports.getAccountsByCustomerId = async (customerId) => {
  const res = await axios.get(
    `http://localhost:3002/accounts/customer/${customerId}`
  );
  return res.data;
};
