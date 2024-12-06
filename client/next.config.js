module.exports = {
  webpack: (config) => {
    config.watchOptions.poll = 300;
    return config;
  },
  env: {
    // Scuffed Implementation, might have to refactor to next 13 AppDir
    API_GATEWAY_URL: 'http://ingress-nginx-controller.ingress-nginx.svc.cluster.local'
  }
};