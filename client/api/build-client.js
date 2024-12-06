import axios from 'axios';

const buildClient = ({ req }) => {
  if (typeof window === 'undefined') {
    // we are on the server!
    const response = axios.create({
      baseURL: process.env.API_GATEWAY_URL,
      headers: req.headers
    });

    return response;
  } else {
    // we are on the browser!
    return axios.create({
      baseURL: '/'
    })
  }
}

export default buildClient;