const axios = require("axios");

exports.getCustomer = async (id) => {
  const res = await axios.get(`http://localhost:3001/customer/${id}`);
  return res.data;
};
